package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type dataType string

const (
	DataTypeJSON dataType = "json"
	DataTypeFile dataType = "file"
	contentType           = "Content-Type"
)

type DataHandler func(data interface{}, w http.ResponseWriter, successCode int) error

type RouteParams struct {
	body   []byte
	Params map[string]string
}

func (p RouteParams) PopulateBody(target interface{}) error {
	return json.Unmarshal(p.body, &target)
}

var dataTypeHandlers = map[dataType]DataHandler{
	DataTypeJSON: jsonHandler,
	DataTypeFile: fileHandler,
}

type RouteHandler func(request *http.Request, p RouteParams) (interface{}, error)

type RouteBuilder struct {
	path        string
	handler     RouteHandler
	methods     []string
	returnCode  int
	apiVersion  string
	dataType    *dataType
	queryParams []QueryParam
}

func (b RouteBuilder) SetApiVersion(v string) RouteBuilder {
	b.apiVersion = v

	return b
}

func (b RouteBuilder) SetHandler(fn RouteHandler) RouteBuilder {
	b.handler = fn

	return b
}

func (b RouteBuilder) SetMethod(methods ...string) RouteBuilder {
	b.methods = methods
	return b
}

func (b RouteBuilder) SetReturnCode(code int) RouteBuilder {
	b.returnCode = code

	return b
}

func (b RouteBuilder) AddQueryParam(q QueryParam) RouteBuilder {
	b.queryParams = append(b.queryParams, q)

	return b
}

func (b RouteBuilder) SetPath(uri string) RouteBuilder {
	b.path = uri

	return b
}

func (b RouteBuilder) SetDataType(t dataType) RouteBuilder {
	b.dataType = &t

	return b
}

func (b RouteBuilder) Validate() error {
	errs := make([]string, 0)

	if b.returnCode == 0 {
		errs = append(errs, "Must provide a return code")
	}

	if b.handler == nil {
		errs = append(errs, "Must provide a handler function")
	}

	if b.path == "" {
		errs = append(errs, "Must provide a URI path")
	}

	if len(b.methods) == 0 {
		errs = append(errs, "Must provide HTTP method(s)")
	}

	if len(errs) != 0 {
		return errors.New(strings.Join(errs, ", "))
	}

	return nil
}

func (b RouteBuilder) Build(router *mux.Router) {
	if validationErr := b.Validate(); validationErr != nil {
		panic(fmt.Sprintf("Failed to build route due to: %s", validationErr.Error()))
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == http.MethodOptions {
			return
		}

		responseBody, err := (func() (i interface{}, err error) {
			p := RouteParams{Params: mux.Vars(r)}

			if r.Body != nil {
				bBody, err := ioutil.ReadAll(r.Body)
				defer r.Body.Close()

				if err != nil {
					return i, err
				}
				p.body = bBody
			}
			return b.handler(r, p)
		})()

		if err != nil {
			sendError(w, err)

			return
		}

		dataHandler, err := getDataHandler(b)

		if err != nil {
			sendError(w, err)

			return
		}

		err = dataHandler(responseBody, w, b.returnCode)

		if err != nil {
			sendError(w, err)

			return
		}
	}

	path := fmt.Sprintf("/api/%s%s", b.apiVersion, b.path)

	hasRequiredQuery := false
	hasQuery := len(b.queryParams) > 0

	/* 	Due to the nature of the mux package we need to register a route twice if it has optional query params.
	* 	Once for the route without any query params and once with the query so we can match in both scenarios.
	* 	However, if the route has a required query param do not register the base route so the router will only match
	*	if the param is provided
	 */
	if hasQuery {
		for _, q := range b.queryParams {
			if q.Required {
				hasRequiredQuery = true
				break
			}
		}
	}

	if hasQuery {
		router.HandleFunc(path, handler).Methods(b.methods...).Queries(buildQuery(b.queryParams)...)

		if !hasRequiredQuery {
			// all params optional, therefore allow matching the route without query params
			router.HandleFunc(path, handler).Methods(b.methods...)
		}

	} else {
		// no query param, just register the route
		router.HandleFunc(path, handler).Methods(b.methods...)
	}

}

func NewV1RouteBuilder() RouteBuilder {
	return RouteBuilder{apiVersion: "v1"}
}

func jsonHandler(data interface{}, w http.ResponseWriter, successCode int) error {
	preparedBody, err := prepareResponseBody(data, nil)

	if err != nil {
		return err
	}

	w.Header().Set(contentType, "application/json")
	w.WriteHeader(successCode)
	w.Write(preparedBody)

	return nil
}

func fileHandler(file interface{}, w http.ResponseWriter, successCode int) error {

	if fileBytes, ok := file.([]byte); !ok {
		return errors.New("handler did not return a file")
	} else {
		w.Header().Set("Expires", "0")
		w.Header().Set("Content-Transfer-Encoding", "binary")
		w.Header().Set("Content-Type", http.DetectContentType(fileBytes))
		w.Header().Set("Content-Length", strconv.Itoa(int(len(fileBytes))))

		w.WriteHeader(successCode)

		reader := bytes.NewReader(fileBytes)
		io.Copy(w, reader)
	}

	return nil
}

func getDataHandler(b RouteBuilder) (DataHandler, error) {
	// unless specified treat all responses as JSON
	if b.dataType == nil {
		return jsonHandler, nil
	}

	if h, ok := dataTypeHandlers[*b.dataType]; ok {
		return h, nil
	} else {
		return nil, fmt.Errorf("no data handler for responses of type %s", *b.dataType)
	}

}

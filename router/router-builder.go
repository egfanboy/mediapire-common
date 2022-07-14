package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type RouteParams struct {
	body        []byte
	PathParams  map[string]string
	QueryParams map[string]string
}

func (p RouteParams) PopulateBody(target interface{}) error {
	return json.Unmarshal(p.body, &target)
}

type RouteHandler func(request *http.Request, p RouteParams) (interface{}, error)

type RouteBuilder struct {
	path       string
	handler    RouteHandler
	methods    []string
	returnCode int
	apiVersion string
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

func (b RouteBuilder) SetPath(uri string) RouteBuilder {
	b.path = uri

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
		responseBody, err := (func() (i interface{}, err error) {
			p := RouteParams{PathParams: mux.Vars(r)}

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

		preparedBody, err := prepareResponseBody(responseBody, nil)

		if err != nil {
			sendError(w, err)

			return
		}

		w.WriteHeader(b.returnCode)
		w.Write(preparedBody)
	}

	path := fmt.Sprintf("/api/%s%s", b.apiVersion, b.path)
	router.HandleFunc(path, handler).Methods(b.methods...)
}

func NewV1RouteBuilder() RouteBuilder {
	return RouteBuilder{apiVersion: "v1"}
}

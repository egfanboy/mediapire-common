package router

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/egfanboy/mediapire-common/exceptions"
)

func prepareResponseBody(data interface{}, key *string) ([]byte, error) {

	if key != nil {
		responseMap := make(map[string]interface{})

		if data != nil {
			responseMap[*key] = data
		}
		data = responseMap
	}

	response, err := json.Marshal(data)

	if err != nil {
		return nil, errors.New("failed to marshal data to JSON")
	}

	return response, nil
}

func sendError(w http.ResponseWriter, errToSend error) {
	errorKey := "error"

	response, err := prepareResponseBody(errToSend.Error(), &errorKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	apiException, ok := errToSend.(*exceptions.ApiException)

	httpCode := http.StatusInternalServerError

	if ok {
		httpCode = apiException.StatusCode
	}

	w.Header().Set(contentType, "application/json")

	w.WriteHeader(httpCode)
	w.Write(response)
}

package router

import (
	"encoding/json"
	"errors"
	"mediapire-common/exceptions"
	"net/http"
)

func prepareResponseBody(data interface{}, key *string) ([]byte, error) {
	bodyKey := "response"

	if key != nil {
		bodyKey = *key
	}

	responseMap := make(map[string]interface{})

	responseMap[bodyKey] = data

	response, err := json.Marshal(responseMap)

	if err != nil {
		return nil, errors.New("failed to marshal data to JSON")
	}

	return response, nil
}

func sendError(w http.ResponseWriter, err error) {
	errorKey := "error"

	response, err := prepareResponseBody(err.Error(), &errorKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	apiException, ok := err.(*exceptions.ApiException)

	httpCode := http.StatusInternalServerError

	if ok {
		httpCode = apiException.StatusCode
	}

	w.WriteHeader(httpCode)
	w.Write(response)
}

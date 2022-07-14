package exceptions

import "net/http"

type ApiException struct {
	err        error
	StatusCode int
}

func (e ApiException) Error() string {
	return e.err.Error()
}

func NewBadRequestException(cause error) *ApiException {
	return &ApiException{err: cause, StatusCode: http.StatusBadRequest}
}

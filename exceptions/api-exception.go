package exceptions

import "net/http"

type ApiException struct {
	Err        error
	StatusCode int
}

func (e ApiException) Error() string {

	if e.Err != nil {
		return e.Err.Error()
	}

	return ""
}

func NewBadRequestException(cause error) *ApiException {
	return &ApiException{Err: cause, StatusCode: http.StatusBadRequest}
}

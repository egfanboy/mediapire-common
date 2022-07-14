package router

import "github.com/gorilla/mux"

func NewRouter() *mux.Router {
	return mux.NewRouter()
}

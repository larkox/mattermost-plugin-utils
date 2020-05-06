package common

import "net/http"

type Router interface {
	PathPrefix(string) Route
	HandleFunc(string, func(http.ResponseWriter, *http.Request)) Route
}

type Route interface {
	Subrouter() Router
	Methods(...string) Route
}

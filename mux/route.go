package mux

import (
	"net/http"
	"reflect"
	"runtime"
)

// Route route
type Route struct {
	Method string
	Name   string
	Path   string
	Func   string
}

func add(m, n, p string, h http.HandlerFunc) {
	routes = append(
		routes,
		Route{
			Method: m,
			Name:   n,
			Path:   p,
			Func:   runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name(),
		},
	)
}

var routes []Route

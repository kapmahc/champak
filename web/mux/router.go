package mux

import "net/http"

// Router http router
type Router interface {
	Add(method, name, path string, hnd http.HandlerFunc)
}

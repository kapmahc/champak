package logging

import (
	"net/http"
	"reflect"
	"runtime"
)

// New new router
func New() *Router {
	return &Router{
		routes: make([]Route, 0),
	}
}

// Router router for logging
type Router struct {
	routes []Route
}

// Add add
func (p *Router) Add(method, name, path string, hnd http.HandlerFunc) {
	p.routes = append(
		p.routes,
		Route{
			Method: method,
			Name:   name,
			Path:   path,
			Func:   runtime.FuncForPC(reflect.ValueOf(hnd).Pointer()).Name(),
		},
	)
}

// URL get url by name
func (p *Router) URL(name string, _ ...string) string {
	for _, r := range p.routes {
		if name == r.Name {
			return r.Path
		}
	}
	return "not-found"
}

// Walk walk routes
func (p *Router) Walk(fn func(Route)) {
	for _, r := range p.routes {
		fn(r)
	}
}

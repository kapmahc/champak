package web

import "github.com/julienschmidt/httprouter"

// NewRouter new logging router
func NewRouter() *LoggingRouter {
	return &LoggingRouter{routes: make([]route, 0)}
}

// Router router
type Router interface {
	GET(path string, handle httprouter.Handle)
	POST(path string, handle httprouter.Handle)
	PUT(path string, handle httprouter.Handle)
	PATCH(path string, handle httprouter.Handle)
	DELETE(path string, handle httprouter.Handle)
}
type route struct {
	method string
	path   string
	handle httprouter.Handle
}

// LoggingRouter router for logging
type LoggingRouter struct {
	routes []route
}

// Walk walk routes
func (p *LoggingRouter) Walk(f func(method, path string, handle httprouter.Handle) error) error {
	for _, r := range p.routes {
		if e := f(r.method, r.path, r.handle); e != nil {
			return e
		}
	}
	return nil
}

// GET http get
func (p *LoggingRouter) GET(path string, handle httprouter.Handle) {
	p.add("GET", path, handle)
}

// POST http post
func (p *LoggingRouter) POST(path string, handle httprouter.Handle) {
	p.add("POST", path, handle)
}

// PUT http put
func (p *LoggingRouter) PUT(path string, handle httprouter.Handle) {
	p.add("PUT", path, handle)
}

// PATCH http patch
func (p *LoggingRouter) PATCH(path string, handle httprouter.Handle) {
	p.add("PATCH", path, handle)
}

// DELETE http delete
func (p *LoggingRouter) DELETE(path string, handle httprouter.Handle) {
	p.add("DELETE", path, handle)
}

func (p *LoggingRouter) add(method string, path string, handle httprouter.Handle) {
	p.routes = append(p.routes, route{
		method: method,
		path:   path,
		handle: handle,
	})
}

package mux

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	_mux "github.com/gorilla/mux"
	"github.com/jinzhu/inflection"
)

const (
	// GET http get
	GET = "GET"
	// POST http post
	POST = "POST"
	// PATCH http patch
	PATCH = "PATCH"
	// PUT http put
	PUT = "PUT"
	// DELETE http delete
	DELETE = "DELETE"
)

// Router rest router
type Router struct {
	Router *_mux.Router
}

// Get http get
func (p *Router) Get(name, path string, hnd http.HandlerFunc) {
	p.add(GET, name, path, hnd)
}

// Post http post
func (p *Router) Post(name, path string, hnd http.HandlerFunc) {
	p.add(POST, name, path, hnd)
}

// Patch http patch
func (p *Router) Patch(name, path string, hnd http.HandlerFunc) {
	p.add(PATCH, name, path, hnd)
}

// Put http put
func (p *Router) Put(name, path string, hnd http.HandlerFunc) {
	p.add(PUT, name, path, hnd)
}

// Delete http delete
func (p *Router) Delete(name, path string, hnd http.HandlerFunc) {
	p.add(DELETE, name, path, hnd)
}

// Crud http crud resources
func (p *Router) Crud(
	name,
	prefix string,
	_new http.HandlerFunc,
	create http.HandlerFunc,
	edit http.HandlerFunc,
	update http.HandlerFunc,
	show http.HandlerFunc,
	index http.HandlerFunc,
	destroy http.HandlerFunc,
) {
	sn := inflection.Singular(name)
	if _new != nil {
		p.Get(fmt.Sprintf("%s.new", sn), fmt.Sprintf("%s/new", prefix), _new)
	}
	if create != nil {
		p.Post("", prefix, create)
	}
	if edit != nil {
		p.Get(fmt.Sprintf("%s.edit", sn), fmt.Sprintf("%s/{id}/edit", prefix), edit)
	}
	if update != nil {
		p.Post("", fmt.Sprintf("%s/{id}", prefix), update)
	}
	if show != nil {
		p.Get(fmt.Sprintf("%s.show", sn), fmt.Sprintf("%s/{id}", prefix), show)
	}
	if destroy != nil {
		p.Delete("", fmt.Sprintf("%s/{id}", prefix), destroy)
	}
	if index != nil {
		p.Get(name, prefix, index)
	}
}

// Form get and post
func (p *Router) Form(name, path string, get http.HandlerFunc, post http.HandlerFunc) {
	p.Get(name, path, get)
	p.Post("", path, post)
}

// URL url by name
func (p *Router) URL(name string, args ...interface{}) string {
	var pairs []string
	for _, v := range args {
		pairs = append(pairs, fmt.Sprintf("%v", v))
	}
	if r := p.Router.Get(name); r != nil {
		u, e := r.URL(pairs...)
		if e == nil {
			return u.String()
		}
		log.Error(e)
	}
	return "not-found"
}

// WalkFunc walk func
type WalkFunc func(method, name, path string) error

// Walk walk routes
func (p *Router) Walk(fn WalkFunc) error {
	return p.Router.Walk(
		func(route *_mux.Route, router *_mux.Router, ancestors []*_mux.Route) error {
			pth, err := route.GetPathTemplate()
			if err != nil {
				return err
			}
			// FIXME
			// https://github.com/gorilla/mux/pull/207
			// method, _ := route.GetInformation(_mux.InformationMethods)
			return fn("", route.GetName(), pth)
		},
	)
}

func (p *Router) add(method, name, path string, hnd http.HandlerFunc) {
	r := p.Router.HandleFunc(path, hnd).Methods(method)
	if name != "" {
		r.Name(name)
	}
}

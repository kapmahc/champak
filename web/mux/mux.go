package mux

import (
	"fmt"
	"net/http"

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

// New new mux
func New(rt Router) *Mux {
	return &Mux{rt: rt}
}

// Mux mux
type Mux struct {
	rt Router
}

// Get http get
func (p *Mux) Get(name, path string, hnd http.HandlerFunc) {
	p.rt.Add(GET, name, path, hnd)
}

// Post http post
func (p *Mux) Post(name, path string, hnd http.HandlerFunc) {
	p.rt.Add(POST, name, path, hnd)
}

// Patch http patch
func (p *Mux) Patch(name, path string, hnd http.HandlerFunc) {
	p.rt.Add(PATCH, name, path, hnd)
}

// Put http put
func (p *Mux) Put(name, path string, hnd http.HandlerFunc) {
	p.rt.Add(PUT, name, path, hnd)
}

// Delete http delete
func (p *Mux) Delete(name, path string, hnd http.HandlerFunc) {
	p.rt.Add(DELETE, name, path, hnd)
}

// Crud http crud resources
func (p *Mux) Crud(
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
func (p *Mux) Form(name, path string, get http.HandlerFunc, post http.HandlerFunc) {
	p.Get(name, path, get)
	p.Post("", path, post)
}

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

var (
	router *_mux.Router
)

// Use use
func Use(r *_mux.Router) {
	router = r
}

// Get http get
func Get(name, path string, hnd http.HandlerFunc) {
	add(GET, name, path, hnd)
}

// Post http post
func Post(name, path string, hnd http.HandlerFunc) {
	add(POST, name, path, hnd)
}

// Patch http patch
func Patch(name, path string, hnd http.HandlerFunc) {
	add(PATCH, name, path, hnd)
}

// Put http put
func Put(name, path string, hnd http.HandlerFunc) {
	add(PUT, name, path, hnd)
}

// Delete http delete
func Delete(name, path string, hnd http.HandlerFunc) {
	add(DELETE, name, path, hnd)
}

// Crud http crud resources
func Crud(
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
		Get(fmt.Sprintf("%s.new", sn), fmt.Sprintf("%s/new", prefix), _new)
	}
	if create != nil {
		Post("", prefix, create)
	}
	if edit != nil {
		Get(fmt.Sprintf("%s.edit", sn), fmt.Sprintf("%s/{id}/edit", prefix), edit)
	}
	if update != nil {
		Post("", fmt.Sprintf("%s/{id}", prefix), update)
	}
	if show != nil {
		Get(fmt.Sprintf("%s.show", sn), fmt.Sprintf("%s/{id}", prefix), show)
	}
	if destroy != nil {
		Delete("", fmt.Sprintf("%s/{id}", prefix), destroy)
	}
	if index != nil {
		Get(name, prefix, index)
	}
}

// Form get and post
func Form(name, path string, get http.HandlerFunc, post http.HandlerFunc) {
	Get(name, path, get)
	Post("", path, post)
}

// URL url by name
func URL(name string, args ...interface{}) string {
	var pairs []string
	for _, v := range args {
		pairs = append(pairs, fmt.Sprintf("%v", v))
	}
	if r := router.Get(name); r != nil {
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
func Walk(fn WalkFunc) error {
	return router.Walk(
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

func add(method, name, path string, hnd http.HandlerFunc) {
	r := router.HandleFunc(path, hnd).Methods(method)
	if name != "" {
		r.Name(name)
	}
}

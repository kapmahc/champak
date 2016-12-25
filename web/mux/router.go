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

// Router http router
type Router interface {
	Add(method, name, path string, hnd http.HandlerFunc)
	URL(name string, pairs ...string) string
}

// Get http get
func Get(name, path string, hnd http.HandlerFunc) {
	router.Add(GET, name, path, hnd)
}

// Post http post
func Post(name, path string, hnd http.HandlerFunc) {
	router.Add(POST, name, path, hnd)
}

// Patch http patch
func Patch(name, path string, hnd http.HandlerFunc) {
	router.Add(PATCH, name, path, hnd)
}

// Put http put
func Put(name, path string, hnd http.HandlerFunc) {
	router.Add(PUT, name, path, hnd)
}

// Delete http delete
func Delete(name, path string, hnd http.HandlerFunc) {
	router.Add(DELETE, name, path, hnd)
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
		switch v.(type) {
		case string:
			pairs = append(pairs, v.(string))
		default:
			pairs = append(pairs, fmt.Sprintf("%v", v))
		}
	}
	return router.URL(name, pairs...)
}

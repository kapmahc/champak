package gorilla

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

// New new gorilla router
func New(r *mux.Router) *Router {
	return &Router{
		rt: r,
	}
}

// Router gorilla router
type Router struct {
	rt *mux.Router
}

// Add add
func (p *Router) Add(method, name, path string, hnd http.HandlerFunc) {
	r := p.rt.HandleFunc(path, hnd).Methods(method)
	if name != "" {
		r.Name(name)
	}
}

// URL url by name
func (p *Router) URL(name string, args ...interface{}) string {
	var pairs []string
	for _, v := range args {
		pairs = append(pairs, fmt.Sprintf("%v", v))
	}
	if r := p.rt.Get(name); r != nil {
		u, e := r.URL(pairs...)
		if e == nil {
			return u.String()
		}
		log.Error(e)
	}
	return "not-found"
}

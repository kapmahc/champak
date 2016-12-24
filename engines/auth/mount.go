package auth

import "github.com/gorilla/mux"

// Mount mount web points
func (p *Engine) Mount(rt *mux.Router) {
	rt.HandleFunc("/", p.getHome).Methods("GET").Name("home")
}

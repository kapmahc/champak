package auth

import "github.com/gorilla/mux"

// Mount mount web points
func (p *Engine) Mount(rt *mux.Router) {
	rt.HandleFunc("/", p.getHome).Methods("GET").Name("home")

	rt.HandleFunc("/users/sign-in", p.getUsersSignIn).
		Methods("GET").
		Name("auth.users.sign-in.fm")
	rt.HandleFunc("/users/sign-in", p.postUsersSignIn).
		Methods("POST").
		Name("auth.users.sign-in")
}

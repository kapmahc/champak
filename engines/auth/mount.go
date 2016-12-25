package auth

import "github.com/kapmahc/champak/web/mux"

// Mount mount web points
func (p *Engine) Mount() {
	lya := "application"
	mux.Get("home", "/", p.Helper.HTML("home", lya, p.getHome))
}

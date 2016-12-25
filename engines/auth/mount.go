package auth

import "github.com/kapmahc/champak/web/mux"

// Mount mount web points
func (p *Engine) Mount() {
	mux.Get("home", "/", p.getHome)
}

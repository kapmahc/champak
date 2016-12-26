package auth

import "github.com/kapmahc/champak/web"

// Mount mount web points
func (p *Engine) Mount(rt web.Router) {
	rt.POST("/users/sign-up", p.W.Form(&fmSignUp{}, p.postUsersSignUp))
}

package auth

import "github.com/kapmahc/champak/web"

// Mount mount web points
func (p *Engine) Mount(rt web.Router) {
	rt.POST("/users/sign-up", p.W.Form(&fmSignUp{}, p.postUsersSignUp))
	rt.GET("/users/confirm/:token", p.W.JSON(p.getUsersConfirm))
	rt.POST("/users/confirm", p.W.Form(&fmEmail{}, p.postUsersConfirm))
	rt.GET("/users/unlock/:token", p.W.JSON(p.getUsersUnlock))
	rt.POST("/users/unlock", p.W.Form(&fmEmail{}, p.postUsersUnlock))
	rt.POST("/users/forgot-password", p.W.Form(&fmEmail{}, p.postUsersForgotPassword))
	rt.POST("/users/reset-password", p.W.Form(&fmResetPassword{}, p.postUsersResetPassword))
}

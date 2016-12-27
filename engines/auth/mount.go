package auth

import (
	"fmt"

	"github.com/kapmahc/champak/web"
)

func (p *Engine) signInURL() string {
	return fmt.Sprintf("%s/users/sign-in", web.Frontend())
}

// Mount mount web points
func (p *Engine) Mount(rt web.Router) {
	rt.POST("/users/sign-in", p.W.Form(&fmSignIn{}, p.postUsersSignIn))
	rt.POST("/users/sign-up", p.W.Form(&fmSignUp{}, p.postUsersSignUp))
	rt.GET("/users/confirm/:token", p.W.Redirect(p.signInURL(), p.getUsersConfirm))
	rt.POST("/users/confirm", p.W.Form(&fmEmail{}, p.postUsersConfirm))
	rt.GET("/users/unlock/:token", p.W.Redirect(p.signInURL(), p.getUsersUnlock))
	rt.POST("/users/unlock", p.W.Form(&fmEmail{}, p.postUsersUnlock))
	rt.POST("/users/forgot-password", p.W.Form(&fmEmail{}, p.postUsersForgotPassword))
	rt.POST("/users/reset-password", p.W.Form(&fmResetPassword{}, p.postUsersResetPassword))

	rt.DELETE("/users/sign-out", p.W.JSON(p.deleteUsersSignOut))
}

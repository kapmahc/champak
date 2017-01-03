package auth

import (
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount web points
func (p *Engine) Mount(rt *gin.Engine) {

	ung := rt.Group("/users")
	ung.GET("/sign-in", p.getUsersSignIn)
	ung.POST(
		"/sign-in",
		web.PostFormHandler("/users/sign-in", &fmSignIn{}, p.postUsersSignIn),
	)
	ung.GET("/sign-up", p.getUsersSignUp)
	ung.POST(
		"/sign-up",
		web.PostFormHandler("/users/sign-up", &fmSignUp{}, p.postUsersSignUp),
	)
	ung.GET("/confirm", p.getUsersConfirm)
	ung.GET(
		"/confirm/:token",
		web.FlashHandler("/users/sign-in", p.getUsersConfirmToken),
	)
	ung.POST(
		"/confirm",
		web.PostFormHandler("/users/confirm", &fmEmail{}, p.postUsersConfirm),
	)
	ung.GET("/unlock", p.getUsersUnlock)
	ung.GET(
		"/unlock/:token",
		web.FlashHandler("/users/sign-in", p.getUsersUnlockToken),
	)
	ung.POST("/unlock",
		web.PostFormHandler("/users/unlock", &fmEmail{}, p.postUsersUnlock),
	)
	ung.GET("/forgot-password", p.getUsersForgotPassword)
	ung.POST(
		"/forgot-password",
		web.PostFormHandler("/users/forgot-password", &fmEmail{}, p.postUsersForgotPassword),
	)
	ung.GET("/reset-password/:token", p.getUsersResetPassword)
	ung.POST(
		"/reset-password",
		web.PostFormHandler("/users/reset-password", &fmResetPassword{}, p.postUsersResetPassword),
	)

	umg := rt.Group("/users", p.Session.MustSignInHandler())
	umg.GET("/info", p.getUsersInfo)
	umg.POST(
		"/info",
		web.PostFormHandler("/users/info", &fmInfo{}, p.postUsersInfo),
	)
	umg.GET("/change-password", p.getUsersChangePassword)
	umg.POST(
		"/change-password",
		web.PostFormHandler("/users/change-password", &fmChangePassword{}, p.postUsersChangePassword),
	)
	umg.GET("/logs", p.getUsersLogs)
	umg.DELETE("/sign-out", p.deleteUsersSignOut)

	// rt.GET("/attachments/*name", p.getAttachment)
	// rt.POST("/attachments", p.Session.MustSignInHandler(), p.postAttachment)
	// rt.DELETE("/attachmetns/:id", p.Session.MustSignInHandler(), p.deleteAttachment)
	//
	// rt.POST("/votes", p.Session.MustSignInHandler(), p.postVotes)
}

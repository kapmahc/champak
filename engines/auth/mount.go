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
		web.PostFormHandler(&fmSignIn{}, p.postUsersSignIn),
	)
	ung.GET("/sign-up", p.getUsersSignUp)
	ung.POST(
		"/sign-up",
		web.PostFormHandler(&fmSignUp{}, p.postUsersSignUp),
	)
	ung.GET("/confirm", p.getUsersConfirm)
	ung.GET(
		"/confirm/:token",
		web.FlashHandler(p.getUsersConfirmToken),
	)
	ung.POST(
		"/confirm",
		web.PostFormHandler(&fmEmail{}, p.postUsersConfirm),
	)
	ung.GET("/unlock", p.getUsersUnlock)
	ung.GET(
		"/unlock/:token",
		web.FlashHandler(p.getUsersUnlockToken),
	)
	ung.POST("/unlock",
		web.PostFormHandler(&fmEmail{}, p.postUsersUnlock),
	)
	ung.GET("/forgot-password", p.getUsersForgotPassword)
	ung.POST(
		"/forgot-password",
		web.PostFormHandler(&fmEmail{}, p.postUsersForgotPassword),
	)
	ung.GET("/reset-password/:token", p.getUsersResetPassword)
	ung.POST(
		"/reset-password",
		web.PostFormHandler(&fmResetPassword{}, p.postUsersResetPassword),
	)

	umg := rt.Group("/users", p.Session.MustSignInHandler())
	umg.GET("/info", p.getUsersInfo)
	umg.POST(
		"/info",
		web.PostFormHandler(&fmInfo{}, p.postUsersInfo),
	)
	umg.GET("/change-password", p.getUsersChangePassword)
	umg.POST(
		"/change-password",
		web.PostFormHandler(&fmChangePassword{}, p.postUsersChangePassword),
	)
	umg.GET("/logs", p.getUsersLogs)
	umg.DELETE("/sign-out", p.deleteUsersSignOut)

	// rt.GET("/attachments/*name", p.getAttachment)
	// rt.POST("/attachments", p.Session.MustSignInHandler(), p.postAttachment)
	// rt.DELETE("/attachmetns/:id", p.Session.MustSignInHandler(), p.deleteAttachment)
	//
	// rt.POST("/votes", p.Session.MustSignInHandler(), p.postVotes)
}

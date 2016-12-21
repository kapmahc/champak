package auth

import (
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Mount mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	ug := rt.Group("/personal")
	ug.GET("/sign-in", p.getUsersSignIn)
	ug.POST(
		"/sign-in",
		web.PostFormHandler("/personal/sign-in", &fmSignIn{}, p.postUsersSignIn),
	)
	ug.GET("/sign-up", p.getUsersSignUp)
	ug.POST(
		"/sign-up",
		web.PostFormHandler("/personal/sign-up", &fmSignUp{}, p.postUsersSignUp),
	)
	ug.GET("/confirm", p.getUsersConfirm)
	ug.GET(
		"/confirm/:token",
		web.FlashHandler("/personal/sign-in", p.getUsersConfirmToken),
	)
	ug.POST(
		"/confirm",
		web.PostFormHandler("/personal/confirm", &fmEmail{}, p.postUsersConfirm),
	)
	ug.GET("/unlock", p.getUsersUnlock)
	ug.GET(
		"/unlock/:token",
		web.FlashHandler("/personal/sign-in", p.getUsersUnlockToken),
	)
	ug.POST("/unlock",
		web.PostFormHandler("/personal/unlock", &fmEmail{}, p.postUsersUnlock),
	)
	ug.GET("/forgot-password", p.getUsersForgotPassword)
	ug.POST(
		"/forgot-password",
		web.PostFormHandler("/personal/forgot-password", &fmEmail{}, p.postUsersForgotPassword),
	)
	ug.GET("/reset-password/:token", p.getUsersResetPassword)
	ug.POST(
		"/reset-password",
		web.PostFormHandler("/personal/reset-password", &fmResetPassword{}, p.postUsersResetPassword),
	)

	ug.GET("/profile", p.getUsersProfile)
	ug.POST("/profile", p.postUsersProfile)
	ug.GET("/change-password", p.getUsersChangePassword)
	ug.POST("/change-password", p.postUsersChangePassword)
	ug.GET("/logs", p.getUsersLogs)
	ug.DELETE("/sign-out", p.deleteUsersSignOut)

	rt.GET("/personal", p.getUsersSelf)
	rt.GET("/attachments/*name", p.getAttachment)
	rt.POST("/attachments", p.postAttachment)
	rt.DELETE("/attachmetns/:id", p.deleteAttachment)

	rt.POST("/votes", p.postVotes)
}

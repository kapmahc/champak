package auth

import gin "gopkg.in/gin-gonic/gin.v1"

// Mount mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	ung := rt.Group("/personal")
	ung.GET("/personal/sign-in", p.getSignIn)
	ung.POST("/personal/sign-in", p.postSignIn)
	ung.GET("/personal/sign-up", p.getSignUp)
	ung.POST("/personal/sign-up", p.postSignUp)
	ung.GET("/personal/confirm", p.getConfirm)
	ung.POST("/personal/confirm", p.postConfirm)
	ung.GET("/personal/unlock", p.getUnlock)
	ung.POST("/personal/unlock", p.postUnlock)
	ung.GET("/personal/forgot-passwrod", p.getForgotPassword)
	ung.POST("/personal/forgot-passwrod", p.postForgotPassword)
	ung.GET("/personal/change-password", p.getChangePassword)
	ung.POST("/personal/change-password", p.postChangePassword)

	umg := rt.Group("/personal", p.Jwt.CurrentUserHandler(true))
	umg.GET("/personal/profile", p.getProfile)
	umg.POST("/personal/profile", p.postProfile)

	rt.GET("/attachments/*name", p.getAttachment)
	rt.POST("/attachments", p.Jwt.CurrentUserHandler(true), p.postAttachment)
	rt.DELETE("/attachmetns/:id", p.Jwt.CurrentUserHandler(true), p.deleteAttachment)

	rt.POST("/votes", p.Jwt.CurrentUserHandler(true), p.postVotes)
}

package auth

import gin "gopkg.in/gin-gonic/gin.v1"

// Mount mount web points
func (p *Engine) Mount(rt *gin.Engine) {
	ung := rt.Group("/personal")
	ung.GET("/sign-in", p.getSignIn)
	ung.POST("/sign-in", p.postSignIn)
	ung.GET("/sign-up", p.getSignUp)
	ung.POST("/sign-up", p.postSignUp)
	ung.GET("/confirm", p.getConfirm)
	ung.POST("/confirm", p.postConfirm)
	ung.GET("/unlock", p.getUnlock)
	ung.POST("/unlock", p.postUnlock)
	ung.GET("/forgot-passwrod", p.getForgotPassword)
	ung.POST("/forgot-passwrod", p.postForgotPassword)
	ung.GET("/change-password", p.getChangePassword)
	ung.POST("/change-password", p.postChangePassword)

	umg := rt.Group("/personal", p.Jwt.CurrentUserHandler(true))
	umg.GET("/profile", p.getProfile)
	umg.POST("/profile", p.postProfile)
	umg.DELETE("/sign-out", p.deleteSignOut)

	rt.GET("/attachments/*name", p.getAttachment)
	rt.POST("/attachments", p.Jwt.CurrentUserHandler(true), p.postAttachment)
	rt.DELETE("/attachmetns/:id", p.Jwt.CurrentUserHandler(true), p.deleteAttachment)

	rt.POST("/votes", p.Jwt.CurrentUserHandler(true), p.postVotes)
}

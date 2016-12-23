package auth

import "github.com/gorilla/mux"

// Mount mount web points
func (p *Engine) Mount(rt *mux.Router) {
	// ung := rt.Group("/personal")
	// ung.GET("/sign-in", p.getUsersSignIn)
	// ung.POST(
	// 	"/sign-in",
	// 	web.PostFormHandler("/personal/sign-in", &fmSignIn{}, p.postUsersSignIn),
	// )
	// ung.GET("/sign-up", p.getUsersSignUp)
	// ung.POST(
	// 	"/sign-up",
	// 	web.PostFormHandler("/personal/sign-up", &fmSignUp{}, p.postUsersSignUp),
	// )
	// ung.GET("/confirm", p.getUsersConfirm)
	// ung.GET(
	// 	"/confirm/:token",
	// 	web.FlashHandler("/personal/sign-in", p.getUsersConfirmToken),
	// )
	// ung.POST(
	// 	"/confirm",
	// 	web.PostFormHandler("/personal/confirm", &fmEmail{}, p.postUsersConfirm),
	// )
	// ung.GET("/unlock", p.getUsersUnlock)
	// ung.GET(
	// 	"/unlock/:token",
	// 	web.FlashHandler("/personal/sign-in", p.getUsersUnlockToken),
	// )
	// ung.POST("/unlock",
	// 	web.PostFormHandler("/personal/unlock", &fmEmail{}, p.postUsersUnlock),
	// )
	// ung.GET("/forgot-password", p.getUsersForgotPassword)
	// ung.POST(
	// 	"/forgot-password",
	// 	web.PostFormHandler("/personal/forgot-password", &fmEmail{}, p.postUsersForgotPassword),
	// )
	// ung.GET("/reset-password/:token", p.getUsersResetPassword)
	// ung.POST(
	// 	"/reset-password",
	// 	web.PostFormHandler("/personal/reset-password", &fmResetPassword{}, p.postUsersResetPassword),
	// )
	//
	// umg := rt.Group("/personal", p.Session.MustSignInHandler())
	// umg.GET("/profile", p.getUsersProfile)
	// umg.POST(
	// 	"/profile",
	// 	web.PostFormHandler("/personal/profile", &fmProfile{}, p.postUsersProfile),
	// )
	// umg.GET("/change-password", p.getUsersChangePassword)
	// umg.POST(
	// 	"/change-password",
	// 	web.PostFormHandler("/personal/change-password", &fmChangePassword{}, p.postUsersChangePassword),
	// )
	// umg.GET("/logs", p.getUsersLogs)
	// umg.DELETE("/sign-out", p.deleteUsersSignOut)
	//
	// rt.GET("/attachments/*name", p.getAttachment)
	// rt.POST("/attachments", p.Session.MustSignInHandler(), p.postAttachment)
	// rt.DELETE("/attachmetns/:id", p.Session.MustSignInHandler(), p.deleteAttachment)
	//
	// rt.POST("/votes", p.Session.MustSignInHandler(), p.postVotes)
}

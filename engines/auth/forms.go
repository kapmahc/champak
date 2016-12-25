package auth

type fmSignIn struct {
	Email      string `form:"email" binding:"required"`
	Password   string `form:"password" binding:"required"`
	RememberMe bool   `form:"rememberMe"`
}

type fmSignUp struct {
	FullName             string `form:"fullName" binding:"required,max=255"`
	Email                string `form:"email" binding:"email"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

type fmEmail struct {
	Email string `form:"email" binding:"email"`
}

type fmResetPassword struct {
	Token                string `form:"token" binding:"required"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

type fmChangePassword struct {
	Password             string `form:"password" binding:"required"`
	NewPassword          string `form:"newPassword" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=NewPassword"`
}

type fmProfile struct {
	FullName string `form:"fullName" binding:"required,max=255"`
	Home     string `form:"home" binding:"required,max=255"`
	Logo     string `form:"logo" binding:"required,max=255"`
}

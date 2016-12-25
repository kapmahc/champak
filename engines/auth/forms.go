package auth

type fmSignIn struct {
	Email      string `form:"email" validate:"required"`
	Password   string `form:"password" validate:"required"`
	RememberMe bool   `form:"rememberMe"`
}

type fmSignUp struct {
	FullName             string `form:"fullName" validate:"required,max=255"`
	Email                string `form:"email" validate:"required,email"`
	Password             string `form:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

type fmEmail struct {
	Email string `form:"email" validate:"required,email"`
}

type fmResetPassword struct {
	Token                string `form:"token" validate:"required"`
	Password             string `form:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

type fmChangePassword struct {
	Password             string `form:"password" validate:"required"`
	NewPassword          string `form:"newPassword" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=NewPassword"`
}

type fmProfile struct {
	FullName string `form:"fullName" validate:"required,max=255"`
	Home     string `form:"home" validate:"required,max=255"`
	Logo     string `form:"logo" validate:"required,max=255"`
}

package site

type fmSiteInfo struct {
	Title       string `form:"title" validate:"required,max=255"`
	SubTitle    string `form:"subTitle" validate:"required,max=32"`
	Keywords    string `form:"keywords" validate:"required,max=64"`
	Description string `form:"description" validate:"required,max=500"`
	Copyright   string `form:"copyright" validate:"required,max=255"`
}

type fmSiteAuthor struct {
	Name  string `form:"name" validate:"required,max=32"`
	Email string `form:"email" validate:"required,email"`
}

type fmSiteSeo struct {
	Google string `form:"google" validate:"required,max=255"`
	Baidu  string `form:"baidu" validate:"required,max=255"`
}

type fmSiteSMTP struct {
	Host                 string `form:"host" validate:"required,max=255"`
	Port                 int    `form:"port"`
	User                 string `form:"user" validate:"required,max=255"`
	Password             string `form:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
	Ssl                  bool   `form:"ssl"`
}

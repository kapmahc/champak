package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getSignIn(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	title := p.I18n.T(lng, "auth.personal.sign-in.title")
	fm := web.NewForm(c, "post", "sign-in", title, "/personal/sign-in")
	fm.AddFields(
		web.NewTextField("fullName", p.I18n.T(lng, "attributes.full-name"), ""),
		web.NewEmailField("email", p.I18n.T(lng, "attributes.email"), ""),
		web.NewPasswordField("password", p.I18n.T(lng, "attributes.password")),
		web.NewPasswordField("passwordConfirmation", p.I18n.T(lng, "attributes.password-confirmation")),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/sign-in", data)
}
func (p *Engine) postSignIn(c *gin.Context) {
	ss := sessions.Default(c)
	ss.AddFlash("demo", web.ALERT)
	ss.Save()
	c.Redirect(http.StatusFound, "/personal/sign-in")
}

func (p *Engine) getSignUp(c *gin.Context) {

}
func (p *Engine) postSignUp(c *gin.Context) {

}

func (p *Engine) getConfirm(c *gin.Context) {

}
func (p *Engine) postConfirm(c *gin.Context) {

}

func (p *Engine) getForgotPassword(c *gin.Context) {

}
func (p *Engine) postForgotPassword(c *gin.Context) {

}

func (p *Engine) getChangePassword(c *gin.Context) {

}
func (p *Engine) postChangePassword(c *gin.Context) {

}

func (p *Engine) getUnlock(c *gin.Context) {

}
func (p *Engine) postUnlock(c *gin.Context) {

}

func (p *Engine) getProfile(c *gin.Context) {

}
func (p *Engine) postProfile(c *gin.Context) {

}

func (p *Engine) deleteSignOut(c *gin.Context) {

}

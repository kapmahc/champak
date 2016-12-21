package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getUsersChangePassword(c *gin.Context) {

}
func (p *Engine) postUsersChangePassword(c *gin.Context) {

}
func (p *Engine) getUsersLogs(c *gin.Context) {

}

func (p *Engine) getUsersProfile(c *gin.Context) {
	user := c.MustGet(CurrentUser).(*User)
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "auth.personal.profile.title")
	fm := web.NewForm(c, "profile", title, "/personal/profile")
	fm.AddFields(
		web.NewTextField("fullName", p.I18n.T(lng, "attributes.full-name"), user.FullName),
		web.NewEmailField("email", p.I18n.T(lng, "attributes.email"), user.Email),
		web.NewTextField("logo", p.I18n.T(lng, "auth.attributes.user.logo"), user.Logo),
		web.NewTextField("home", p.I18n.T(lng, "auth.attributes.user.home"), user.Home),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/form", data)
}
func (p *Engine) postUsersProfile(c *gin.Context) {

}

func (p *Engine) deleteUsersSignOut(c *gin.Context) {
	user := c.MustGet(CurrentUser).(*User)
	ss := sessions.Default(c)
	ss.Clear()
	ss.Save()
	lng := c.MustGet(web.LOCALE).(string)
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lng, "auth.logs.sign-out"))
	c.JSON(http.StatusOK, gin.H{"to": "/"})
}

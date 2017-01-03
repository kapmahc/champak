package auth

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getUsersChangePassword(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "auth.users.change-password.title")
	fm := web.NewForm(c, "change-password", title, "/users/change-password")
	fm.AddFields(
		web.NewPasswordField("password", p.I18n.T(lng, "attributes.password")),
		web.NewPasswordField("newPassword", p.I18n.T(lng, "attributes.newPassword")),
		web.NewPasswordField("passwordConfirmation", p.I18n.T(lng, "attributes.passwordConfirmation")),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/form", data)
}

type fmChangePassword struct {
	Password             string `form:"password" binding:"required"`
	NewPassword          string `form:"newPassword" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=NewPassword"`
}

func (p *Engine) postUsersChangePassword(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	user := c.MustGet(CurrentUser).(*User)
	fm := o.(*fmChangePassword)
	if !p.Security.Chk([]byte(fm.Password), user.Password) {
		return p.I18n.E(lng, "auth.errors.bad-password")
	}

	if err := p.Db.
		Model(user).
		Update("password", p.Security.Sum([]byte(fm.NewPassword))).Error; err != nil {
		return err
	}
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lng, "auth.logs.change-password"))
	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "auth.messages.change-password-success"), web.NOTICE)
	ss.Save()
	c.Redirect(http.StatusFound, "/users/change-password")
	return nil
}

func (p *Engine) getUsersLogs(c *gin.Context) {
	user := c.MustGet(CurrentUser).(*User)
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	if err := p.Db.Model(user).
		Order("id DESC").Limit(60).
		Related(&user.Logs).
		Error; err != nil {
		log.Error(err)
	}
	data["title"] = p.I18n.T(lng, "auth.users.logs.title")
	data["logs"] = user.Logs
	c.HTML(http.StatusOK, "auth/users/logs", data)
}

func (p *Engine) getUsersProfile(c *gin.Context) {
	user := c.MustGet(CurrentUser).(*User)
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)

	title := p.I18n.T(lng, "auth.users.profile.title")
	fm := web.NewForm(c, "profile", title, "/users/profile")
	email := web.NewEmailField("email", p.I18n.T(lng, "attributes.email"), user.Email)
	email.Require = false
	email.ReadOnly = true
	fm.AddFields(
		web.NewTextField("fullName", p.I18n.T(lng, "attributes.fullName"), user.FullName),
		email,
		web.NewTextField("logo", p.I18n.T(lng, "auth.attributes.user.logo"), user.Logo),
		web.NewTextField("home", p.I18n.T(lng, "auth.attributes.user.home"), user.Home),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/form", data)
}

type fmProfile struct {
	FullName string `form:"fullName" binding:"required,max=255"`
	Home     string `form:"home" binding:"required,max=255"`
	Logo     string `form:"logo" binding:"required,max=255"`
}

func (p *Engine) postUsersProfile(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	user := c.MustGet(CurrentUser).(*User)
	fm := o.(*fmProfile)

	if err := p.Db.
		Model(user).
		Updates(map[string]interface{}{
			"full_name": fm.FullName,
			"home":      fm.Home,
			"logo":      fm.Logo,
		}).Error; err != nil {
		return err
	}
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lng, "auth.logs.update-profile"))
	c.Redirect(http.StatusFound, "/users/profile")
	return nil
}

func (p *Engine) deleteUsersSignOut(c *gin.Context) {
	user := c.MustGet(CurrentUser).(*User)
	ss := sessions.Default(c)
	ss.Clear()
	ss.Save()
	lng := c.MustGet(web.LOCALE).(string)
	p.Dao.Log(user.ID, c.ClientIP(), p.I18n.T(lng, "auth.logs.sign-out"))
	c.JSON(http.StatusOK, gin.H{web.TO: "/"})
}

package auth

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

const (
	actConfirm       = "confirm"
	actUnlock        = "unlock"
	actResetPassword = "reset-password"
)

func (p *Engine) getUsersSignIn(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	title := p.I18n.T(lng, "auth.personal.sign-in.title")
	fm := web.NewForm(c, "sign-in", title, "/personal/sign-in")
	fm.AddFields(
		web.NewEmailField("email", p.I18n.T(lng, "attributes.email"), ""),
		web.NewPasswordField("password", p.I18n.T(lng, "attributes.password")),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/non-sign-in", data)
}

type fmSignIn struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (p *Engine) postUsersSignIn(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmSignIn)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}
	ip := c.ClientIP()
	if !p.Security.Chk([]byte(fm.Password), user.Password) {
		p.Dao.Log(user.ID, p.I18n.T(lng, "auth.logs.sign-in-failed", ip))
		return p.I18n.E(lng, "auth.errors.email-password-not-match")
	}
	if !user.IsConfirm() {
		return p.I18n.E(lng, "auth.errors.user-not-confirm")
	}
	if user.IsLock() {
		return p.I18n.E(lng, "auth.errors.user-is-lock")
	}

	p.Dao.signIn(user.ID, ip)
	p.Dao.Log(user.ID, p.I18n.T(lng, "auth.logs.sign-in-success", ip))
	ss := sessions.Default(c)
	ss.Set("uid", user.UID)
	ss.Save()
	c.Redirect(http.StatusFound, "/")
	return nil
}

func (p *Engine) getUsersSignUp(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	title := p.I18n.T(lng, "auth.personal.sign-up.title")
	fm := web.NewForm(c, "sign-up", title, "/personal/sign-up")
	fm.AddFields(
		web.NewTextField("fullName", p.I18n.T(lng, "attributes.full-name"), ""),
		web.NewEmailField("email", p.I18n.T(lng, "attributes.email"), ""),
		web.NewPasswordField("password", p.I18n.T(lng, "attributes.password")),
		web.NewPasswordField("passwordConfirmation", p.I18n.T(lng, "attributes.password-confirmation")),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/non-sign-in", data)
}

type fmSignUp struct {
	FullName             string `form:"fullName" binding:"min=2,max=255"`
	Email                string `form:"email" binding:"email"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) postUsersSignUp(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmSignUp)

	var count int
	if err := p.Db.
		Model(&User{}).
		Where("email = ?", fm.Email).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return p.I18n.E(lng, "auth.errors.email-already-exists")
	}

	user, err := p.Dao.AddEmailUser(fm.FullName, fm.Email, fm.Password)
	if err != nil {
		return err
	}
	p.Dao.Log(user.ID, p.I18n.T(lng, "auth.logs.sign-up"))
	if err = p.sendEmail(lng, user, actConfirm); err != nil {
		log.Error(err)
	}

	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "auth.messages.email-for-confirm"), web.NOTICE)
	ss.Save()
	c.Redirect(http.StatusFound, "/personal/sign-in")
	return nil
}

func (p *Engine) getUsersConfirmToken(c *gin.Context) error {
	lng := c.MustGet(web.LOCALE).(string)
	user, err := p.parseToken(c.Param("token"), actConfirm)
	if err != nil {
		return err
	}
	if user.IsConfirm() {
		return p.I18n.E(lng, "auth.errors.user-already-confirm")
	}
	if err = p.Db.Model(user).Update("confirm_at", time.Now()).Error; err != nil {
		return err
	}
	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "auth.messages.confirm-success"), web.NOTICE)
	ss.Save()
	c.Redirect(http.StatusFound, "personal/sign-in")
	return nil
}

func (p *Engine) getUsersConfirm(c *gin.Context) {

	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	title := p.I18n.T(lng, "auth.personal.confirm.title")
	fm := web.NewForm(c, "confirm", title, "/personal/confirm")
	fm.AddFields(
		web.NewEmailField("email", p.I18n.T(lng, "attributes.email"), ""),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/non-sign-in", data)
}

type fmEmail struct {
	Email string `form:"email" binding:"email"`
}

func (p *Engine) postUsersConfirm(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmEmail)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}
	if user.IsConfirm() {
		return p.I18n.E(lng, "auth.errors.user-already-confirm")
	}

	if err = p.sendEmail(lng, user, actConfirm); err != nil {
		log.Error(err)
	}
	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "auth.messages.email-for-confirm"), web.NOTICE)
	ss.Save()
	c.Redirect(http.StatusFound, "personal/sign-in")
	return nil
}

func (p *Engine) getUsersForgotPassword(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	title := p.I18n.T(lng, "auth.personal.forgot-password.title")
	fm := web.NewForm(c, "forgot-password", title, "/personal/forgot-password")
	fm.AddFields(
		web.NewEmailField("email", p.I18n.T(lng, "attributes.email"), ""),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/non-sign-in", data)
}

func (p *Engine) postUsersForgotPassword(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmEmail)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}
	if err = p.sendEmail(lng, user, actResetPassword); err != nil {
		log.Error(err)
	}
	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "auth.messages.email-for-reset-password"), web.NOTICE)
	ss.Save()
	c.Redirect(http.StatusFound, "personal/sign-in")
	return nil
}

func (p *Engine) getUsersResetPassword(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	title := p.I18n.T(lng, "auth.personal.reset-password.title")
	fm := web.NewForm(c, "reset-password", title, "/personal/reset-password")
	fm.AddFields(
		web.NewHiddenField("token", c.Param("token")),
		web.NewPasswordField("password", p.I18n.T(lng, "attributes.password")),
		web.NewPasswordField("passwordConfirmation", p.I18n.T(lng, "attributes.password-confirmation")),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/non-sign-in", data)
}

type fmResetPassword struct {
	Token                string `form:"token" binding:"required"`
	Password             string `form:"password" binding:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" binding:"eqfield=Password"`
}

func (p *Engine) postUsersResetPassword(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmResetPassword)
	user, err := p.parseToken(fm.Token, actResetPassword)
	if err != nil {
		return err
	}
	if err = p.Db.Model(user).
		Update("password", p.Security.Sum([]byte(fm.Password))).Error; err != nil {
		return err
	}
	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "auth.messages.reset-password-success"), web.NOTICE)
	ss.Save()
	c.Redirect(http.StatusFound, "personal/sign-in")
	return nil
}

func (p *Engine) getUsersUnlockToken(c *gin.Context) error {
	lng := c.MustGet(web.LOCALE).(string)
	user, err := p.parseToken(c.Param("token"), actConfirm)
	if err != nil {
		return err
	}
	if !user.IsLock() {
		return p.I18n.E(lng, "auth.errors.user-not-lock")
	}
	if err = p.Db.Model(user).Update(map[string]interface{}{"locked_at": nil}).Error; err != nil {
		return err
	}
	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "auth.messages.unlock-success"), web.NOTICE)
	ss.Save()
	c.Redirect(http.StatusFound, "personal/sign-in")
	return nil
}

func (p *Engine) getUsersUnlock(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	title := p.I18n.T(lng, "auth.personal.unlock.title")
	fm := web.NewForm(c, "unlock", title, "/personal/unlock")
	fm.AddFields(
		web.NewEmailField("email", p.I18n.T(lng, "attributes.email"), ""),
	)

	data["title"] = title
	data["form"] = fm
	c.HTML(http.StatusOK, "auth/non-sign-in", data)
}
func (p *Engine) postUsersUnlock(c *gin.Context, o interface{}) error {
	lng := c.MustGet(web.LOCALE).(string)
	fm := o.(*fmEmail)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}
	if !user.IsLock() {
		return p.I18n.E(lng, "auth.errors.user-not-lock")
	}

	if err = p.sendEmail(lng, user, actUnlock); err != nil {
		log.Error(err)
	}
	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, "auth.messages.email-for-unlock"), web.NOTICE)
	ss.Save()
	c.Redirect(http.StatusFound, "personal/sign-in")
	return nil
}

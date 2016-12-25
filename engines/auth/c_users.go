package auth

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	sessions "github.com/goincremental/negroni-sessions"
	"github.com/kapmahc/champak/web"
	"github.com/kapmahc/champak/web/crypto"
	"github.com/kapmahc/champak/web/i18n"
	"github.com/kapmahc/champak/web/mux"
)

const (
	actConfirm       = "confirm"
	actUnlock        = "unlock"
	actResetPassword = "reset-password"
)

func (p *Engine) getUsersSignIn(wrt http.ResponseWriter, req *http.Request) (web.H, error) {
	lng := req.Context().Value(web.LOCALE).(string)
	data := req.Context().Value(web.DATA).(web.H)
	title := i18n.T(lng, "auth.users.sign-in.title")
	fm := web.NewForm(req, "sign-in", title, mux.URL("auth.users.sign-in"))
	fm.AddFields(
		web.NewEmailField("email", i18n.T(lng, "attributes.email"), ""),
		web.NewPasswordField("password", i18n.T(lng, "attributes.password")),
	)

	data["title"] = title
	data["form"] = fm

	return data, nil
}

func (p *Engine) postUsersSignIn(wrt http.ResponseWriter, req *http.Request, o interface{}) error {
	lng := req.Context().Value(web.LOCALE).(string)
	fm := o.(*fmSignIn)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}
	ip := req.Context().Value(web.ClientIP).(string)
	if !crypto.Chk([]byte(fm.Password), user.Password) {
		p.Dao.Log(user.ID, ip, i18n.T(lng, "auth.logs.sign-in-failed"))
		return i18n.E(lng, "auth.errors.email-password-not-match")
	}
	if !user.IsConfirm() {
		return i18n.E(lng, "auth.errors.user-not-confirm")
	}
	if user.IsLock() {
		return i18n.E(lng, "auth.errors.user-is-lock")
	}

	p.Dao.signIn(user.ID, ip)
	p.Dao.Log(user.ID, ip, i18n.T(lng, "auth.logs.sign-in-success"))
	ss := sessions.GetSession(req)
	ss.Set("uid", user.UID)
	return nil
}

func (p *Engine) getUsersSignUp(wrt http.ResponseWriter, req *http.Request) (web.H, error) {
	lng := req.Context().Value(web.LOCALE).(string)
	data := req.Context().Value(web.DATA).(web.H)
	title := i18n.T(lng, "auth.users.sign-up.title")
	fm := web.NewForm(req, "sign-up", title, mux.URL("auth.users.sign-up"))
	fm.AddFields(
		web.NewTextField("fullName", i18n.T(lng, "attributes.fullName"), ""),
		web.NewEmailField("email", i18n.T(lng, "attributes.email"), ""),
		web.NewPasswordField("password", i18n.T(lng, "attributes.password")),
		web.NewPasswordField("passwordConfirmation", i18n.T(lng, "attributes.passwordConfirmation")),
	)

	data["title"] = title
	data["form"] = fm
	return data, nil
}

func (p *Engine) postUsersSignUp(wrt http.ResponseWriter, req *http.Request, o interface{}) error {
	lng := req.Context().Value(web.LOCALE).(string)
	fm := o.(*fmSignUp)

	var count int
	if err := p.Db.
		Model(&User{}).
		Where("email = ?", fm.Email).
		Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return i18n.E(lng, "auth.errors.email-already-exists")
	}

	user, err := p.Dao.AddEmailUser(fm.FullName, fm.Email, fm.Password)
	if err != nil {
		return err
	}
	ip := req.Context().Value(web.ClientIP).(string)
	p.Dao.Log(user.ID, ip, i18n.T(lng, "auth.logs.sign-up"))
	p.sendEmail(lng, user, actConfirm)

	ss := sessions.GetSession(req)
	ss.AddFlash(i18n.T(lng, "auth.messages.email-for-confirm"), web.NOTICE)

	return nil
}

func (p *Engine) getUsersConfirmToken(wrt http.ResponseWriter, req *http.Request) error {
	lng := req.Context().Value(web.LOCALE).(string)
	user, err := p.parseToken(mux.Param(req, "token"), actConfirm)
	if err != nil {
		return err
	}
	if user.IsConfirm() {
		return i18n.E(lng, "auth.errors.user-already-confirm")
	}
	if err = p.Db.Model(user).Update("confirmed_at", time.Now()).Error; err != nil {
		return err
	}
	p.Dao.Log(user.ID, req.Context().Value(web.ClientIP).(string), i18n.T(lng, "auth.logs.confirm"))
	ss := sessions.GetSession(req)
	ss.AddFlash(i18n.T(lng, "auth.messages.confirm-success"), web.NOTICE)
	return nil
}

func (p *Engine) getUsersConfirm(wrt http.ResponseWriter, req *http.Request) (web.H, error) {

	lng := req.Context().Value(web.LOCALE).(string)
	data := req.Context().Value(web.DATA).(web.H)
	title := i18n.T(lng, "auth.users.confirm.title")
	fm := web.NewForm(req, "confirm", title, mux.URL("auth.users.confirm"))
	fm.AddFields(
		web.NewEmailField("email", i18n.T(lng, "attributes.email"), ""),
	)

	data["title"] = title
	data["form"] = fm
	return data, nil
}

func (p *Engine) postUsersConfirm(wrt http.ResponseWriter, req *http.Request, o interface{}) error {
	lng := req.Context().Value(web.LOCALE).(string)
	fm := o.(*fmEmail)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}
	if user.IsConfirm() {
		return i18n.E(lng, "auth.errors.user-already-confirm")
	}

	p.sendEmail(lng, user, actConfirm)
	ss := sessions.GetSession(req)
	ss.AddFlash(i18n.T(lng, "auth.messages.email-for-confirm"), web.NOTICE)

	return nil
}

func (p *Engine) getUsersForgotPassword(wrt http.ResponseWriter, req *http.Request) (web.H, error) {
	lng := req.Context().Value(web.LOCALE).(string)
	data := req.Context().Value(web.DATA).(web.H)
	title := i18n.T(lng, "auth.users.forgot-password.title")
	fm := web.NewForm(req, "forgot-password", title, mux.URL("auth.users.forgot-password"))
	fm.AddFields(
		web.NewEmailField("email", i18n.T(lng, "attributes.email"), ""),
	)

	data["title"] = title
	data["form"] = fm
	return data, nil
}

func (p *Engine) postUsersForgotPassword(wrt http.ResponseWriter, req *http.Request, o interface{}) error {
	lng := req.Context().Value(web.LOCALE).(string)
	fm := o.(*fmEmail)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}
	p.sendEmail(lng, user, actResetPassword)
	ss := sessions.GetSession(req)
	ss.AddFlash(i18n.T(lng, "auth.messages.email-for-reset-password"), web.NOTICE)

	return nil
}

func (p *Engine) getUsersResetPassword(wrt http.ResponseWriter, req *http.Request) (web.H, error) {
	lng := req.Context().Value(web.LOCALE).(string)
	data := req.Context().Value(web.DATA).(web.H)
	title := i18n.T(lng, "auth.users.reset-password.title")
	fm := web.NewForm(req, "reset-password", title, mux.URL("auth.users.reset-password"))
	fm.AddFields(
		web.NewHiddenField("token", mux.Param(req, "token")),
		web.NewPasswordField("password", i18n.T(lng, "attributes.password")),
		web.NewPasswordField("passwordConfirmation", i18n.T(lng, "attributes.passwordConfirmation")),
	)

	data["title"] = title
	data["form"] = fm
	return data, nil
}

func (p *Engine) postUsersResetPassword(wrt http.ResponseWriter, req *http.Request, o interface{}) error {
	lng := req.Context().Value(web.LOCALE).(string)
	fm := o.(*fmResetPassword)
	user, err := p.parseToken(fm.Token, actResetPassword)
	if err != nil {
		return err
	}
	if err = p.Db.Model(user).
		Update("password", crypto.Sum([]byte(fm.Password))).Error; err != nil {
		return err
	}
	p.Dao.Log(user.ID, req.Context().Value(web.ClientIP).(string), i18n.T(lng, "auth.logs.reset-password"))
	ss := sessions.GetSession(req)
	ss.AddFlash(i18n.T(lng, "auth.messages.reset-password-success"), web.NOTICE)

	return nil
}

func (p *Engine) getUsersUnlockToken(wrt http.ResponseWriter, req *http.Request) error {
	lng := req.Context().Value(web.LOCALE).(string)
	user, err := p.parseToken(mux.Param(req, "token"), actUnlock)
	if err != nil {
		return err
	}
	if !user.IsLock() {
		return i18n.E(lng, "auth.errors.user-not-lock")
	}
	if err = p.Db.Model(user).Update(map[string]interface{}{"locked_at": nil}).Error; err != nil {
		return err
	}
	p.Dao.Log(user.ID, req.Context().Value(web.ClientIP).(string), i18n.T(lng, "auth.logs.unlock"))
	ss := sessions.GetSession(req)
	ss.AddFlash(i18n.T(lng, "auth.messages.unlock-success"), web.NOTICE)

	return nil
}

func (p *Engine) getUsersUnlock(wrt http.ResponseWriter, req *http.Request) (web.H, error) {
	lng := req.Context().Value(web.LOCALE).(string)
	data := req.Context().Value(web.DATA).(web.H)
	title := i18n.T(lng, "auth.users.unlock.title")
	fm := web.NewForm(req, "unlock", title, mux.URL("auth.users.unlock"))
	fm.AddFields(
		web.NewEmailField("email", i18n.T(lng, "attributes.email"), ""),
	)

	data["title"] = title
	data["form"] = fm
	return data, nil
}

func (p *Engine) postUsersUnlock(wrt http.ResponseWriter, req *http.Request, o interface{}) error {
	lng := req.Context().Value(web.LOCALE).(string)
	fm := o.(*fmEmail)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return err
	}
	if !user.IsLock() {
		return i18n.E(lng, "auth.errors.user-not-lock")
	}

	p.sendEmail(lng, user, actUnlock)
	ss := sessions.GetSession(req)
	ss.AddFlash(i18n.T(lng, "auth.messages.email-for-unlock"), web.NOTICE)

	return nil
}

func (p *Engine) getUsersChangePassword(wrt http.ResponseWriter, req *http.Request) (web.H, error) {
	lng := req.Context().Value(web.LOCALE).(string)
	data := req.Context().Value(web.DATA).(web.H)

	title := i18n.T(lng, "auth.users.change-password.title")
	fm := web.NewForm(req, "change-password", title, mux.URL("auth.users.change-password"))
	fm.AddFields(
		web.NewPasswordField("password", i18n.T(lng, "attributes.password")),
		web.NewPasswordField("newPassword", i18n.T(lng, "attributes.newPassword")),
		web.NewPasswordField("passwordConfirmation", i18n.T(lng, "attributes.passwordConfirmation")),
	)

	data["title"] = title
	data["form"] = fm
	return data, nil
}

func (p *Engine) postUsersChangePassword(wrt http.ResponseWriter, req *http.Request, o interface{}) error {
	lng := req.Context().Value(web.LOCALE).(string)
	user := req.Context().Value(CurrentUser).(*User)
	fm := o.(*fmChangePassword)
	if !crypto.Chk([]byte(fm.Password), user.Password) {
		return i18n.E(lng, "auth.errors.bad-password")
	}

	if err := p.Db.
		Model(user).
		Update("password", crypto.Sum([]byte(fm.NewPassword))).Error; err != nil {
		return err
	}
	p.Dao.Log(user.ID, req.Context().Value(web.ClientIP).(string), i18n.T(lng, "auth.logs.change-password"))
	ss := sessions.GetSession(req)
	ss.AddFlash(i18n.T(lng, "auth.messages.change-password-success"), web.NOTICE)

	return nil
}

func (p *Engine) getUsersLogs(wrt http.ResponseWriter, req *http.Request) (web.H, error) {
	user := req.Context().Value(CurrentUser).(*User)
	lng := req.Context().Value(web.LOCALE).(string)
	data := req.Context().Value(web.DATA).(web.H)
	if err := p.Db.Model(user).
		Order("id DESC").Limit(60).
		Related(&user.Logs).
		Error; err != nil {
		log.Error(err)
	}
	data["title"] = i18n.T(lng, "auth.users.logs.title")
	data["logs"] = user.Logs
	return data, nil
}

func (p *Engine) getUsersProfile(wrt http.ResponseWriter, req *http.Request) (web.H, error) {
	user := req.Context().Value(CurrentUser).(*User)
	lng := req.Context().Value(web.LOCALE).(string)
	data := req.Context().Value(web.DATA).(web.H)

	title := i18n.T(lng, "auth.users.profile.title")
	fm := web.NewForm(req, "profile", title, mux.URL("auth.users.profile"))
	email := web.NewEmailField("email", i18n.T(lng, "attributes.email"), user.Email)
	email.Require = false
	email.ReadOnly = true
	fm.AddFields(
		web.NewTextField("fullName", i18n.T(lng, "attributes.fullName"), user.FullName),
		email,
		web.NewTextField("logo", i18n.T(lng, "auth.attributes.user.logo"), user.Logo),
		web.NewTextField("home", i18n.T(lng, "auth.attributes.user.home"), user.Home),
	)

	data["title"] = title
	data["form"] = fm
	return data, nil
}

func (p *Engine) postUsersProfile(wrt http.ResponseWriter, req *http.Request, o interface{}) error {
	lng := req.Context().Value(web.LOCALE).(string)
	user := req.Context().Value(CurrentUser).(*User)
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
	p.Dao.Log(user.ID, req.Context().Value(web.ClientIP).(string), i18n.T(lng, "auth.logs.update-profile"))
	return nil
}

func (p *Engine) deleteUsersSignOut(wrt http.ResponseWriter, req *http.Request) (web.H, error) {
	user := req.Context().Value(CurrentUser).(*User)
	ss := sessions.GetSession(req)
	ss.Clear()

	lng := req.Context().Value(web.LOCALE).(string)
	p.Dao.Log(user.ID, req.Context().Value(web.ClientIP).(string), i18n.T(lng, "auth.logs.sign-out"))
	return web.H{web.TO: "/"}, nil
}

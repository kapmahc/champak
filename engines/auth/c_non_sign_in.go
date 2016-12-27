package auth

import (
	"net/http"
	"time"

	"github.com/SermoDigital/jose/jws"
	"github.com/julienschmidt/httprouter"
	"github.com/kapmahc/champak/web"
)

const (
	actConfirm       = "confirm"
	actUnlock        = "unlock"
	actResetPassword = "reset-password"
)

func (p *Engine) postUsersSignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params, o interface{}) (interface{}, error) {
	lng := r.Context().Value(web.LOCALE).(string)
	fm := o.(*fmSignUp)

	var count int
	if err := p.Db.
		Model(&User{}).
		Where("email = ?", fm.Email).
		Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, p.I18n.E(lng, "auth.errors.email-already-exists")
	}

	user, err := p.Dao.AddEmailUser(fm.FullName, fm.Email, fm.Password)
	if err != nil {
		return nil, err
	}
	p.Dao.Log(user.ID, p.W.ClientIP(r), p.I18n.T(lng, "auth.logs.sign-up"))
	p.sendEmail(lng, user, actConfirm)

	return web.H{"message": p.I18n.T(lng, "auth.messages.email-for-confirm")}, nil
}

func (p *Engine) postUsersSignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params, o interface{}) (interface{}, error) {
	lng := r.Context().Value(web.LOCALE).(string)
	fm := o.(*fmSignIn)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	ip := p.W.ClientIP(r)
	if !p.Security.Chk([]byte(fm.Password), user.Password) {
		p.Dao.Log(user.ID, ip, p.I18n.T(lng, "auth.logs.sign-in-failed"))
		return nil, p.I18n.E(lng, "auth.errors.email-password-not-match")
	}
	if !user.IsConfirm() {
		return nil, p.I18n.E(lng, "auth.errors.user-not-confirm")
	}
	if user.IsLock() {
		return nil, p.I18n.E(lng, "auth.errors.user-is-lock")
	}

	p.Dao.signIn(user.ID, ip)
	p.Dao.Log(user.ID, ip, p.I18n.T(lng, "auth.logs.sign-in-success"))

	var cm jws.Claims
	cm.Set("name", user.FullName)
	cm.Set("uid", user.UID)
	tkn, err := p.Jwt.Sum(cm, 7)
	return web.H{"token": tkn}, err
}

func (p *Engine) getUsersConfirm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	lng := r.Context().Value(web.LOCALE).(string)
	user, err := p.parseToken(lng, ps.ByName("token"), actConfirm)
	if err != nil {
		return nil, err
	}
	if user.IsConfirm() {
		return nil, p.I18n.E(lng, "auth.errors.user-already-confirm")
	}
	if err = p.Db.Model(user).Update("confirmed_at", time.Now()).Error; err != nil {
		return nil, err
	}
	p.Dao.Log(user.ID, p.W.ClientIP(r), p.I18n.T(lng, "auth.logs.confirm"))

	return web.H{"message": p.I18n.T(lng, "auth.messages.confirm-success")}, nil
}

func (p *Engine) postUsersConfirm(w http.ResponseWriter, r *http.Request, _ httprouter.Params, o interface{}) (interface{}, error) {
	lng := r.Context().Value(web.LOCALE).(string)
	fm := o.(*fmEmail)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	if user.IsConfirm() {
		return nil, p.I18n.E(lng, "auth.errors.user-already-confirm")
	}

	p.sendEmail(lng, user, actConfirm)
	return web.H{"message": p.I18n.T(lng, "auth.messages.email-for-confirm")}, nil
}

func (p *Engine) postUsersForgotPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params, o interface{}) (interface{}, error) {
	lng := r.Context().Value(web.LOCALE).(string)
	fm := o.(*fmEmail)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	p.sendEmail(lng, user, actResetPassword)

	return web.H{"message": p.I18n.T(lng, "auth.messages.email-for-reset-password")}, nil
}

func (p *Engine) postUsersResetPassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params, o interface{}) (interface{}, error) {
	lng := r.Context().Value(web.LOCALE).(string)
	fm := o.(*fmResetPassword)
	user, err := p.parseToken(lng, fm.Token, actResetPassword)
	if err != nil {
		return nil, err
	}
	if err = p.Db.Model(user).
		Update("password", p.Security.Sum([]byte(fm.Password))).Error; err != nil {
		return nil, err
	}
	p.Dao.Log(user.ID, p.W.ClientIP(r), p.I18n.T(lng, "auth.logs.reset-password"))

	return web.H{"message": p.I18n.T(lng, "auth.messages.reset-password-success")}, nil
}

func (p *Engine) getUsersUnlock(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (interface{}, error) {
	lng := r.Context().Value(web.LOCALE).(string)
	user, err := p.parseToken(lng, ps.ByName("token"), actConfirm)
	if err != nil {
		return nil, err
	}
	if !user.IsLock() {
		return nil, p.I18n.E(lng, "auth.errors.user-not-lock")
	}
	if err = p.Db.Model(user).Update(map[string]interface{}{"locked_at": nil}).Error; err != nil {
		return nil, err
	}
	p.Dao.Log(user.ID, p.W.ClientIP(r), p.I18n.T(lng, "auth.logs.unlock"))

	return web.H{"message": p.I18n.T(lng, "auth.messages.unlock-success")}, nil
}

func (p *Engine) postUsersUnlock(w http.ResponseWriter, r *http.Request, _ httprouter.Params, o interface{}) (interface{}, error) {
	lng := r.Context().Value(web.LOCALE).(string)
	fm := o.(*fmEmail)
	user, err := p.Dao.GetByEmail(fm.Email)
	if err != nil {
		return nil, err
	}
	if !user.IsLock() {
		return nil, p.I18n.E(lng, "auth.errors.user-not-lock")
	}

	p.sendEmail(lng, user, actUnlock)

	return web.H{"message": p.I18n.T(lng, "auth.messages.email-for-unlock")}, nil
}

// -----------------------------------------------------------------------------

func (p *Engine) parseToken(lng, tkn, act string) (*User, error) {
	cm, err := p.Jwt.Validate([]byte(tkn))
	if err != nil {
		return nil, err
	}
	if act != cm.Get("act").(string) {
		return nil, p.I18n.E(lng, "auth.errors.bad-action")
	}
	return p.Dao.GetUserByUID(cm.Get("uid").(string))
}

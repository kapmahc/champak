package auth

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kapmahc/champak/web"
)

func (p *Engine) getUsersLogs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (interface{}, error) {
	user := r.Context().Value(CurrentUser).(*User)
	err := p.Db.Model(user).
		Order("id DESC").Limit(60).
		Related(&user.Logs).
		Error

	return user.Logs, err
}

func (p *Engine) getUsersInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (interface{}, error) {
	user := r.Context().Value(CurrentUser).(*User)
	return user, nil
}

func (p *Engine) postUsersInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params, o interface{}) (interface{}, error) {
	lng := r.Context().Value(web.LOCALE).(string)
	user := r.Context().Value(CurrentUser).(*User)
	fm := o.(*fmUserInfo)

	if err := p.Db.
		Model(user).
		Updates(map[string]interface{}{
			"full_name": fm.FullName,
			"home":      fm.Home,
			"logo":      fm.Logo,
		}).Error; err != nil {
		return nil, err
	}
	p.Dao.Log(user.ID, p.W.ClientIP(r), p.I18n.T(lng, "auth.logs.update-profile"))
	return web.H{}, nil
}

func (p *Engine) postUsersChangePassword(w http.ResponseWriter, r *http.Request, _ httprouter.Params, o interface{}) (interface{}, error) {
	lng := r.Context().Value(web.LOCALE).(string)
	user := r.Context().Value(CurrentUser).(*User)
	fm := o.(*fmChangePassword)
	if !p.Security.Chk([]byte(fm.Password), user.Password) {
		return nil, p.I18n.E(lng, "auth.errors.bad-password")
	}

	if err := p.Db.
		Model(user).
		Update("password", p.Security.Sum([]byte(fm.NewPassword))).Error; err != nil {
		return nil, err
	}
	p.Dao.Log(user.ID, p.W.ClientIP(r), p.I18n.T(lng, "auth.logs.change-password"))
	return web.H{"message": p.I18n.T(lng, "auth.messages.change-password-success")}, nil

}

func (p *Engine) deleteUsersSignOut(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (interface{}, error) {
	user := r.Context().Value(CurrentUser).(*User)
	lng := r.Context().Value(web.LOCALE).(string)
	p.Dao.Log(user.ID, p.W.ClientIP(r), p.I18n.T(lng, "auth.logs.sign-out"))
	return web.H{}, nil
}

package auth

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kapmahc/champak/web"
)

const (
	actConfirm       = "confirm"
	actUnlock        = "unlock"
	actResetPassword = "reset-password"
)

type fmSignUp struct {
	FullName             string `form:"fullName" validate:"required,max=255"`
	Email                string `form:"email" validate:"required,email"`
	Password             string `form:"password" validate:"min=6,max=32"`
	PasswordConfirmation string `form:"passwordConfirmation" validate:"eqfield=Password"`
}

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

func (p *Engine) sendEmail(lng string, user *User, act string) {
	// TODO
}

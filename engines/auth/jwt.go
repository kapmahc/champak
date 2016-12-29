package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	log "github.com/Sirupsen/logrus"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/kapmahc/champak/web"
	"github.com/unrolled/render"
)

const (
	// CurrentUser current user
	CurrentUser = web.KEY("current-user")
)

//Jwt jwt helper
type Jwt struct {
	Key    []byte               `inject:"jwt.key"`
	Method crypto.SigningMethod `inject:"jwt.method"`
	Dao    *Dao                 `inject:""`
	I18n   *web.I18n            `inject:""`
	R      *render.Render       `inject:""`
}

//Validate check jwt
func (p *Jwt) Validate(buf []byte) (jwt.Claims, error) {
	tk, err := jws.ParseJWT(buf)
	if err != nil {
		return nil, err
	}
	if err = tk.Validate(p.Key, p.Method); err != nil {
		return nil, err
	}
	return tk.Claims(), nil
}

//Parse parse from request
func (p *Jwt) Parse(r *http.Request) (jwt.Claims, error) {
	tk, err := jws.ParseJWTFromRequest(r)
	if err != nil {
		return nil, err
	}
	if err = tk.Validate(p.Key, p.Method); err != nil {
		return nil, err
	}
	return tk.Claims(), nil
}

func (p *Jwt) key(kid string) string {
	return fmt.Sprintf("token://%s", kid)
}

//Sum create jwt token
func (p *Jwt) Sum(cm jws.Claims, days int) ([]byte, error) {
	kid := uuid.New().String()
	now := time.Now()
	cm.SetNotBefore(now)
	cm.SetExpiration(now.AddDate(0, 0, days))
	cm.Set("kid", kid)
	//TODO using kid

	jt := jws.NewJWT(cm, p.Method)
	return jt.Serialize(p.Key)
}

func (p *Jwt) getUserFromRequest(req *http.Request) (*User, error) {
	cm, err := p.Parse(req)
	if err != nil {
		return nil, err
	}
	user, err := p.Dao.GetUserByUID(cm.Get("uid").(string))
	if err != nil {
		return nil, err
	}
	if !user.IsConfirm() || user.IsLock() {
		return nil, errors.New("bad user status")
	}
	return user, nil
}

func (p *Jwt) ServeHTTP(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	user, err := p.getUserFromRequest(req)
	if err == nil {
		ctx := context.WithValue(req.Context(), CurrentUser, user)
		req = req.WithContext(ctx)
	} else {
		log.Debug(err)
	}
	next(wrt, req)
}

// MustSignIn must sign-in
func (p *Jwt) MustSignIn(h httprouter.Handle) httprouter.Handle {
	return func(wrt http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		lng := req.Context().Value(web.LOCALE).(string)
		if req.Context().Value(CurrentUser) == nil {
			p.R.Text(wrt, http.StatusForbidden, p.I18n.T(lng, "auth.errors.please-sign-in"))
			return
		}
		h(wrt, req, ps)
	}
}

// MustAdmin must have admin role
func (p *Jwt) MustAdmin(h httprouter.Handle) httprouter.Handle {
	return func(wrt http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		lng := req.Context().Value(web.LOCALE).(string)
		user := req.Context().Value(CurrentUser)
		if user == nil || !p.Dao.Is(user.(*User).ID, RoleAdmin) {
			p.R.Text(wrt, http.StatusForbidden, p.I18n.T(lng, "auth.errors.not-allow"))
			return
		}
		h(wrt, req, ps)
	}
}

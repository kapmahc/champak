package auth

import (
	"fmt"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/SermoDigital/jose/jwt"
	log "github.com/Sirupsen/logrus"
	"github.com/google/uuid"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

const (
	// CurrentUser current user
	CurrentUser = "current_user"
)

//Jwt jwt helper
type Jwt struct {
	Key     []byte               `inject:"jwt.key"`
	Method  crypto.SigningMethod `inject:"jwt.method"`
	Dao     *Dao                 `inject:""`
	I18n    *web.I18n            `inject:""`
	Session *Session             `inject:""`
}

//Validate check jwt
func (p *Jwt) Validate(buf []byte) (jwt.Claims, error) {
	tk, err := jws.ParseJWT(buf)
	if err != nil {
		return nil, err
	}
	err = tk.Validate(p.Key, p.Method)
	return tk.Claims(), err
}

//CurrentUserHandler inject current user
func (p *Jwt) CurrentUserHandler(c *gin.Context) {
	tkn, err := jws.ParseFromRequest(c.Request, jws.Compact)
	if err == nil {
		if err = tkn.Verify(p.Key, p.Method); err == nil {
			data := tkn.Payload().(map[string]interface{})
			if uid, ok := data["uid"]; ok {
				err = p.Session.SetCurrentUser(c, uid.(string))
			}
		}
	}
	log.Debug(err)
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

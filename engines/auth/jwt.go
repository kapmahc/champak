package auth

import (
	"fmt"
	"net/http"
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
	Key    []byte               `inject:"jwt.key"`
	Method crypto.SigningMethod `inject:"jwt.method"`
	Dao    *Dao                 `inject:""`
	I18n   *web.I18n            `inject:""`
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

//MustAdminHandler check must have admin role
func (p *Jwt) MustAdminHandler() gin.HandlerFunc {
	return p.MustRolesHandler("admin")
}

//MustRolesHandler check must have one roles at least
func (p *Jwt) MustRolesHandler(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet("user").(*User)
		for _, a := range p.Dao.Authority(u.ID, "-", 0) {
			for _, r := range roles {
				if a == r {
					return
				}
			}
		}
		c.AbortWithStatus(http.StatusForbidden)
	}
}

//CurrentUserHandler inject current user
func (p *Jwt) CurrentUserHandler(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	tkn, err := jws.ParseFromRequest(c.Request, jws.Compact)

	if err == nil {
		if err = tkn.Verify(p.Key, p.Method); err == nil {
			var user User
			data := tkn.Payload().(map[string]interface{})
			if err = p.Dao.Db.Where("uid = ?", data["uid"]).First(&user).Error; err == nil {
				if !user.IsConfirm() {
					err = p.I18n.E(lng, "auth.errors.user-not-confirm")
				} else if user.IsLock() {
					err = p.I18n.E(lng, "auth.errors.user-is-lock")
				} else {
					c.Set(CurrentUser, &user)
					data := c.MustGet(web.DATA).(gin.H)
					data[CurrentUser] = gin.H{
						"full_name": user.FullName,
						"uid":       user.UID,
					}
					c.Set(web.DATA, user)
					return
				}
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

package auth

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-contrib/sessions"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Session session
type Session struct {
	Dao  *Dao      `inject:""`
	I18n *web.I18n `inject:""`
}

// SetCurrentUser set currrent user
func (p *Session) SetCurrentUser(c *gin.Context, uid string) error {
	lng := c.MustGet(web.LOCALE).(string)
	user, err := p.Dao.GetUserByUID(uid)
	if err != nil {
		return err
	}
	if !user.IsConfirm() {
		return p.I18n.E(lng, "auth.errors.user-not-confirm")
	}
	if user.IsLock() {
		return p.I18n.E(lng, "auth.errors.user-is-lock")
	}
	c.Set(CurrentUser, user)
	data := c.MustGet(web.DATA).(gin.H)
	data[CurrentUser] = gin.H{
		"full_name": user.FullName,
		"uid":       user.UID,
	}
	c.Set(CurrentUser, user)
	return nil
}

//CurrentUserHandler inject current user
func (p *Session) CurrentUserHandler(c *gin.Context) {

	ss := sessions.Default(c)
	if uid := ss.Get("uid"); uid != nil {
		if err := p.SetCurrentUser(c, uid.(string)); err != nil {
			log.Debug(err)
		}
	}
}

//MustSignInHandler check must have admin role
func (p *Session) MustSignInHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Get(CurrentUser); ok {
			data := c.MustGet(web.DATA).(gin.H)
			var links []web.Dropdown
			web.Walk(func(en web.Engine) error {
				items := en.Dashboard(c)
				links = append(links, items...)
				return nil
			})
			data["links"] = links
			c.Set(web.DATA, data)
			return
		}
		p.gotoSignIn(c, "auth.errors.please-sign-in")
	}
}

//MustAdminHandler check must have admin role
func (p *Session) MustAdminHandler() gin.HandlerFunc {
	return p.MustRolesHandler(RoleAdmin)
}

//MustRolesHandler check must have one roles at least
func (p *Session) MustRolesHandler(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(CurrentUser).(*User)
		for _, a := range p.Dao.Authority(u.ID, "-", 0) {
			for _, r := range roles {
				if a == r {
					return
				}
			}
		}

		p.gotoSignIn(c, "auth.errors.not-allow")
	}
}

func (p *Session) gotoSignIn(c *gin.Context, msg string) {
	lng := c.MustGet(web.LOCALE).(string)
	ss := sessions.Default(c)
	ss.AddFlash(p.I18n.T(lng, msg), web.ALERT)
	ss.Save()
	c.Redirect(http.StatusFound, "/users/sign-in")
}

package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/kapmahc/champak/web"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getUsersProfile(c *gin.Context) {

}
func (p *Engine) postUsersProfile(c *gin.Context) {

}

func (p *Engine) deleteUsersSignOut(c *gin.Context) {
	user := c.MustGet(CurrentUser).(*User)
	ss := sessions.Default(c)
	ss.Clear()
	ss.Save()
	lng := c.MustGet(web.LOCALE).(string)
	p.Dao.Log(user.ID, p.I18n.T(lng, "auth.logs.sign-out", c.ClientIP()))
	c.JSON(http.StatusOK, gin.H{"to": "/"})
}

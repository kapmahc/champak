package forum

import (
	"github.com/kapmahc/champak/engines/auth"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) check(c *gin.Context, id uint) error {
	lng := c.MustGet(web.LOCALE).(string)
	user := c.MustGet(auth.CurrentUser).(*auth.User)
	if id != user.ID && !p.Dao.Is(user.ID, auth.RoleAdmin) {
		return p.I18n.E(lng, "auth.errors.not-allow")
	}
	return nil
}

package site

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/kapmahc/champak/engines/auth"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) indexAdminUsers(c *gin.Context) {
	lng := c.MustGet(web.LOCALE).(string)
	data := c.MustGet(web.DATA).(gin.H)
	data["title"] = p.I18n.T(lng, "site.admin.users.title")

	var items []auth.User
	if err := p.Db.Order("last_sign_in_at DESC").Find(&items).Error; err != nil {
		log.Error(err)
	}
	data["items"] = items
	c.HTML(http.StatusOK, "admin/users", data)
}

package auth

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/kapmahc/champak/web"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func (p *Engine) getUsersSelf(c *gin.Context) {
	data := c.MustGet(web.DATA).(gin.H)
	data["links"] = dashboard
	c.HTML(http.StatusOK, "auth/self", data)
}

// -----------------------------------------------------------------------------

var dashboard = make(map[string]interface{})

// Dashboard add dashboard links
func Dashboard(label string, links ...[]web.Link) {
	if _, ok := dashboard[label]; ok {
		log.Warn("already register for label %s", label)
	}
	dashboard[label] = links
}

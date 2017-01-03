package site

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Home home
func (p *Engine) Home() gin.HandlerFunc {
	return p.indexNotices
}

func (p *Engine) getHome(c *gin.Context) {
	fns := make(map[string]gin.HandlerFunc)
	web.Walk(func(en web.Engine) error {
		fns[reflect.TypeOf(en).String()] = en.Home()
		return nil
	})
	fn, ok := fns[fmt.Sprintf("*%s.Engine", viper.GetString("server.root"))]
	if !ok {
		fn = p.newLeaveWord
	}
	fn(c)
}

func (p *Engine) getDashboard(c *gin.Context) {
	data := c.MustGet(web.DATA).(gin.H)
	lng := c.MustGet(web.LOCALE).(string)
	data["title"] = p.I18n.T(lng, "header.dashboard")
	c.HTML(http.StatusOK, "site/dashboard", data)
}

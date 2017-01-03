package site

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/kapmahc/champak/web"
	"github.com/spf13/viper"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Home home
func (p *Engine) Home(c *gin.Context) {
	p.indexNotices(c)
}

func (p *Engine) getHome(c *gin.Context) {
	web.Walk(func(en web.Engine) error {
		if fmt.Sprintf("*%s.Engine", viper.GetString("server.root")) == reflect.TypeOf(en).String() {
			en.Home(c)
			return errors.New("")
		}
		return nil
	})
}

func (p *Engine) getDashboard(c *gin.Context) {
	data := c.MustGet(web.DATA).(gin.H)
	lng := c.MustGet(web.LOCALE).(string)
	data["title"] = p.I18n.T(lng, "header.dashboard")
	c.HTML(http.StatusOK, "site/dashboard", data)
}

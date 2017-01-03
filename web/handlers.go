package web

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin/binding"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// FlashHandler show flash
func FlashHandler(ego string, fn func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fn(c); err != nil {
			ss := sessions.Default(c)
			ss.AddFlash(err.Error(), ALERT)
			ss.Save()
			c.Redirect(http.StatusFound, ego)
		}
	}
}

// PostFormHandler fix gin bind error return 400
func PostFormHandler(ego string, fm interface{}, fn func(*gin.Context, interface{}) error) gin.HandlerFunc {
	return FlashHandler(ego, func(c *gin.Context) error {
		err := binding.FormPost.Bind(c.Request, fm)
		if err == nil {
			err = fn(c, fm)
		} else {
			err = errors.New(strings.Replace(err.Error(), "\n", "<br/>", -1))
		}
		return err
	})
}

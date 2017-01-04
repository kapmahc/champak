package web

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin/binding"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// JSON render json
func JSON(fn func(*gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if val, err := fn(c); err == nil {
			c.JSON(http.StatusOK, val)
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	}
}

// HTML render html
func HTML(fn func(*gin.Context) (string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		tpl, err := fn(c)
		if err != nil {
			ss := sessions.Default(c)
			ss.AddFlash(err.Error(), ALERT)
			ss.Save()
		}
		c.HTML(http.StatusOK, tpl, c.MustGet(DATA))
	}
}

// RedirectTo redirect to
func RedirectTo(fn func(*gin.Context) (string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		url, err := fn(c)
		if err != nil {
			ss := sessions.Default(c)
			ss.AddFlash(err.Error(), ALERT)
			ss.Save()
		}

		if nex, ok := c.Get("next"); ok {
			url = nex.(string)
		}
		if !c.Writer.Written() {
			c.Redirect(http.StatusFound, url)
		}
	}
}

// PostFormHandler fix gin bind error return 400
func PostFormHandler(fm interface{}, fn func(*gin.Context, interface{}) error) gin.HandlerFunc {
	return RedirectTo(func(c *gin.Context) (string, error) {
		err := binding.FormPost.Bind(c.Request, fm)
		if err == nil {
			err = fn(c, fm)
		} else {
			err = errors.New(strings.Replace(err.Error(), "\n", "<br/>", -1))
		}
		return c.PostForm("next"), err
	})
}

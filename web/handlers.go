package web

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin/binding"
	"github.com/gorilla/csrf"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// FlashsHandler flashs handler
func FlashsHandler(c *gin.Context) {
	data := c.MustGet(DATA).(gin.H)
	ss := sessions.Default(c)
	for _, k := range []string{ALERT, NOTICE} {
		data[k] = ss.Flashes(k)
	}
	ss.Save()
	c.Set(DATA, data)
}

// CsrfHandler csrf handler
func CsrfHandler(c *gin.Context) {
	tkn := csrf.Token(c.Request)
	c.Writer.Header().Set("X-CSRF-Token", tkn)
	data := c.MustGet(DATA).(gin.H)
	data["csrf"] = tkn
	c.Set(DATA, data)
}

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
		data := c.MustGet(DATA).(gin.H)
		if err != nil {
			alt := data[ALERT].([]interface{})
			alt = append(alt, err.Error())
			data[ALERT] = alt
		}
		c.HTML(http.StatusOK, tpl, data)
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

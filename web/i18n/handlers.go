package i18n

import (
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/text/language"
	gin "gopkg.in/gin-gonic/gin.v1"
)

const (
	// LOCALE locale key
	LOCALE = "locale"
)

//LocaleHandler detect locale from http header
func LocaleHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 1. Check URL arguments.
		lng := c.Request.URL.Query().Get(LOCALE)

		// 2. Get language information from cookies.
		if len(lng) == 0 {
			if ck, er := c.Request.Cookie(LOCALE); er == nil {
				lng = ck.Value
			}
		}

		// 3. Get language information from 'Accept-Language'.
		if len(lng) == 0 {
			al := c.Request.Header.Get("Accept-Language")
			if len(al) > 4 {
				lng = al[:5]
			}
		}
		tag, err := language.Parse(lng)
		if err != nil {
			log.Error(err)
			tag = language.AmericanEnglish
		}

		// Write cookie
		http.SetCookie(c.Writer, &http.Cookie{
			Name:    LOCALE,
			Value:   tag.String(),
			Expires: time.Now().Add(7 * 24 * time.Hour),
			Path:    "/",
		})

		c.Set(LOCALE, tag.String())

	}
}

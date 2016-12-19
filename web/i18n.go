package web

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	"golang.org/x/text/language"
	gin "gopkg.in/gin-gonic/gin.v1"
)

const (
	// LOCALE locale key
	LOCALE = "locale"
)

//Locale locale model
type Locale struct {
	Model

	Lang    string `gorm:"not null;type:varchar(8);index"`
	Code    string `gorm:"not null;index;type:VARCHAR(255)"`
	Message string `gorm:"not null;type:varchar(800)"`
}

// -----------------------------------------------------------------------------

//I18n i18n
type I18n struct {
	Db        *gorm.DB       `inject:""`
	Cache     *Cache         `inject:""`
	Languages []language.Tag `inject:"languages"`
}

//Tm translate by lang tag
func (p *I18n) Tm(lang string, code string, args ...interface{}) string {
	msg := p.Get(lang, code)
	if len(msg) == 0 {
		return code
	}
	return fmt.Sprintf(msg, args...)
}

//T translate by lang tag
func (p *I18n) T(lang string, code string, args ...interface{}) string {
	msg := p.Get(lang, code)
	if len(msg) == 0 {
		return code
	}
	return fmt.Sprintf(msg, args...)
}

//Set set locale
func (p *I18n) Set(lng string, code, message string) {
	var l Locale
	var err error
	if p.Db.Where("lang = ? AND code = ?", lng, code).First(&l).RecordNotFound() {
		l.Lang = lng
		l.Code = code
		l.Message = message
		err = p.Db.Create(&l).Error
	} else {
		l.Message = message
		err = p.Db.Save(&l).Error
	}
	if err != nil {
		log.Error(err)
	}
}

//Get get locale
func (p *I18n) Get(lng string, code string) string {
	var l Locale
	if err := p.Db.Where("lang = ? AND code = ?", lng, code).First(&l).Error; err != nil {
		log.Error(err)
	}
	return l.Message

}

//Del del locale
func (p *I18n) Del(lng string, code string) {
	if err := p.Db.Where("lang = ? AND code = ?", lng, code).Delete(Locale{}).Error; err != nil {
		log.Error(err)
	}
}

//Codes list locale keys
func (p *I18n) Codes(lang string) ([]string, error) {
	var keys []string
	err := p.Db.Model(&Locale{}).Where("lang = ?", lang).Pluck("code", &keys).Error

	return keys, err
}

//Handler detect locale from http header
func (p *I18n) Handler() gin.HandlerFunc {
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

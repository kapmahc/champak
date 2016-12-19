package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	"golang.org/x/text/language"
)

type localeKeyType string

const (
	// LOCALE locale key
	LOCALE = localeKeyType("locale")
)

//Locale locale model
type Locale struct {
	Model

	Lang    string `gorm:"not null;type:varchar(8);index"`
	Code    string `gorm:"not null;index;type:VARCHAR(255)"`
	Message string `gorm:"not null;type:varchar(800)"`
}

// TableName table name
func (Locale) TableName() string {
	return "locales"
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

// Sync sync records
func (p *I18n) Sync(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		const ext = ".ini"
		name := info.Name()
		if info.Mode().IsRegular() && filepath.Ext(name) == ext {
			log.Debugf("Find locale file %s", path)
			if err != nil {
				return err
			}
			lang := name[0 : len(name)-len(ext)]
			if _, err := language.Parse(lang); err != nil {
				return err
			}
			cfg, err := ini.Load(path)
			if err != nil {
				return err
			}

			items := cfg.Section(ini.DEFAULT_SECTION).KeysHash()
			for k, v := range items {
				var l Locale
				if p.Db.Where("lang = ? AND code = ?", lang, k).First(&l).RecordNotFound() {
					l.Lang = lang
					l.Code = k
					l.Message = v
					if err := p.Db.Create(&l).Error; err != nil {
						return err
					}
				}
			}
			log.Infof("find %d items", len(items))
		}
		return nil
	})
}

//ServeHTTP detect locale from http header
func (p *I18n) ServeHTTP(wrt http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	key := string(LOCALE)
	// 1. Check URL arguments.
	lng := req.URL.Query().Get(key)

	// 2. Get language information from cookies.
	if len(lng) == 0 {
		if ck, er := req.Cookie(key); er == nil {
			lng = ck.Value
		}
	}

	// 3. Get language information from 'Accept-Language'.
	if len(lng) == 0 {
		al := req.Header.Get("Accept-Language")
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
	http.SetCookie(wrt, &http.Cookie{
		Name:    key,
		Value:   tag.String(),
		Expires: time.Now().Add(7 * 24 * time.Hour),
		Path:    "/",
	})

	ctx := context.WithValue(req.Context(), LOCALE, tag.String())
	next(wrt, req.WithContext(ctx))
}

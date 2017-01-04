package web

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	gin "gopkg.in/gin-gonic/gin.v1"

	log "github.com/Sirupsen/logrus"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	"golang.org/x/text/language"
)

//Locale locale model
type Locale struct {
	Model

	Lang    string `gorm:"not null;type:varchar(8);index" json:"lang"`
	Code    string `gorm:"not null;index;type:VARCHAR(255)" json:"code"`
	Message string `gorm:"not null;type:varchar(800)" json:"message"`
}

// TableName table name
func (Locale) TableName() string {
	return "locales"
}

// -----------------------------------------------------------------------------

//I18n i18n
type I18n struct {
	Db    *gorm.DB `inject:""`
	Cache *Cache   `inject:""`
}

// F format message
func (p *I18n) F(lng, code string, obj interface{}) (string, error) {
	msg, err := p.getMessage(lng, code)
	if err != nil {
		return "", err
	}
	tpl, err := template.New("").Parse(msg)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, obj)
	return buf.String(), err
}

//E create an i18n error
func (p *I18n) E(lang string, code string, args ...interface{}) error {
	msg, err := p.getMessage(lang, code)
	if err != nil {
		return errors.New(code)
	}
	return fmt.Errorf(msg, args...)
}

//T translate by lang tag
func (p *I18n) T(lng string, code string, args ...interface{}) string {
	msg, err := p.getMessage(lng, code)
	if err != nil {
		return code
	}
	return fmt.Sprintf(msg, args...)
}

//Set set locale
func (p *I18n) Set(lng string, code, message string) error {
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
	return err
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
func (p *I18n) Del(lng, code string) {
	if err := p.Db.Where("lang = ? AND code = ?", lng, code).Delete(Locale{}).Error; err != nil {
		log.Error(err)
	}
}

func (p *I18n) getMessage(lng, code string) (string, error) {
	key := p.key(lng, code)
	var msg string
	if err := p.Cache.Get(key, &msg); err == nil {
		return msg, nil
	}

	var l Locale
	if err := p.Db.
		Select("message").
		Where("lang = ? AND code = ?", lng, code).
		First(&l).Error; err != nil {
		return "", err
	}
	p.Cache.Set(key, l.Message, time.Hour*24)
	return l.Message, nil
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

func (p *I18n) key(lng, code string) string {
	return fmt.Sprintf("%s/%s", lng, code)
}

//Items list all items
func (p *I18n) Items(lng string) map[string]interface{} {
	rt := make(map[string]interface{})
	var items []Locale
	if err := p.Db.
		Select([]string{"code", "message"}).
		Where("lang = ?", lng).
		Find(&items).Error; err != nil {
		log.Error(err)
	}

	for _, l := range items {
		if strings.HasPrefix(l.Code, "web.") {
			k := l.Code[4:]
			codes := strings.Split(k, ".")
			tmp := rt
			for i, c := range codes {
				if i+1 == len(codes) {
					tmp[c] = l.Message
				} else {
					if tmp[c] == nil {
						tmp[c] = make(map[string]interface{})
					}
					tmp = tmp[c].(map[string]interface{})
				}
			}

		}

	}
	return rt
}

// -----------------------------------------------------------------------------

const (
	// LOCALE locale key
	LOCALE = "locale"
)

// Handler parse locale from request
func (p *I18n) Handler(languages ...string) (gin.HandlerFunc, error) {
	var tags []language.Tag
	for _, l := range languages {
		if lng, err := language.Parse(l); err == nil {
			tags = append(tags, lng)
		} else {
			return nil, err
		}
	}
	matcher := language.NewMatcher(tags)
	return func(c *gin.Context) {
		write := false
		// 1. Check URL arguments.
		lng := c.Request.URL.Query().Get(LOCALE)

		// 2. Get language information from cookies.
		if len(lng) == 0 {
			if ck, er := c.Request.Cookie(LOCALE); er == nil {
				lng = ck.Value
			} else {
				write = true
			}
		}

		// 3. Get language information from 'Accept-Language'.
		if len(lng) == 0 {
			write = true
			al := c.Request.Header.Get("Accept-Language")
			if len(al) > 4 {
				lng = al[:5]
			}
		}

		tag, _, _ := matcher.Match(language.Make(lng))

		// Write cookie
		if write {
			http.SetCookie(c.Writer, &http.Cookie{
				Name:    LOCALE,
				Value:   tag.String(),
				Expires: time.Now().AddDate(10, 0, 0),
				Path:    "/",
			})
		}

		c.Set(LOCALE, tag.String())
		c.Set(DATA, gin.H{
			"locale":    tag.String(),
			"languages": languages,
		})

	}, nil
}

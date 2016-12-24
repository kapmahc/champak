package i18n

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"golang.org/x/text/language"

	log "github.com/Sirupsen/logrus"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/cache"
)

var db *gorm.DB

// DS set data source
func DS(d *gorm.DB) {
	db = d
}

// F format message
func F(lng, code string, obj interface{}) (string, error) {
	msg, err := getMessage(lng, code)
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
func E(lang string, code string, args ...interface{}) error {
	msg, err := getMessage(lang, code)
	if err != nil {
		return errors.New(code)
	}
	return fmt.Errorf(msg, args...)
}

//T translate by lang tag
func T(lng string, code string, args ...interface{}) string {
	msg, err := getMessage(lng, code)
	if err != nil {
		return code
	}
	return fmt.Sprintf(msg, args...)
}

//Set set locale
func Set(lng string, code, message string) error {
	var l Locale
	var err error
	if db.Where("lang = ? AND code = ?", lng, code).First(&l).RecordNotFound() {
		l.Lang = lng
		l.Code = code
		l.Message = message
		err = db.Create(&l).Error
	} else {
		l.Message = message
		err = db.Save(&l).Error
	}
	return err
}

//Del del locale
func Del(lng, code string) {
	if err := db.Where("lang = ? AND code = ?", lng, code).Delete(Locale{}).Error; err != nil {
		log.Error(err)
	}
}

func key(lng, code string) string {
	return fmt.Sprintf("%s/%s", lng, code)
}

func getMessage(lng, code string) (string, error) {
	var v string
	k := key(lng, code)
	if err := cache.Get(k, &v); err == nil {
		return v, nil
	}
	var l Locale
	if err := db.
		Select("message").
		Where("lang = ? AND code = ?", lng, code).
		First(&l).Error; err != nil {
		return "", err
	}

	cache.Set(k, l.Message, 60*60*24)
	return l.Message, nil
}

//Codes list locale keys
func Codes(lang string) ([]string, error) {
	var keys []string
	err := db.Model(&Locale{}).Where("lang = ?", lang).Pluck("code", &keys).Error

	return keys, err
}

// Sync sync records
func Sync(dir string) error {
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
				if db.Where("lang = ? AND code = ?", lang, k).First(&l).RecordNotFound() {
					l.Lang = lang
					l.Code = k
					l.Message = v
					if err := db.Create(&l).Error; err != nil {
						return err
					}
				}
			}
			log.Infof("find %d items", len(items))
		}
		return nil
	})
}

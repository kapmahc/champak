package i18n

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/go-ini/ini"
	"github.com/kapmahc/champak/web/cache"
	"golang.org/x/text/language"
)

var (
	store Store
)

// Use use
func Use(s Store) {
	store = s
}

// F text template message
func F(lang, code string, obj interface{}) (string, error) {
	msg, err := get(lang, code)
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
	msg, err := get(lang, code)
	if err != nil {
		return errors.New(code)
	}
	return fmt.Errorf(msg, args...)
}

//T translate by lang tag
func T(lang string, code string, args ...interface{}) string {
	msg, err := get(lang, code)
	if err != nil {
		return code
	}
	return fmt.Sprintf(msg, args...)
}

//Set set locale
func Set(lang string, code, message string) error {
	return store.Set(lang, code, message, true)
}

//All all items
func All(lang string) (map[string]string, error) {
	return store.All(lang)
}

// Del delete
func Del(lang, code string) error {
	return store.Del(lang, code)
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
				if err := store.Set(lang, k, v, false); err != nil {
					return err
				}
			}
			log.Infof("find %d items", len(items))
		}
		return nil
	})
}

func get(lang, code string) (string, error) {
	k := fmt.Sprintf("locales/%s/%s", lang, code)
	var v string
	e := cache.Get(k, &v)
	if e == nil {
		return v, nil
	}
	v, e = store.Get(lang, code)
	if e == nil {
		if e = cache.Set(k, v, 24*time.Hour); e != nil {
			log.Error(e)
		}
	}
	return v, e
}

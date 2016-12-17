package i18n

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/go-ini/ini"
	"golang.org/x/text/language"
)

//Store i18n store
type Store interface {
	Set(lang string, code, message string)
	Get(lang string, code string) string
	Del(lang string, code string)
	Keys(lang string) ([]string, error)
}

//I18n i18n helper
type I18n struct {
	Store Store `inject:""`

	Locales map[string]map[string]string
}

//Items list all items
func (p *I18n) Items(lng string) map[string]interface{} {
	rt := make(map[string]interface{})
	if items, ok := p.Locales[lng]; ok {
		for k, v := range items {
			if strings.HasPrefix(k, "web.") {
				k = k[4:]
				codes := strings.Split(k, ".")
				tmp := rt
				for i, c := range codes {
					if i+1 == len(codes) {
						tmp[c] = v
					} else {
						if tmp[c] == nil {
							tmp[c] = make(map[string]interface{})
						}
						tmp = tmp[c].(map[string]interface{})
					}
				}

			}
		}
	}
	return rt
}

//Exist is lang exist?
func (p *I18n) Exist(lang string) bool {
	_, ok := p.Locales[lang]
	return ok
}

//Load load locales from filesystem
func (p *I18n) Load(dir string) error {
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

			p.Locales[lang] = cfg.Section(ini.DEFAULT_SECTION).KeysHash()
			log.Infof("find %d items", len(p.Locales[lang]))
		}
		return nil
	})
}

func (p *I18n) set(lng *language.Tag, code, message string) {
	lang := lng.String()
	if _, ok := p.Locales[lang]; !ok {
		p.Locales[lang] = make(map[string]string)
	}
	p.Locales[lang][code] = message
}

//T translate by lang tag
func (p *I18n) T(lang string, code string, args ...interface{}) string {
	msg := p.Store.Get(lang, code)
	if len(msg) == 0 {
		if items, ok := p.Locales[lang]; ok {
			msg = items[code]
		}
	}
	if len(msg) == 0 {
		msg = code
	}
	return fmt.Sprintf(msg, args...)
}

package web_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kapmahc/champak/web"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/language"
)

var lang = language.SimplifiedChinese.String()

func OpenDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	return db, nil
}

func TestI18n(t *testing.T) {
	db, err := OpenDatabase()
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&web.Locale{})
	db.Model(&web.Locale{}).AddUniqueIndex("idx_locales_lang_code", "lang", "code")
	p := &web.I18n{Db: db, Cache: &web.Cache{Redis: OpenRedis(), Namespace: "test"}}
	key := "hello"
	val := "你好"
	p.Set(lang, key, val)
	p.Set(lang, key+".1", val)
	if val1 := p.Get(lang, key); val != val1 {
		t.Errorf("want %s, get %s", val, val1)
	}
	ks, err := p.Codes(lang)
	if err != nil {
		t.Fatal(err)
	}
	if len(ks) == 0 {
		t.Errorf("empty keys")
	} else {
		t.Log(ks)
	}
	p.Del(lang, key)
}

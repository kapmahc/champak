package i18n_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kapmahc/champak/cache"
	"github.com/kapmahc/champak/i18n"
	"golang.org/x/text/language"
)

var lang = language.SimplifiedChinese.String()

func openDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	db.AutoMigrate(&i18n.Locale{})
	db.Model(&i18n.Locale{}).AddUniqueIndex("idx_locales_lang_code", "lang", "code")
	return db, nil
}

func TestI18n(t *testing.T) {
	db, err := openDatabase()
	if err != nil {
		t.Fatal(err)
	}
	i18n.DS(db)
	cache.New("localhost", 6379, 0, "test")

	key := "hello"
	val := "你好"
	i18n.Set(lang, key, val)
	i18n.Set(lang, key+".1", val)
	if val1 := i18n.T(lang, key); val != val1 {
		t.Errorf("want %s, get %s", val, val1)
	}
	ks, err := i18n.Codes(lang)
	if err != nil {
		t.Fatal(err)
	}
	if len(ks) == 0 {
		t.Errorf("empty keys")
	} else {
		t.Log(ks)
	}
	i18n.Del(lang, key)
}

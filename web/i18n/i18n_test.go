package i18n_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kapmahc/champak/web/cache"
	"github.com/kapmahc/champak/web/cache/redis"
	"github.com/kapmahc/champak/web/i18n"
	_gorm "github.com/kapmahc/champak/web/i18n/gorm"
	"golang.org/x/text/language"
)

var lang = language.SimplifiedChinese.String()

func TestGorm(t *testing.T) {
	db, err := openDb()
	if err != nil {
		t.Fatal(err)
	}

	testI18n(t, _gorm.New(db, true))

}

func openDb() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return db, nil
}

func testI18n(t *testing.T, s i18n.Store) {
	cache.Use(redis.New("localhost", 6379, 0, "test"))
	i18n.Use(s)

	key := "hello"
	val := "你好"
	i18n.Set(lang, key, val)
	i18n.Set(lang, key+".1", val)
	if val1 := i18n.T(lang, key); val != val1 {
		t.Errorf("want %s, get %s", val, val1)
	}
	ks, err := i18n.All(lang)
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

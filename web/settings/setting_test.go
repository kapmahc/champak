package settings_test

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kapmahc/champak/web/crypto"
	"github.com/kapmahc/champak/web/settings"
)

func openDb() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return db, nil
}

func TestSettings(t *testing.T) {
	if err := crypto.Use([]byte("1234567890123456"), []byte("123456")); err != nil {
		t.Fatal(err)
	}
	db, err := openDb()
	if err != nil {
		t.Fatal(err)
	}
	settings.Use(db, true)

	key := "hello"
	val := "你好"
	var tmp string
	if err := settings.Set(key, val, true); err != nil {
		t.Fatal(err)
	}
	if err := settings.Get(key, &tmp); err != nil {
		t.Fatal(err)
	}
	if tmp != val {
		t.Fatalf("want %s, get %s", val, tmp)
	}

	if err := settings.Set(key, val, false); err != nil {
		t.Fatal(err)
	}
	if err := settings.Get(key, &tmp); err != nil {
		t.Fatal(err)
	}
	if tmp != val {
		t.Fatalf("want %s, get %s", val, tmp)
	}

}

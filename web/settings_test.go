package web_test

import (
	"testing"

	"github.com/kapmahc/champak/web"
)

func TestSettings(t *testing.T) {
	db, err := OpenDatabase()
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&web.Setting{})
	en, err := NewSecurity()
	if err != nil {
		t.Fatal(err)
	}
	p := &web.Settings{Db: db, Security: en}
	key := "hello"
	val := "你好"
	var val1 string
	if err := p.Set(key, val, true); err != nil {
		t.Fatal(err)
	}
	if err := p.Get(key, &val1); err != nil {
		t.Fatal(err)
	}
	if val != val1 {
		t.Errorf("want %s, get %s", val, val1)
	}

	if err := p.Set(key, val, false); err != nil {
		t.Fatal(err)
	}
	if err := p.Get(key, &val1); err != nil {
		t.Fatal(err)
	}
	if val != val1 {
		t.Errorf("want %s, get %s", val, val1)
	}
}

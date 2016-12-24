package cache_test

import (
	"testing"

	"github.com/kapmahc/champak/cache"
)

type S struct {
	IV int
	SV string
}

func TestCache(t *testing.T) {
	cache.New("localhost", 6379, 0, "test")

	s1 := S{IV: 100, SV: "hello, champak!"}
	if err := cache.Set("hello", &s1, 60); err != nil {
		t.Fatal(err)
	}
	var s2 S
	if err := cache.Get("hello", &s2); err == nil {
		if s1.IV != s2.IV || s1.SV != s2.SV {
			t.Fatalf("wang %v get %v", s1, s2)
		}
	} else {
		t.Fatal(err)
	}

	if keys, err := cache.Keys(); err == nil {
		t.Logf("keys: %v", keys)
	} else {
		t.Fatal(err)
	}

	if err := cache.Flush(); err != nil {
		t.Fatal(err)
	}
}

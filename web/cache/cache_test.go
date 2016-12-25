package cache_test

import (
	"testing"
	"time"

	"github.com/kapmahc/champak/web/cache"
	"github.com/kapmahc/champak/web/cache/redis"
)

type S struct {
	IV int
	SV string
}

func TestRedis(t *testing.T) {
	testCache(t, redis.New("localhost", 6379, 0, "test"))
}

func testCache(t *testing.T, s cache.Store) {
	cache.Use(s)
	s1 := S{IV: 100, SV: "hello, champak!"}
	if err := cache.Set("hello", &s1, 24*time.Hour); err != nil {
		t.Fatal(err)
	}
	var s2 S
	if err := cache.Get("hello", &s2); err == nil {
		t.Logf("s1 = %+v", s1)
		t.Logf("s2 = %+v", s2)
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

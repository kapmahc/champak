package ops_test

import (
	"testing"

	"github.com/garyburd/redigo/redis"
)

func TestRedisInfo(t *testing.T) {
	con, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		t.Fatal(err)
	}
	rst, err := redis.String(con.Do("INFO"))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", rst)
}

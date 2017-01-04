package auth_test

import (
	"reflect"
	"testing"

	"github.com/kapmahc/champak/engines/auth"
)

func TestTypeName(t *testing.T) {
	var u auth.User
	t.Log(reflect.TypeOf(u).String())
	t.Log(reflect.TypeOf(&u).String())
	t.Log(reflect.Indirect(reflect.ValueOf(u)).Field(0).Type().Name())
}

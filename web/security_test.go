package web_test

import (
	"crypto/aes"
	"testing"

	"github.com/kapmahc/champak/web"
)

func NewSecurity() (*web.Security, error) {
	cip, err := aes.NewCipher([]byte("1234567890123456"))
	if err != nil {
		return nil, err
	}
	return &web.Security{Cip: cip, Key: []byte("123456")}, nil
}
func TestSecurity(t *testing.T) {
	en, err := NewSecurity()
	if err != nil {
		t.Fatal(err)
	}
	hello := "Hello, Champak!"
	code, err := en.Encrypt([]byte(hello))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(code))
	plain, err := en.Decrypt(code)
	if err != nil {
		t.Fatal(err)
	}
	if string(plain) != hello {
		t.Fatalf("wang %s get %s", hello, string(plain))
	}

	code = en.Sum([]byte(hello))
	t.Log(string(code))
	if !en.Chk([]byte(hello), code) {
		t.Fatalf("check password failed")
	}

}

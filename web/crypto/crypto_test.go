package crypto_test

import (
	"testing"

	"github.com/kapmahc/champak/web/crypto"
)

const hello = "Hello, Champak!"

func TestAes(t *testing.T) {
	if err := crypto.Use([]byte("1234567890123456"), []byte("123456")); err != nil {
		t.Fatal(err)
	}

	code, err := crypto.Encrypt([]byte(hello))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(code))
	plain, err := crypto.Decrypt(code)
	if err != nil {
		t.Fatal(err)
	}
	if string(plain) != hello {
		t.Fatalf("wang %s get %s", hello, string(plain))
	}

}

func TestHmac(t *testing.T) {
	code := crypto.Sum([]byte(hello))
	t.Log(string(code))
	if !crypto.Chk([]byte(hello), code) {
		t.Fatalf("check password failed")
	}
}

func TestRand(t *testing.T) {
	s1 := crypto.Rand(8)
	s2 := crypto.Rand(8)
	if s1 == s2 {
		t.Fail()
	}
}

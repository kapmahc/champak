package crypto

import (
	"crypto/hmac"
	"crypto/sha512"
)

var hmacKey []byte

// Sum sum hmac
func Sum(plain []byte) []byte {
	mac := hmac.New(sha512.New, hmacKey)
	return mac.Sum(plain)
}

// Chk chk hmac
func Chk(plain, code []byte) bool {
	return hmac.Equal(Sum(plain), code)
}

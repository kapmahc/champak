package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

var aesCip cipher.Block

// AesKey set aes-key
func AesKey(k []byte) error {
	var err error
	aesCip, err = aes.NewCipher(k)
	return err
}

// Encrypt aes encrypt
func Encrypt(buf []byte) ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(aesCip, iv)
	val := make([]byte, len(buf))
	cfb.XORKeyStream(val, buf)

	return append(val, iv...), nil
}

// Decrypt aes decrypt
func Decrypt(buf []byte) ([]byte, error) {
	bln := len(buf)
	cln := bln - aes.BlockSize
	ct := buf[0:cln]
	iv := buf[cln:bln]

	cfb := cipher.NewCFBDecrypter(aesCip, iv)
	val := make([]byte, cln)
	cfb.XORKeyStream(val, ct)
	return val, nil
}

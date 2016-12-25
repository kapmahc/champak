package crypto

import "crypto/aes"

// Use set keys
func Use(ak, hk []byte) error {
	hmacKey = hk
	var err error
	aesCip, err = aes.NewCipher(ak)
	return err
}

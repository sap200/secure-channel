package generator

import (
	"crypto/rand"
	"crypto/rsa"
)

func NewGenerator() rsa.PrivateKey {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	return *privKey
}

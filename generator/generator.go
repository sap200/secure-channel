// Package generator
// generates new private and public key pair for client or server
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

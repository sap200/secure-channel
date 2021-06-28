package store

import (
	"crypto/rsa"
)

var Store map[string]rsa.PublicKey

func InitStore() map[string]rsa.PublicKey {
	return map[string]rsa.PublicKey{}
}

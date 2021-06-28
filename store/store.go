// package store stores the public key of all the connections made to securely communicate with them
package store

import (
	"crypto/rsa"
)

// Store variable is a key value mapping for actually storing the public key corresponding to
// the string representation of connection's remote address
var Store map[string]rsa.PublicKey

func InitStore() map[string]rsa.PublicKey {
	return map[string]rsa.PublicKey{}
}

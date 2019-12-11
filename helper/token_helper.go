package helper

import (
	"crypto/rand"
	"fmt"
)

// GenerateRecoveryToken used to generate recovery token
func GenerateRecoveryToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

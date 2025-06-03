package core

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateNonce generates a random nonce
func GenerateNonce() (string, error) {
	nonce := make([]byte, 16)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(nonce), nil
}

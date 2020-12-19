package security

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateBase64RandomToken generates a base64 random token
func GenerateBase64RandomToken() (string, error) {

	key := make([]byte, 64)

	if _, error := rand.Read(key); error != nil {
		return "", error
	}

	base64Token := base64.StdEncoding.EncodeToString(key)

	return base64Token, nil
}

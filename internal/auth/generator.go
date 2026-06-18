package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateToken() (string, error) {
	var data [32]byte
	if _, err := rand.Read(data[:]); err != nil {
		return "", fmt.Errorf("generate local token: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(data[:]), nil
}

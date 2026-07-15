package shortsfactory

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func hashJSON(value any) (string, error) {
	body, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(body)
	return hex.EncodeToString(sum[:]), nil
}

func mustHash(value any) string {
	hash, err := hashJSON(value)
	if err != nil {
		panic(err)
	}
	return hash
}

package auth

import (
	"errors"
	"os"
	"strings"
)

func Read(root string) (string, error) {
	data, err := os.ReadFile(LocalTokenPath(root))
	if err != nil {
		return "", err
	}
	token := strings.TrimSpace(string(data))
	if token == "" {
		return "", errors.New("local token is empty")
	}
	return token, nil
}

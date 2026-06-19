package linear

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type tokenConfig struct {
	Value  string
	Source string
}

func loadToken(root string) (tokenConfig, error) {
	if value := strings.TrimSpace(os.Getenv("LINEAR_API_KEY")); value != "" {
		return tokenConfig{Value: value, Source: "env:LINEAR_API_KEY"}, nil
	}
	tokenPath := filepath.Join(root, "data", "private", "linear-token.txt")
	data, err := os.ReadFile(tokenPath)
	if err != nil {
		return tokenConfig{}, err
	}
	value := strings.TrimSpace(string(data))
	if value == "" {
		return tokenConfig{}, errors.New("Linear token file is empty")
	}
	return tokenConfig{Value: value, Source: "file:data/private/linear-token.txt"}, nil
}

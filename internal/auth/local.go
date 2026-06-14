package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const localTokenRelativePath = "data/private/local-token.txt"

type LocalTokenStatus struct {
	Configured bool   `json:"configured"`
	Path       string `json:"path"`
	Mode       string `json:"mode,omitempty"`
	Message    string `json:"message"`
}

type LocalTokenResult struct {
	Configured bool   `json:"configured"`
	Path       string `json:"path"`
	Token      string `json:"token,omitempty"`
	Rotated    bool   `json:"rotated"`
	Message    string `json:"message"`
}

func LocalTokenPath(root string) string {
	return filepath.Join(root, filepath.FromSlash(localTokenRelativePath))
}

func Status(root string) LocalTokenStatus {
	path := LocalTokenPath(root)
	info, err := os.Stat(path)
	if err != nil {
		return LocalTokenStatus{
			Configured: false,
			Path:       localTokenRelativePath,
			Message:    "local LAN token is not configured",
		}
	}
	return LocalTokenStatus{
		Configured: true,
		Path:       localTokenRelativePath,
		Mode:       info.Mode().Perm().String(),
		Message:    "local LAN token is configured",
	}
}

func Create(root string, rotate bool) (LocalTokenResult, error) {
	path := LocalTokenPath(root)
	if !rotate {
		if _, err := os.Stat(path); err == nil {
			return LocalTokenResult{
				Configured: true,
				Path:       localTokenRelativePath,
				Message:    "local LAN token already exists; use rotate to replace it",
			}, nil
		} else if !errors.Is(err, os.ErrNotExist) {
			return LocalTokenResult{}, err
		}
	}
	token, err := generateToken()
	if err != nil {
		return LocalTokenResult{}, err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return LocalTokenResult{}, err
	}
	if err := os.WriteFile(path, []byte(token+"\n"), 0o600); err != nil {
		return LocalTokenResult{}, err
	}
	return LocalTokenResult{
		Configured: true,
		Path:       localTokenRelativePath,
		Token:      token,
		Rotated:    rotate,
		Message:    "local LAN token written; store it only on trusted household clients",
	}, nil
}

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

func generateToken() (string, error) {
	var data [32]byte
	if _, err := rand.Read(data[:]); err != nil {
		return "", fmt.Errorf("generate local token: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(data[:]), nil
}

package auth

import (
	"errors"
	"os"
	"path/filepath"
)

func Create(root string, rotate bool) (LocalTokenResult, error) {
	path := LocalTokenPath(root)
	if !rotate {
		if existing, err := existingTokenResult(path); err != nil || existing.Configured {
			return existing, err
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

func existingTokenResult(path string) (LocalTokenResult, error) {
	if _, err := os.Stat(path); err == nil {
		return LocalTokenResult{
			Configured: true,
			Path:       localTokenRelativePath,
			Message:    "local LAN token already exists; use rotate to replace it",
		}, nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return LocalTokenResult{}, err
	}
	return LocalTokenResult{}, nil
}

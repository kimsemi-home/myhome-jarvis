package auth

import "path/filepath"

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

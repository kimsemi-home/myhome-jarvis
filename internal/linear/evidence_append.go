package linear

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func AppendWriteEvidence(root string, action string, issueKey string) error {
	action = strings.TrimSpace(action)
	if action == "" {
		return errors.New("linear write evidence action is required")
	}
	path := filepath.Join(root, filepath.FromSlash(WriteEvidenceRelativePath))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(WriteEvidence{
		At:       time.Now().UTC().Format(time.RFC3339),
		Action:   action,
		IssueKey: publicIssueKey(issueKey),
		Synced:   true,
	})
}

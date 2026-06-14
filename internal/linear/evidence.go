package linear

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const WriteEvidenceRelativePath = "data/private/linear-write-evidence.jsonl"

type WriteEvidence struct {
	At       string `json:"at"`
	Action   string `json:"action"`
	IssueKey string `json:"issue_key,omitempty"`
	Synced   bool   `json:"synced"`
}

type WriteEvidenceStatus struct {
	EvidencePath         string         `json:"evidence_path"`
	SyncedMutationCount  int            `json:"synced_mutation_count"`
	HasSyncedMutation    bool           `json:"has_synced_mutation"`
	LatestSyncedMutation *WriteEvidence `json:"latest_synced_mutation,omitempty"`
}

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

func WriteEvidenceStatusForRoot(root string) (WriteEvidenceStatus, error) {
	return WriteEvidenceStatusForPath(root, WriteEvidenceRelativePath)
}

func WriteEvidenceStatusForPath(root string, relativePath string) (WriteEvidenceStatus, error) {
	relativePath = strings.TrimSpace(relativePath)
	if relativePath == "" {
		relativePath = WriteEvidenceRelativePath
	}
	status := WriteEvidenceStatus{
		EvidencePath: privateRelativePath(filepathJoinSlash(root, relativePath)),
	}
	path := filepath.Join(root, filepath.FromSlash(relativePath))
	file, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return status, nil
	}
	if err != nil {
		return status, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	for {
		var event WriteEvidence
		err := decoder.Decode(&event)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return status, err
		}
		event.Action = strings.TrimSpace(event.Action)
		event.IssueKey = publicIssueKey(event.IssueKey)
		if !event.Synced || event.Action == "" {
			continue
		}
		status.SyncedMutationCount++
		status.HasSyncedMutation = true
		latest := event
		status.LatestSyncedMutation = &latest
	}
	return status, nil
}

func publicIssueKey(value string) string {
	value = strings.TrimSpace(value)
	parts := strings.Split(value, "-")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return ""
	}
	for _, char := range parts[0] {
		if !((char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')) {
			return ""
		}
	}
	for _, char := range parts[1] {
		if char < '0' || char > '9' {
			return ""
		}
	}
	return strings.ToUpper(parts[0]) + "-" + parts[1]
}

package linear

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

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
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(relativePath)))
	if errors.Is(err, os.ErrNotExist) {
		return status, nil
	}
	if err != nil {
		return status, err
	}
	defer file.Close()

	return readWriteEvidenceStatus(file, status)
}

func readWriteEvidenceStatus(reader io.Reader, status WriteEvidenceStatus) (WriteEvidenceStatus, error) {
	decoder := json.NewDecoder(reader)
	for {
		var event WriteEvidence
		err := decoder.Decode(&event)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return status, err
		}
		addSyncedWriteEvidence(&status, event)
	}
	return status, nil
}

func addSyncedWriteEvidence(status *WriteEvidenceStatus, event WriteEvidence) {
	event.Action = strings.TrimSpace(event.Action)
	event.IssueKey = publicIssueKey(event.IssueKey)
	if !event.Synced || event.Action == "" {
		return
	}
	status.SyncedMutationCount++
	status.HasSyncedMutation = true
	latest := event
	status.LatestSyncedMutation = &latest
}

package linear

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func readReplayedEntryIDs(root string) (map[string]bool, error) {
	path := filepath.Join(root, filepath.FromSlash(OfflineReplayRelativePath))
	file, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return map[string]bool{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	replayed := map[string]bool{}
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		var record replayRecord
		if err := json.Unmarshal(scanner.Bytes(), &record); err != nil {
			return nil, err
		}
		if record.Synced && strings.TrimSpace(record.EntryID) != "" {
			replayed[record.EntryID] = true
		}
	}
	return replayed, scanner.Err()
}

func appendReplayRecord(root string, entry queuedOfflineAction) error {
	path := filepath.Join(root, filepath.FromSlash(OfflineReplayRelativePath))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(replayRecord{
		At:       time.Now().UTC().Format(time.RFC3339),
		EntryID:  entry.EntryID,
		Kind:     entry.Kind,
		IssueKey: replayIssueKey(entry),
		Synced:   true,
	})
}

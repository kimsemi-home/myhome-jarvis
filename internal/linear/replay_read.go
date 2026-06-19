package linear

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func readQueuedOfflineActions(root string) ([]queuedOfflineAction, error) {
	path := filepath.Join(root, filepath.FromSlash(OfflineQueueRelativePath))
	file, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	entries := []queuedOfflineAction{}
	for scanner.Scan() {
		entry, ok, err := decodeQueuedOfflineAction(scanner.Text())
		if err != nil {
			return nil, err
		}
		if ok {
			entries = append(entries, entry)
		}
	}
	return entries, scanner.Err()
}

func decodeQueuedOfflineAction(line string) (queuedOfflineAction, bool, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return queuedOfflineAction{}, false, nil
	}
	var entry queuedOfflineAction
	if err := json.Unmarshal([]byte(line), &entry); err != nil {
		return queuedOfflineAction{}, false, err
	}
	entry.Kind = strings.TrimSpace(entry.Kind)
	entry.EntryID = offlineEntryID(line)
	return entry, true, nil
}

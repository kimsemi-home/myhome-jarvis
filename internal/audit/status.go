package audit

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

func CommandIntentStatusForRoot(root string) (CommandIntentStatus, error) {
	status := CommandIntentStatus{
		Path:      commandIntentRelativePath,
		CheckedAt: time.Now().UTC().Format(time.RFC3339),
	}
	file, err := os.Open(commandIntentPath(root))
	if errors.Is(err, os.ErrNotExist) {
		return status, nil
	}
	if err != nil {
		return status, err
	}
	defer file.Close()
	return readCommandIntentStatus(file, status)
}

func readCommandIntentStatus(file *os.File, status CommandIntentStatus) (CommandIntentStatus, error) {
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if err := addCommandIntentLine(&status, line); err != nil {
			return status, err
		}
	}
	return status, scanner.Err()
}

func addCommandIntentLine(status *CommandIntentStatus, line string) error {
	var event CommandIntentEvent
	if err := json.Unmarshal([]byte(line), &event); err != nil {
		return err
	}
	status.Exists = true
	status.Count++
	status.Last = &event
	return nil
}

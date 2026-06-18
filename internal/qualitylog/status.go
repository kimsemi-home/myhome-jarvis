package qualitylog

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

func StatusForRoot(root string) (Status, error) {
	status := Status{
		Path:      RelativePath,
		CheckedAt: time.Now().UTC().Format(time.RFC3339),
	}
	file, err := os.Open(runPath(root))
	if errors.Is(err, os.ErrNotExist) {
		return status, nil
	}
	if err != nil {
		return status, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var run Run
		if err := json.Unmarshal([]byte(line), &run); err != nil {
			return status, err
		}
		status.Exists = true
		status.Count++
		status.Last = &run
	}
	if err := scanner.Err(); err != nil {
		return status, err
	}
	return status, nil
}

package evidence

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func countJSONL(root string, rel string) (int, bool, string, error) {
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(rel)))
	if errors.Is(err, os.ErrNotExist) {
		return 0, false, "", nil
	}
	if err != nil {
		return 0, false, "", err
	}
	defer file.Close()

	count := 0
	last := ""
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		count++
		last = laterRFC3339(last, jsonlAt(line))
	}
	if err := scanner.Err(); err != nil {
		return 0, false, "", err
	}
	return count, true, last, nil
}

func jsonlAt(line string) string {
	var row struct {
		At string `json:"at"`
	}
	if err := json.Unmarshal([]byte(line), &row); err != nil {
		return ""
	}
	return row.At
}

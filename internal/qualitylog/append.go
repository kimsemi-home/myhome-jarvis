package qualitylog

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func AppendRun(root string, run Run) error {
	if strings.TrimSpace(root) == "" {
		return errors.New("root is required")
	}
	if strings.TrimSpace(run.At) == "" {
		run.At = time.Now().UTC().Format(time.RFC3339)
	}
	path := runPath(root)
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := json.Marshal(run)
	if err != nil {
		return err
	}
	data = append(data, '\n')
	_, err = file.Write(data)
	return err
}

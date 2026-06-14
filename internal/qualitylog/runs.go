package qualitylog

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const RelativePath = "data/private/quality/runs.jsonl"

type Step struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Run struct {
	At             string `json:"at"`
	OK             bool   `json:"ok"`
	DurationMillis int64  `json:"duration_millis"`
	StepCount      int    `json:"step_count"`
	PassCount      int    `json:"pass_count"`
	FailCount      int    `json:"fail_count"`
	SkipCount      int    `json:"skip_count"`
	Steps          []Step `json:"steps"`
}

type Status struct {
	Path      string `json:"path"`
	Exists    bool   `json:"exists"`
	Count     int    `json:"count"`
	Last      *Run   `json:"last,omitempty"`
	CheckedAt string `json:"checked_at"`
}

func NewRun(started time.Time, ok bool, steps []Step) Run {
	run := Run{
		At:             started.UTC().Format(time.RFC3339),
		OK:             ok,
		DurationMillis: time.Since(started).Milliseconds(),
		StepCount:      len(steps),
		Steps:          make([]Step, 0, len(steps)),
	}
	for _, step := range steps {
		step.Name = strings.TrimSpace(step.Name)
		step.Status = normalizeStatus(step.Status)
		if step.Name == "" {
			step.Name = "unknown"
		}
		switch step.Status {
		case "pass":
			run.PassCount++
		case "fail":
			run.FailCount++
		case "skip":
			run.SkipCount++
		}
		run.Steps = append(run.Steps, step)
	}
	return run
}

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

func runPath(root string) string {
	return filepath.Join(root, filepath.FromSlash(RelativePath))
}

func normalizeStatus(status string) string {
	status = strings.TrimSpace(strings.ToLower(status))
	switch status {
	case "pass", "fail", "skip":
		return status
	default:
		return "unknown"
	}
}

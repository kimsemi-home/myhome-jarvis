package pdca

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func scanCycles(root string, policy Policy, status *Status) error {
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(policy.PrivateCycleLedger)))
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer file.Close()
	status.Exists = true
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		if scanCycleLine(policy, status, scanner.Text()) {
			status.CycleCount++
		}
	}
	return scanner.Err()
}

func scanCycleLine(policy Policy, status *Status, text string) bool {
	line := strings.TrimSpace(text)
	if line == "" {
		return false
	}
	var cycle Cycle
	if err := json.Unmarshal([]byte(line), &cycle); err != nil || invalidCycle(policy, cycle) {
		status.InvalidCycleCount++
		return false
	}
	status.ByStatus[cycle.Status]++
	if cycle.Status == "closed" {
		status.ClosedCount++
	} else {
		status.OpenCount++
	}
	status.LastObservedAt = later(status.LastObservedAt, cycle.At)
	return true
}

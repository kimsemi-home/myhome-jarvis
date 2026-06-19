package learning

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := newStatus(policy)
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(policy.PrivateLedger)))
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
		observation, ok, err := decodeObservationLine(scanner.Text())
		if err != nil {
			return status, err
		}
		if ok {
			applyObservation(&status, observation)
		}
	}
	if err := scanner.Err(); err != nil {
		return status, err
	}
	return status, nil
}

func decodeObservationLine(line string) (Observation, bool, error) {
	line = strings.TrimSpace(line)
	if line == "" {
		return Observation{}, false, nil
	}
	var observation Observation
	if err := json.Unmarshal([]byte(line), &observation); err != nil {
		return Observation{}, false, err
	}
	return observation, true, nil
}

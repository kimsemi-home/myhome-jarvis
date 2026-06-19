package evidence

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func inspectLearningSource(
	root string,
	policy Policy,
	source PrivateSource,
	status *Status,
	artifactRefs map[string]bool,
) (SourceStatus, error) {
	sourceStatus := newSourceStatus(source)
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(source.Path)))
	if errors.Is(err, os.ErrNotExist) {
		return sourceStatus, nil
	}
	if err != nil {
		return sourceStatus, err
	}
	defer file.Close()

	sourceStatus.Present = true
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		if err := inspectLearningLine(
			root,
			policy,
			strings.TrimSpace(scanner.Text()),
			status,
			artifactRefs,
			&sourceStatus,
		); err != nil {
			return sourceStatus, err
		}
	}
	return sourceStatus, scanner.Err()
}

func inspectLearningLine(
	root string,
	policy Policy,
	line string,
	status *Status,
	artifactRefs map[string]bool,
	sourceStatus *SourceStatus,
) error {
	if line == "" {
		return nil
	}
	var observation learningObservation
	if err := json.Unmarshal([]byte(line), &observation); err != nil {
		return err
	}
	addLearningObservation(status, sourceStatus, observation)
	return addLearningEvidenceRefs(root, policy, observation, status, artifactRefs)
}

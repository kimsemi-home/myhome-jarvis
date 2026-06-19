package evidence

import (
	"errors"
	"os"
	"path/filepath"
)

func addLearningObservation(
	status *Status,
	sourceStatus *SourceStatus,
	observation learningObservation,
) {
	sourceStatus.Count++
	status.NodeCount++
	status.ByNodeKind[sourceStatus.NodeKind]++
	if observation.Status != "closed" {
		status.OpenLearningCount++
	}
	updateLastObserved(status, observation.At)
}

func addLearningEvidenceRefs(
	root string,
	policy Policy,
	observation learningObservation,
	status *Status,
	artifactRefs map[string]bool,
) error {
	for _, ref := range observation.EvidenceRefs {
		normalized, err := normalizeEvidenceRef(policy, ref)
		if err != nil {
			return err
		}
		artifactRefs[normalized] = true
		status.EdgeCount++
		status.ByEdgeKind["supports"]++
		if err := inspectEvidenceRef(root, normalized, status); err != nil {
			return err
		}
	}
	return nil
}

func inspectEvidenceRef(root string, normalized string, status *Status) error {
	_, err := os.Stat(filepath.Join(root, filepath.FromSlash(normalized)))
	if errors.Is(err, os.ErrNotExist) {
		status.DanglingEvidenceRefCount++
		return nil
	}
	return err
}

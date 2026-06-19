package pdca

import (
	"errors"
	"os"
	"path/filepath"
	"time"
)

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := newStatus(policy)
	for _, step := range policy.Steps {
		if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(step.Artifact))); err == nil {
			status.ReadyStepCount++
		} else if errors.Is(err, os.ErrNotExist) {
			status.MissingArtifactCount++
		} else {
			return Status{}, err
		}
	}
	if err := scanCycles(root, policy, &status); err != nil {
		return Status{}, err
	}
	status.Ready = status.MissingArtifactCount == 0 && status.InvalidCycleCount == 0
	return status, nil
}

func newStatus(policy Policy) Status {
	return Status{
		PolicyPath:          PolicyRelativePath,
		LedgerPath:          policy.PrivateCycleLedger,
		StepCount:           len(policy.Steps),
		EvidenceSourceCount: len(policy.EvidenceSources),
		ByStatus:            map[string]int{},
		CheckedAt:           time.Now().UTC().Format(time.RFC3339),
	}
}

package evidencequality

import (
	"errors"
	"time"
)

func classifySnapshotError(status *Status, err error) {
	if errors.Is(err, errMissingEvidenceRef) {
		status.MissingEvidenceCount++
	} else {
		status.InvalidSnapshotCount++
	}
}

func applySnapshot(status *Status, policy Policy, snapshot Snapshot, checkedAt time.Time) {
	status.SnapshotCount++
	status.ByQualityLevel[snapshot.QualityLevel]++
	status.ByMappingConfidence[snapshot.MappingConfidence]++
	status.ByPurpose[snapshot.Purpose]++
	for _, reason := range snapshot.ReassessmentReasons {
		status.ByReassessmentReason[reason]++
	}
	if isStale(policy, snapshot, checkedAt) {
		status.StaleSnapshotCount++
	}
	if snapshot.QualityLevel == "low" {
		status.LowQualityCount++
	}
	if snapshot.QualityLevel == "blocked" {
		status.BlockedQualityCount++
	}
	if snapshot.MappingConfidence == "low" || snapshot.MappingConfidence == "unknown" {
		status.MappingDriftCount++
	}
	status.LastObservedAt = laterRFC3339(status.LastObservedAt, snapshot.At)
}

func reassessmentDebt(status Status) int {
	return status.InvalidSnapshotCount +
		status.MissingEvidenceCount +
		status.StaleSnapshotCount +
		status.LowQualityCount +
		status.BlockedQualityCount +
		status.MappingDriftCount
}

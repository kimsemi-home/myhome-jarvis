package evidencequality

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func normalizeSnapshot(policy Policy, snapshot Snapshot) (Snapshot, error) {
	normalized := Snapshot{
		ID:                  publicText(snapshot.ID),
		At:                  publicText(snapshot.At),
		EvidenceRef:         filepath.ToSlash(strings.TrimSpace(snapshot.EvidenceRef)),
		Purpose:             normalizeToken(snapshot.Purpose),
		QualityLevel:        normalizeToken(snapshot.QualityLevel),
		SchemaVersion:       publicText(snapshot.SchemaVersion),
		OntologyVersion:     publicText(snapshot.OntologyVersion),
		MappingConfidence:   normalizeToken(snapshot.MappingConfidence),
		AssessedBy:          normalizeToken(snapshot.AssessedBy),
		ReassessmentReasons: normalizeList(snapshot.ReassessmentReasons),
	}
	if err := validateSnapshot(policy, normalized); err != nil {
		return Snapshot{}, err
	}
	return normalized, nil
}

func validateSnapshot(policy Policy, snapshot Snapshot) error {
	if snapshot.At == "" {
		return fmt.Errorf("evidence quality snapshot at timestamp is required")
	}
	if _, err := time.Parse(time.RFC3339, snapshot.At); err != nil {
		return fmt.Errorf("evidence quality snapshot at timestamp is invalid")
	}
	if snapshot.EvidenceRef == "" {
		return errMissingEvidenceRef
	}
	if err := validateRef(policy, snapshot.EvidenceRef); err != nil {
		return err
	}
	if err := rejectSensitiveText(snapshot.EvidenceRef); err != nil {
		return err
	}
	return validateSnapshotEnums(policy, snapshot)
}

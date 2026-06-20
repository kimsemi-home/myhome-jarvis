package evidence

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func inspectIntegrityLine(root string, policy Policy, status *IntegrityStatus, line string) error {
	if strings.TrimSpace(line) == "" {
		return nil
	}
	var observation learningObservation
	if err := json.Unmarshal([]byte(line), &observation); err != nil {
		return err
	}
	for _, ref := range observation.EvidenceRefs {
		if err := inspectIntegrityRef(root, policy, status, ref); err != nil {
			return err
		}
	}
	return nil
}

func inspectIntegrityRef(root string, policy Policy, status *IntegrityStatus, ref string) error {
	normalized, err := normalizeEvidenceRef(policy, ref)
	if err != nil {
		return err
	}
	prefixIndex := integrityPrefixIndex(status.PrefixCounts, normalized)
	if prefixIndex < 0 {
		return nil
	}
	status.CheckedEvidenceRefCount++
	status.PrefixCounts[prefixIndex].CheckedCount++
	_, statErr := os.Stat(filepath.Join(root, filepath.FromSlash(normalized)))
	if errors.Is(statErr, os.ErrNotExist) {
		status.DanglingEvidenceRefCount++
		status.PrefixCounts[prefixIndex].DanglingCount++
		return nil
	}
	if statErr != nil {
		return statErr
	}
	status.PresentEvidenceRefCount++
	status.PrefixCounts[prefixIndex].PresentCount++
	return nil
}

func integrityPrefixIndex(prefixes []IntegrityPrefix, ref string) int {
	for index, prefix := range prefixes {
		if strings.HasPrefix(ref, prefix.Prefix) {
			return index
		}
	}
	return -1
}

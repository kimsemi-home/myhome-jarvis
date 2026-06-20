package codexcost

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func validateAttributionCore(policy Policy, record AttributionRecord) error {
	if record.At == "" || record.Amount <= 0 || record.Basis == "" {
		return fmt.Errorf("codex cost attribution core fields are required")
	}
	if _, err := time.Parse(time.RFC3339, record.At); err != nil {
		return fmt.Errorf("codex cost attribution time must be RFC3339: %w", err)
	}
	if !contains(normalizeList(policy.LoopScopes), record.Scope) {
		return fmt.Errorf("codex cost attribution scope %q is not allowed", record.Scope)
	}
	if !contains(normalizeList(policy.UnitKinds), record.UnitKind) {
		return fmt.Errorf("codex cost attribution unit kind %q is not allowed", record.UnitKind)
	}
	return validateSubjectKey(policy, record.SubjectKey)
}

func validateSubjectKey(policy Policy, subject string) error {
	if subject == "" || len(subject) > policy.AttributionSubjectMaxLength {
		return fmt.Errorf("codex cost attribution subject is invalid")
	}
	if strings.Contains(subject, "://") || strings.Contains(subject, "\\") ||
		strings.Contains(subject, "..") ||
		filepath.IsAbs(filepath.FromSlash(subject)) {
		return fmt.Errorf("codex cost attribution subject must be a safe key")
	}
	return nil
}

func validateCostRef(policy Policy, costRef string) error {
	if costRef == "" || len(costRef) > policy.AttributionCostRefMaxLength {
		return fmt.Errorf("codex cost attribution cost ref is invalid")
	}
	for _, char := range costRef {
		if !safeCostRefChar(char) {
			return fmt.Errorf("codex cost attribution cost ref must be a safe token")
		}
	}
	return nil
}

func safeCostRefChar(char rune) bool {
	return char >= 'a' && char <= 'z' ||
		char >= '0' && char <= '9' ||
		char == '_' || char == '-' || char == ':'
}

func validateAttributionRefs(policy Policy, refs []string) error {
	if len(refs) == 0 {
		return errMissingEvidenceRef
	}
	for _, ref := range refs {
		if err := validateRef(policy, ref); err != nil {
			return err
		}
	}
	return nil
}

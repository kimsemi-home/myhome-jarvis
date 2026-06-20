package codexcost

import (
	"fmt"
	"strings"
	"time"
)

func normalizeRecord(policy Policy, record Record) (Record, error) {
	record.At = strings.TrimSpace(record.At)
	record.Scope = normalizeToken(record.Scope)
	record.UnitKind = normalizeToken(record.UnitKind)
	record.Status = normalizeToken(record.Status)
	record.EvidenceRefs = normalizeRefs(record.EvidenceRefs)
	if record.At == "" || record.Amount <= 0 {
		return Record{}, fmt.Errorf("codex cost time and positive amount are required")
	}
	if _, err := time.Parse(time.RFC3339, record.At); err != nil {
		return Record{}, fmt.Errorf("codex cost time must be RFC3339: %w", err)
	}
	if !contains(normalizeList(policy.LoopScopes), record.Scope) {
		return Record{}, fmt.Errorf("codex cost scope %q is not allowed", record.Scope)
	}
	if !contains(normalizeList(policy.UnitKinds), record.UnitKind) {
		return Record{}, fmt.Errorf("codex cost unit kind %q is not allowed", record.UnitKind)
	}
	if !contains(normalizeList(policy.RecordStatuses), record.Status) {
		return Record{}, fmt.Errorf("codex cost status %q is not allowed", record.Status)
	}
	if len(record.EvidenceRefs) == 0 {
		return Record{}, errMissingEvidenceRef
	}
	for _, ref := range record.EvidenceRefs {
		if err := validateRef(policy, ref); err != nil {
			return Record{}, err
		}
	}
	return record, nil
}

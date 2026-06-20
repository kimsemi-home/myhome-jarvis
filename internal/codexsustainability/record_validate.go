package codexsustainability

import (
	"fmt"
	"strings"
	"time"
)

func normalizeRecord(policy Policy, record Record) (Record, error) {
	record.At = strings.TrimSpace(record.At)
	record.RecordKind = normalizeToken(record.RecordKind)
	record.Metric = normalizeToken(record.Metric)
	record.EvidenceRefs = normalizeRefs(record.EvidenceRefs)
	if record.At == "" {
		return Record{}, fmt.Errorf("codex sustainability time is required")
	}
	if _, err := time.Parse(time.RFC3339, record.At); err != nil {
		return Record{}, fmt.Errorf("codex sustainability time must be RFC3339")
	}
	if !contains(normalizeList(policy.RecordKinds), record.RecordKind) {
		return Record{}, fmt.Errorf("codex sustainability kind %q is not allowed", record.RecordKind)
	}
	if len(record.EvidenceRefs) == 0 {
		return Record{}, errMissingEvidenceRef
	}
	for _, ref := range record.EvidenceRefs {
		if err := validateRef(policy, ref); err != nil {
			return Record{}, err
		}
	}
	return normalizeByKind(policy, record)
}

func normalizeByKind(policy Policy, record Record) (Record, error) {
	switch record.RecordKind {
	case "usage_sample", "cycle_sample":
		if !contains(normalizeList(policy.Metrics), record.Metric) || record.Amount <= 0 {
			return Record{}, fmt.Errorf("codex sustainability metric sample is invalid")
		}
	case "trend_baseline":
		if record.TrendBaselineVersion == "" || record.TrendMeasuredAt == "" ||
			record.Metric != "elapsed_cycle_minutes" || record.Amount <= 0 {
			return Record{}, fmt.Errorf("codex sustainability trend baseline is invalid")
		}
	case "feature_proposal":
		if err := validateProposal(record); err != nil {
			return Record{}, err
		}
	}
	return record, nil
}

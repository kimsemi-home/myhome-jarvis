package codexcost

import (
	"strings"
	"time"
)

func normalizeUsageRecord(policy Policy, request RecordRequest, now time.Time) (Record, error) {
	recordedAt, err := normalizeRecordedAt(request.At, now)
	if err != nil {
		return Record{}, err
	}
	record := Record{
		At:           recordedAt,
		Scope:        normalizeToken(request.Scope),
		UnitKind:     normalizeToken(request.UnitKind),
		Amount:       request.Amount,
		Status:       normalizeToken(request.Status),
		EvidenceRefs: normalizeRefs(request.EvidenceRefs),
	}
	if record.Status == "" {
		record.Status = "recorded"
	}
	if record.Amount >= policy.ReviewUnitThreshold {
		record.Status = "review_required"
	}
	normalized, err := normalizeRecord(policy, record)
	if err != nil {
		return Record{}, err
	}
	normalized.SemanticHash = usageSemanticHash(normalized)
	return normalized, nil
}

func normalizeRecordedAt(value string, now time.Time) (string, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return now.UTC().Format(time.RFC3339), nil
	}
	parsed, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return "", err
	}
	return parsed.UTC().Format(time.RFC3339), nil
}

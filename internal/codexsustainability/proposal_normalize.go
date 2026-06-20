package codexsustainability

import (
	"strings"
	"time"
)

func normalizeProposalRecord(
	policy Policy,
	request ProposalRecordRequest,
	now time.Time,
) (Record, error) {
	recordedAt, err := normalizeRecordedAt(request.At, now)
	if err != nil {
		return Record{}, err
	}
	record := Record{
		At:                    recordedAt,
		RecordKind:            "feature_proposal",
		ProposalID:            publicKey(request.ProposalID),
		CostPerAcceptedChange: request.CostPerAcceptedChange,
		MedianCycleMinutes:    request.MedianCycleMinutes,
		CacheSavingsUnits:     request.CacheSavingsUnits,
		DefectReworkRate:      request.DefectReworkRate,
		MonetizationRef:       publicKey(request.MonetizationRef),
		EvidenceRefs:          normalizeRefs(request.EvidenceRefs),
	}
	return normalizeRecord(policy, record)
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

func publicKey(value string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(value)), "-")
}

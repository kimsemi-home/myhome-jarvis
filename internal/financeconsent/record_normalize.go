package financeconsent

import (
	"fmt"
	"time"
)

func normalizeRecordRequest(
	policy Policy,
	request RecordRequest,
	now time.Time,
) (Record, error) {
	recordedAt, err := normalizeRecordedAt(request.At, now)
	if err != nil {
		return Record{}, err
	}
	expiresAt, err := normalizeOptionalTime(request.ExpiresAt)
	if err != nil {
		return Record{}, err
	}
	record := Record{
		At:               recordedAt,
		ConsentKind:      normalizeToken(request.ConsentKind),
		SubjectScope:     normalizeToken(request.SubjectScope),
		Status:           normalizeToken(request.Status),
		ReviewStatus:     normalizeToken(request.ReviewStatus),
		AuthorityProfile: normalizeToken(request.AuthorityProfile),
		EvidenceRefs:     normalizeRefs(request.EvidenceRefs),
		ExpiresAt:        expiresAt,
	}
	if !privateSafeScope(record.SubjectScope) ||
		!scopeMatchesKind(record.ConsentKind, record.SubjectScope) ||
		!validateRecord(record, policy).Active {
		return Record{}, fmt.Errorf("finance consent record is not active and public-safe")
	}
	return record, nil
}

func privateSafeScope(value string) bool {
	if value == "" {
		return false
	}
	for _, char := range value {
		if (char < 'a' || char > 'z') && char != '_' && char != '-' {
			return false
		}
	}
	return true
}

func scopeMatchesKind(kind string, scope string) bool {
	allowed := map[string][]string{
		"finance_connector": {"user", "owner_read_only"},
		"spouse_scope":      {"spouse", "spouse_read_only"},
		"household_scope":   {"household", "household_read_only"},
	}
	return contains(allowed[kind], scope)
}

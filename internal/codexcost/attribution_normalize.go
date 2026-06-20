package codexcost

import (
	"strings"
	"time"
)

func normalizeAttribution(
	policy Policy,
	request AttributionRequest,
	now time.Time,
) (AttributionRecord, error) {
	record := AttributionRecord{
		At: request.At, Scope: request.Scope, SubjectKey: request.SubjectKey,
		CostRef: request.CostRef, UnitKind: request.UnitKind,
		Amount: request.Amount, Basis: request.Basis, EvidenceRefs: request.EvidenceRefs,
	}
	if strings.TrimSpace(record.At) == "" {
		record.At = now.Format(time.RFC3339)
	}
	normalized, err := normalizeAttributionRecord(policy, record)
	if err != nil {
		return AttributionRecord{}, err
	}
	normalized.SubjectHash = attributionSubjectHash(normalized.SubjectKey)
	normalized.SemanticHash = attributionSemanticHash(normalized)
	return normalized, nil
}

func normalizeAttributionRecord(
	policy Policy,
	record AttributionRecord,
) (AttributionRecord, error) {
	record.At = strings.TrimSpace(record.At)
	record.Scope = normalizeToken(record.Scope)
	record.SubjectKey = strings.TrimSpace(record.SubjectKey)
	record.CostRef = normalizeToken(record.CostRef)
	record.UnitKind = normalizeToken(record.UnitKind)
	record.Basis = normalizeToken(record.Basis)
	record.EvidenceRefs = normalizeRefs(record.EvidenceRefs)
	if err := validateAttributionCore(policy, record); err != nil {
		return AttributionRecord{}, err
	}
	if strings.TrimSpace(record.SubjectHash) == "" {
		record.SubjectHash = attributionSubjectHash(record.SubjectKey)
	}
	if strings.TrimSpace(record.CostRef) == "" {
		record.CostRef = attributionCostRef(record)
	}
	if err := validateCostRef(policy, record.CostRef); err != nil {
		return AttributionRecord{}, err
	}
	return record, validateAttributionRefs(policy, record.EvidenceRefs)
}

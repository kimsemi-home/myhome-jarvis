package codexcost

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

func attributionSemanticHash(record AttributionRecord) string {
	sum := sha256.Sum256([]byte(strings.Join([]string{
		record.CostRef,
		record.Scope,
		record.SubjectKey,
		record.UnitKind,
		strconv.FormatInt(record.Amount, 10),
		record.Basis,
		strings.Join(record.EvidenceRefs, "\x1f"),
	}, "\x00")))
	return "cost_attr_" + hex.EncodeToString(sum[:])[:16]
}

func attributionSubjectHash(subject string) string {
	sum := sha256.Sum256([]byte("subject\x00" + subject))
	return "subject_" + hex.EncodeToString(sum[:])[:16]
}

func attributionCostRef(record AttributionRecord) string {
	sum := sha256.Sum256([]byte(strings.Join([]string{
		record.UnitKind,
		strconv.FormatInt(record.Amount, 10),
		strings.Join(record.EvidenceRefs, "\x1f"),
	}, "\x00")))
	return "cost_ref_" + hex.EncodeToString(sum[:])[:16]
}

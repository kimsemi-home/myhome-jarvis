package codexcost

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

func usageSemanticHash(record Record) string {
	sum := sha256.Sum256([]byte(strings.Join([]string{
		record.Scope,
		record.UnitKind,
		strconv.FormatInt(record.Amount, 10),
		strings.Join(record.EvidenceRefs, "\x1f"),
	}, "\x00")))
	return "cost_" + hex.EncodeToString(sum[:])[:16]
}

package localfinanceevidence

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"
)

func receiptHash(receipt Receipt) string {
	values := []string{receipt.SchemaVersion, receipt.Component, receipt.Capability,
		receipt.ExecutionMode, "false", receipt.ArtifactHash}
	if receipt.ExternalWritesEnabled {
		values[4] = "true"
	}
	return digest(strings.Join(append(values, receipt.Checks...), "\n"))
}

func aggregateHash(manifest Manifest) string {
	receipts := append([]Receipt{}, manifest.Receipts...)
	sort.Slice(receipts, func(i, j int) bool { return receipts[i].Component < receipts[j].Component })
	values := []string{manifest.SchemaVersion, manifest.Month, "false"}
	if manifest.ExternalWritesEnabled {
		values[2] = "true"
	}
	for _, receipt := range receipts {
		values = append(values, receipt.Component+"\x00"+receipt.ReceiptHash)
	}
	return digest(strings.Join(values, "\n"))
}

func digest(value string) string {
	hash := sha256.Sum256([]byte(value))
	return hex.EncodeToString(hash[:])
}

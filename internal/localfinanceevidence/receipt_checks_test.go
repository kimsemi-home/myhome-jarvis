package localfinanceevidence

import (
	"path/filepath"
	"slices"
	"testing"
)

func TestShortsReceiptRequiresConnectionReadiness(t *testing.T) {
	path := filepath.Join("..", "..", "fixtures", "local_finance", "manifest.json")
	manifest, err := Read(path)
	if err != nil {
		t.Fatal(err)
	}
	for index := range manifest.Receipts {
		if manifest.Receipts[index].Component != "shorts" {
			continue
		}
		manifest.Receipts[index].Checks = slices.DeleteFunc(manifest.Receipts[index].Checks, func(value string) bool {
			return value == "connection-readiness-plan-sealed"
		})
		manifest.Receipts[index].ReceiptHash = receiptHash(manifest.Receipts[index])
	}
	manifest.AggregateHash = aggregateHash(manifest)
	if err := Validate(manifest); err == nil {
		t.Fatal("accepted Shorts evidence without the sealed connection-readiness plan")
	}
}

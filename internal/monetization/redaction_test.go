package monetization

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestStatusDoesNotExposePrivateRevenueFields(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/monetization/experiments.jsonl", ledgerFixture())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range [][]byte{
		[]byte("private revenue note"),
		[]byte("private_revenue_notes"),
		[]byte("evidence_refs"),
		[]byte("raw_revenue_amount"),
	} {
		if bytes.Contains(body, forbidden) {
			t.Fatalf("status leaked %s in %s", forbidden, body)
		}
	}
}

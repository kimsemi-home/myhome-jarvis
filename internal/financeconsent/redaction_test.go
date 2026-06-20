package financeconsent

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestStatusDoesNotExposePrivateConsentPayload(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/finance/consent.jsonl", activeLedgerFixture())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range [][]byte{
		[]byte("private consent note"),
		[]byte("private_notes"),
		[]byte("evidence_refs"),
		[]byte("account_id"),
		[]byte("card_number"),
		[]byte("raw_transaction"),
	} {
		if bytes.Contains(body, forbidden) {
			t.Fatalf("finance consent status leaked %s in %s", forbidden, body)
		}
	}
}

package financeconsent

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRecordConsentAppendsPrivateRecordAndReturnsRedactedResult(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	result, err := RecordConsent(root, []byte(`{
		"consent_kind":"finance_connector",
		"subject_scope":"owner_read_only",
		"status":"granted",
		"review_status":"approved",
		"authority_profile":"finance_review_only",
		"evidence_refs":["docs/finance-consent.md"]
	}`))
	if err != nil {
		t.Fatal(err)
	}
	if result.ConsentKind != "finance_connector" ||
		result.SubjectScope != "owner_read_only" {
		t.Fatalf("result = %#v", result)
	}
	if result.EvidenceRefCount != 1 || result.MissingRequiredConsentCount != 2 {
		t.Fatalf("redacted result = %#v", result)
	}
	body := readFinanceConsentLedger(t, root)
	if !bytes.Contains(body, []byte(`"evidence_refs"`)) {
		t.Fatalf("ledger missing private evidence refs: %s", body)
	}
	if bytes.Contains(mustFinanceConsentJSON(t, result), []byte("evidence_refs")) {
		t.Fatalf("result leaked evidence refs: %#v", result)
	}
}

func TestRecordConsentCanReachReadyReadOnly(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	for _, payload := range activeConsentPayloads() {
		if _, err := RecordConsent(root, []byte(payload)); err != nil {
			t.Fatal(err)
		}
	}
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.ReadinessState != "ready_read_only" ||
		status.ActiveConsentCount != 3 ||
		status.ConsentDebtCount != 0 {
		t.Fatalf("status = %#v", status)
	}
}

func readFinanceConsentLedger(t *testing.T, root string) []byte {
	t.Helper()
	body, err := os.ReadFile(filepath.Join(root, "data/private/finance/consent.jsonl"))
	if err != nil {
		t.Fatal(err)
	}
	return body
}

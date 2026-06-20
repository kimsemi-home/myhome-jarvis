package financeconsent

import "testing"

func TestMissingLedgerReturnsBlockedPublicStatus(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.Exists || status.ReadinessState != "blocked" {
		t.Fatalf("status = %#v", status)
	}
	if status.MissingRequiredConsentCount != 3 || status.ConsentDebtCount != 3 {
		t.Fatalf("consent debt = %#v", status)
	}
}

func TestActiveLedgerIsReadyButStillReadOnly(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/finance/consent.jsonl", activeLedgerFixture())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.ReadinessState != "ready_read_only" || status.ActiveConsentCount != 3 {
		t.Fatalf("status = %#v", status)
	}
	if status.ForbiddenActionEnabledCount != 0 || status.FinanceMode != "read_only_review_only" {
		t.Fatalf("finance mode = %#v", status)
	}
}

func TestDebtLedgerCountsReviewAndMissingEvidence(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/finance/consent.jsonl", debtLedgerFixture())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.ReadinessState != "blocked" || status.InvalidRecordCount != 1 {
		t.Fatalf("status = %#v", status)
	}
	if status.ReviewRequiredCount != 1 || status.MissingEvidenceCount != 1 ||
		status.RevokedOrExpiredCount != 1 {
		t.Fatalf("debt counts = %#v", status)
	}
}

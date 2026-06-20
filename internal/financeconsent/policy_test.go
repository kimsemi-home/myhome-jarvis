package financeconsent

import "testing"

func TestPolicyRequiresPrivateLedgerAndReadOnlyReviewOnly(t *testing.T) {
	policy := testPolicy()
	policy.PrivateConsentLedger = "generated/consent.jsonl"
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected public ledger path to fail")
	}

	policy = testPolicy()
	policy.ReviewOnly = false
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected non-review-only finance to fail")
	}
}

func TestPolicyRejectsWriteActions(t *testing.T) {
	policy := testPolicy()
	policy.TransferActionsAllowed = true
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected transfer action to fail")
	}

	policy = testPolicy()
	policy.ExternalWritesAllowed = true
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected external writes to fail")
	}
}

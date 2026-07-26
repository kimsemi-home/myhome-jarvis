package financeconnector

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadConfiguredBlocksLiveWithoutActiveConsent(t *testing.T) {
	t.Setenv("MYHOME_FINANCE_MODE", ModeMyData)
	t.Setenv("MYHOME_FINANCE_ALLOW_EXTERNAL", "true")
	t.Setenv("MYHOME_FINANCE_1PASSWORD_REF", "op://vault/item/token")
	t.Setenv("MYHOME_MYDATA_BASE_URL", "http://127.0.0.1:1")
	loaded, err := LoadConfigured(context.Background(), repoRootForConsentTest(t))
	if err == nil || loaded.Transactions != nil || !strings.Contains(err.Error(), "ready_read_only") {
		t.Fatalf("loaded = %#v err = %v", loaded, err)
	}
}

func repoRootForConsentTest(t *testing.T) string {
	repo, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	repo = filepath.Join(repo, "..", "..")
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "generated"), 0o755); err != nil {
		t.Fatal(err)
	}
	policy, err := os.ReadFile(filepath.Join(repo, "generated", "finance_consent.generated.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "generated", "finance_consent.generated.json"), policy, 0o644); err != nil {
		t.Fatal(err)
	}
	return root
}

func consentRootWithActiveLedger(t *testing.T) string {
	root := repoRootForConsentTest(t)
	ledger := filepath.Join(root, "data", "private", "finance", "consent.jsonl")
	if err := os.MkdirAll(filepath.Dir(ledger), 0o700); err != nil {
		t.Fatal(err)
	}
	body := `{"at":"2026-06-19T00:00:00Z","consent_kind":"finance_connector","subject_scope":"user","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["docs/finance-consent.md"]}
{"at":"2026-06-19T00:01:00Z","consent_kind":"spouse_scope","subject_scope":"spouse","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["docs/finance-consent.md"]}
{"at":"2026-06-19T00:02:00Z","consent_kind":"household_scope","subject_scope":"household","status":"granted","review_status":"approved","authority_profile":"finance_review_only","evidence_refs":["docs/finance-consent.md"]}
`
	if err := os.WriteFile(ledger, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
	return root
}

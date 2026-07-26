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

package security

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCheckHistoryAllowsPrivateKeepPlaceholder(t *testing.T) {
	root := initGitRepo(t)
	privateDir := filepath.Join(root, "data", "private")
	if err := os.MkdirAll(privateDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(privateDir, ".keep"), []byte(""), 0o600); err != nil {
		t.Fatal(err)
	}
	commitAll(t, root, "add private placeholder")
	report, err := CheckHistory(root)
	if err != nil {
		t.Fatal(err)
	}
	if !report.OK {
		t.Fatalf("expected private placeholder to be allowed, got %+v", report.Findings)
	}
}

func TestCheckHistoryAllowsRedactedTokenPlaceholder(t *testing.T) {
	root := initGitRepo(t)
	body := []byte("request := \"/health?local_token=redacted-value\"\n")
	if err := os.WriteFile(filepath.Join(root, "server_test.go"), body, 0o600); err != nil {
		t.Fatal(err)
	}
	commitAll(t, root, "add redacted placeholder")
	report, err := CheckHistory(root)
	if err != nil {
		t.Fatal(err)
	}
	if !report.OK {
		t.Fatalf("expected redacted token placeholder to be allowed, got %+v", report.Findings)
	}
}

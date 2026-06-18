package security

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckHistoryRejectsPrivateMarkerInCommittedContent(t *testing.T) {
	root := initGitRepo(t)
	privatePath := "/" + "Users" + "/" + strings.Join([]string{"al", "ice"}, "") + "/project"
	if err := os.WriteFile(filepath.Join(root, "notes.md"), []byte("old path: "+privatePath+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	commitAll(t, root, "add private marker")
	report, err := CheckHistory(root)
	if err != nil {
		t.Fatal(err)
	}
	if report.OK {
		t.Fatal("expected private marker in history to be rejected")
	}
	if !historyHasCode(report, "history_private_identity") {
		t.Fatalf("expected private identity finding, got %+v", report.Findings)
	}
}

func TestCheckHistoryRejectsPrivateDataPath(t *testing.T) {
	root := initGitRepo(t)
	privateDir := filepath.Join(root, "data", "private")
	if err := os.MkdirAll(privateDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(privateDir, "local-token.txt"), []byte("local only\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	commitAll(t, root, "add private file")
	report, err := CheckHistory(root)
	if err != nil {
		t.Fatal(err)
	}
	if report.OK {
		t.Fatal("expected private data path in history to be rejected")
	}
	if !historyHasCode(report, "history_private_data_path") {
		t.Fatalf("expected private data path finding, got %+v", report.Findings)
	}
}

package security

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStatusForRootAggregatesCurrentAndHistoryWithoutRoot(t *testing.T) {
	root := initGitRepo(t)
	if err := os.WriteFile(filepath.Join(root, "README.md"), []byte("ok\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	commitAll(t, root, "initial")
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.OK || !status.CurrentOK || !status.HistoryOK {
		t.Fatalf("expected clean status, got %#v", status)
	}
	if status.CurrentFindingCount != 0 || status.HistoryFindingCount != 0 {
		t.Fatalf("unexpected finding counts: %#v", status)
	}
	if status.CheckedAt == "" {
		t.Fatal("expected checked_at to be set")
	}
}

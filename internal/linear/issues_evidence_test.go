package linear

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteEvidenceRedactsNonIssueKeys(t *testing.T) {
	root := t.TempDir()
	if err := AppendWriteEvidence(root, "linear_comment", "550e8400-e29b-41d4-a716-446655440000"); err != nil {
		t.Fatal(err)
	}
	status, err := WriteEvidenceStatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.LatestSyncedMutation == nil {
		t.Fatalf("expected latest write evidence: %#v", status)
	}
	if status.LatestSyncedMutation.IssueKey != "" {
		t.Fatalf("non-issue key leaked into evidence: %#v", status.LatestSyncedMutation)
	}
	payload, err := os.ReadFile(filepath.Join(root, WriteEvidenceRelativePath))
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(payload), "550e8400") {
		t.Fatalf("raw id leaked into evidence file: %s", string(payload))
	}
}

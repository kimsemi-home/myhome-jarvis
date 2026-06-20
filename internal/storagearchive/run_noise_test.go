package storagearchive

import (
	"path/filepath"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

func TestRunBlocksArchiveOnNoiseBudgetBreach(t *testing.T) {
	root := t.TempDir()
	source := domain.PrivateLogSource{Key: "quality", Path: "data/private/quality/runs.jsonl", Format: "jsonl"}
	policy := testStoragePolicy(source)
	writeStoragePolicy(t, root, policy)
	writePrivateFile(t, root, source.Path,
		`{"source":"quality","kind":"run","evidence_ref":"a"}`+"\n"+
			`{"source":"quality","kind":"run","evidence_ref":"a"}`+"\n")
	report, err := RunForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	result := report.Results[0]
	if result.State != "budget_breach" || result.ArchivePath != "" {
		t.Fatalf("expected budget breach without archive, got %#v", result)
	}
	entry := readManifestEntry(t, root, report.ManifestPath)
	if entry.BudgetVerdict != "breach" || entry.RawPayloadStored {
		t.Fatalf("manifest breach entry = %#v", entry)
	}
	matches, err := filepath.Glob(filepath.Join(root, "data/private/archive/quality-*.jsonl.gz"))
	if err != nil {
		t.Fatal(err)
	}
	if len(matches) != 0 {
		t.Fatalf("unexpected archive artifacts: %v", matches)
	}
}

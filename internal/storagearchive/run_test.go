package storagearchive

import (
	"strings"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

func TestRunSkipsMissingSourceSafely(t *testing.T) {
	root := t.TempDir()
	source := domain.PrivateLogSource{Key: "quality", Path: "data/private/quality/runs.jsonl", Format: "jsonl"}
	writeStoragePolicy(t, root, testStoragePolicy(source))
	report, err := RunForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if report.ArchivedCount != 0 || report.SkippedCount != 1 || !report.PublicSafe {
		t.Fatalf("missing source report = %#v", report)
	}
	if report.Results[0].State != "missing" || report.Results[0].ArchivePath != "" {
		t.Fatalf("missing source result = %#v", report.Results[0])
	}
}

func TestRunArchivesJSONLAndAppendsManifest(t *testing.T) {
	root := t.TempDir()
	source := domain.PrivateLogSource{Key: "quality", Path: "data/private/quality/runs.jsonl", Format: "jsonl"}
	writeStoragePolicy(t, root, testStoragePolicy(source))
	writePrivateFile(t, root, source.Path, `{"source":"quality","kind":"run","evidence_ref":"a"}`+"\n")
	report, err := RunForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	result := report.Results[0]
	if result.State != "archived" || result.ArchivePath == "" || !result.BudgetOK {
		t.Fatalf("archive result = %#v", result)
	}
	if result.CompressionRatioPercent <= 0 || report.CompressionRatioPercent <= 0 {
		t.Fatalf("compression evidence missing = %#v / %#v", result, report)
	}
	if result.ConfigEvidenceSHA256 == "" || report.ConfigEvidenceSHA256 == "" {
		t.Fatalf("config evidence missing = %#v / %#v", result, report)
	}
	assertGzipContains(t, root, result.ArchivePath, `"source":"quality"`)
	entry := readManifestEntry(t, root, report.ManifestPath)
	if entry.ArchivePath != result.ArchivePath || entry.RawPayloadStored != true {
		t.Fatalf("manifest entry = %#v", entry)
	}
	if entry.CompressionRatioPercent <= 0 || entry.ConfigEvidenceSHA256 == "" {
		t.Fatalf("manifest evidence = %#v", entry)
	}
	if strings.Contains(mustJSON(t, report), "local-token") {
		t.Fatalf("public report leaked private marker: %#v", report)
	}
}

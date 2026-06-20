package codexsustainability

import (
	"encoding/json"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
	"github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"
)

func TestCaptureQualityRunRecordsStorageArchiveCacheMetrics(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	sourceA := domain.PrivateLogSource{
		Key: "quality_a", Path: "data/private/archive-source/a.jsonl", Format: "jsonl",
	}
	sourceB := domain.PrivateLogSource{
		Key: "quality_b", Path: "data/private/archive-source/b.jsonl", Format: "jsonl",
	}
	writeStoragePolicy(t, root, sourceA, sourceB)
	writeFile(t, root, sourceA.Path, `{"source":"a","kind":"run","evidence_ref":"a"}`+"\n")
	if _, err := storagearchive.RunForRoot(root); err != nil {
		t.Fatal(err)
	}
	writeFile(t, root, sourceB.Path, `{"source":"b","kind":"run","evidence_ref":"b"}`+"\n")
	writeQualityRun(t, root, true)
	capture, err := captureQualityRunAt(root, mustTime(t, "2026-06-20T01:00:00Z"))
	if err != nil {
		t.Fatal(err)
	}
	if capture.RecordedRecordCount != 5 || capture.StorageCacheHitCount != 1 ||
		capture.StorageCacheMissCount != 1 || capture.StorageCacheSavingsUnits <= 0 {
		t.Fatalf("capture cache metrics = %#v", capture)
	}
	status, err := statusForRootAt(root, mustTime(t, "2026-06-20T01:00:00Z"))
	if err != nil {
		t.Fatal(err)
	}
	if status.CacheHitCount != 1 || status.CacheMissCount != 1 ||
		status.CacheSavingsUnits != capture.StorageCacheSavingsUnits {
		t.Fatalf("status cache metrics = %#v", status)
	}
}

func writeStoragePolicy(t *testing.T, root string, sources ...domain.PrivateLogSource) {
	t.Helper()
	policy := domain.StoragePolicy{
		PrivateLogSources: sources,
		LogArchive: domain.LogArchivePolicy{
			Mode: "compress_then_archive", Compression: "gzip",
			ArchiveRoot: "data/private/archive", ArchiveExtension: ".jsonl.gz",
			ManifestPath:     "data/private/archive/manifest.jsonl",
			ConfigIsEvidence: true, ConfigHashInputs: []string{
				"private_log_sources", "log_archive", "evidence_noise_budget",
			},
		},
		EvidenceNoiseBudget: domain.EvidenceNoiseBudget{
			Enabled: true, MaxNoiseRatioPercent: 20,
			MaxLowSignalRecordsPerWindow: 10,
			DedupeKeyFields:              []string{"source", "kind", "evidence_ref"},
			ConfigEvidenceField:          "evidence_noise_budget", BreachBlocksArchive: true,
		},
	}
	body, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, root, storagearchive.PolicyRelativePath, string(body)+"\n")
}

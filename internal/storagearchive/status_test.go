package storagearchive

import "testing"

func TestStatusSummarizesArchivePolicy(t *testing.T) {
	status, err := StatusForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	if !status.PublicSafe || !status.ArchiveReady || !status.NoiseBudgetReady {
		t.Fatalf("archive status = %#v", status)
	}
	if status.Compression != "gzip" ||
		status.CompressionArchivePattern != "compress_then_archive" {
		t.Fatalf("compression policy = %#v", status)
	}
	if status.PrivateLogSourceCount != 12 || status.MaxNoiseRatioPercent > 25 {
		t.Fatalf("archive counts = %#v", status)
	}
	assertStorageArchiveSourceKeys(t, status.PrivateLogSourceKeys)
	if status.ConfigEvidenceField != "evidence_noise_budget" ||
		!status.BreachBlocksArchive {
		t.Fatalf("noise budget evidence = %#v", status)
	}
	if status.ConfigEvidenceSHA256 == "" || len(status.ConfigHashInputs) != 3 {
		t.Fatalf("config evidence hash = %#v", status)
	}
	if status.ManifestBudgetBreachCount != 0 ||
		status.ManifestInvalidEntryCount != 0 {
		t.Fatalf("manifest health = %#v", status)
	}
}

func TestStatusSummarizesArchiveManifest(t *testing.T) {
	root := t.TempDir()
	source := privateQualitySource()
	writeStoragePolicy(t, root, testStoragePolicy(source))
	writePrivateFile(t, root, source.Path,
		`{"source":"quality","kind":"run","evidence_ref":"a"}`+"\n")
	if _, err := RunForRoot(root); err != nil {
		t.Fatal(err)
	}
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.ManifestPresent ||
		status.ManifestEntryCount != 1 ||
		status.ManifestArchivedCount != 1 {
		t.Fatalf("manifest summary = %#v", status)
	}
	if status.ManifestCompressionRatio <= 0 ||
		status.ManifestLastArchivedAt == "" {
		t.Fatalf("manifest compression evidence = %#v", status)
	}
	if status.ManifestBudgetBreachCount != 0 ||
		status.ManifestInvalidEntryCount != 0 {
		t.Fatalf("manifest noise health = %#v", status)
	}
}

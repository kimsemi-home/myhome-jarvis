package storagearchive

import "testing"

func TestRunReportsPolicyEvidence(t *testing.T) {
	root := t.TempDir()
	source := privateQualitySource()
	writeStoragePolicy(t, root, testStoragePolicy(source))

	report, err := RunForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	evidence := report.PolicyEvidence
	if evidence.CompressionArchivePattern != "compress_then_archive" {
		t.Fatalf("archive pattern = %#v", evidence)
	}
	if evidence.Compression != "gzip" || !evidence.NoiseBudgetEnabled {
		t.Fatalf("compression/noise evidence = %#v", evidence)
	}
	if evidence.MaxNoiseRatioPercent != 20 ||
		evidence.MaxLowSignalRecordsPerWindow != 10 {
		t.Fatalf("noise budget thresholds = %#v", evidence)
	}
	if evidence.NoiseBudgetWindow != "per_quality_run" ||
		len(evidence.DedupeKeyFields) != 3 {
		t.Fatalf("noise budget shape = %#v", evidence)
	}
	if evidence.ConfigEvidenceSHA256 != report.ConfigEvidenceSHA256 ||
		evidence.ConfigEvidenceField != "evidence_noise_budget" {
		t.Fatalf("config evidence mismatch = %#v / %#v", evidence, report)
	}
	if !evidence.ConfigIsEvidence || !evidence.BreachBlocksArchive {
		t.Fatalf("evidence policy flags = %#v", evidence)
	}
}

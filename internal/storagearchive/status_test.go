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
	if status.PrivateLogSourceCount != 8 || status.MaxNoiseRatioPercent > 25 {
		t.Fatalf("archive counts = %#v", status)
	}
	if !containsKey(status.PrivateLogSourceKeys, "codex_cost_attribution") {
		t.Fatalf("archive sources = %#v", status.PrivateLogSourceKeys)
	}
	if !containsKey(status.PrivateLogSourceKeys, "monetization") {
		t.Fatalf("archive sources = %#v", status.PrivateLogSourceKeys)
	}
	if status.ConfigEvidenceField != "evidence_noise_budget" ||
		!status.BreachBlocksArchive {
		t.Fatalf("noise budget evidence = %#v", status)
	}
	if status.ConfigEvidenceSHA256 == "" || len(status.ConfigHashInputs) != 3 {
		t.Fatalf("config evidence hash = %#v", status)
	}
}

func containsKey(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}

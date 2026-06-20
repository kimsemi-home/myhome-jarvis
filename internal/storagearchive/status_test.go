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
	if status.PrivateLogSourceCount != 5 || status.MaxNoiseRatioPercent > 25 {
		t.Fatalf("archive counts = %#v", status)
	}
	if status.ConfigEvidenceField != "evidence_noise_budget" ||
		!status.BreachBlocksArchive {
		t.Fatalf("noise budget evidence = %#v", status)
	}
}

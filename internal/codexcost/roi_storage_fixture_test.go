package codexcost

import (
	"encoding/json"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
	"github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"
)

func writeStoragePolicy(t *testing.T, root string) {
	t.Helper()
	policy := domain.StoragePolicy{
		PrivateLogSources: []domain.PrivateLogSource{
			{Key: "codex_cost", Path: "data/private/codex-cost/usage.jsonl", Format: "jsonl"},
			{Key: "codex_cost_attribution", Path: "data/private/codex-cost/attribution.jsonl", Format: "jsonl"},
			{Key: "codex_sustainability", Path: "data/private/codex-sustainability/evidence.jsonl", Format: "jsonl"},
		},
		LogArchive: domain.LogArchivePolicy{
			Mode: "compress_then_archive", Compression: "gzip",
			ArchiveRoot: "data/private/archive", ArchiveExtension: ".jsonl.gz",
			ManifestPath:     "data/private/archive/manifest.jsonl",
			ConfigIsEvidence: true,
			ConfigHashInputs: []string{
				"private_log_sources", "log_archive", "evidence_noise_budget",
			},
		},
		EvidenceNoiseBudget: domain.EvidenceNoiseBudget{
			Enabled: true, MaxNoiseRatioPercent: 20,
			MaxLowSignalRecordsPerWindow: 10,
			ConfigEvidenceField:          "evidence_noise_budget",
			BreachBlocksArchive:          true,
		},
	}
	body, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, root, storagearchive.PolicyRelativePath, string(body)+"\n")
}

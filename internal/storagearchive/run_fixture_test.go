package storagearchive

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

func writeStoragePolicy(t *testing.T, root string, policy domain.StoragePolicy) {
	t.Helper()
	body, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	writePrivateFile(t, root, PolicyRelativePath, string(body)+"\n")
}

func writePrivateFile(t *testing.T, root string, rel string, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}

func testStoragePolicy(source domain.PrivateLogSource) domain.StoragePolicy {
	return domain.StoragePolicy{
		PrivateLogSources: []domain.PrivateLogSource{source},
		LogArchive: domain.LogArchivePolicy{
			Mode:                    "compress_then_archive",
			Compression:             "gzip",
			ArchiveRoot:             "data/private/archive",
			ArchiveExtension:        ".jsonl.gz",
			ManifestPath:            "data/private/archive/manifest.jsonl",
			RawPayloadPublicAllowed: false,
			ConfigIsEvidence:        true,
			ConfigHashInputs:        []string{"log_archive", "evidence_noise_budget"},
		},
		EvidenceNoiseBudget: domain.EvidenceNoiseBudget{
			Enabled:                      true,
			MaxNoiseRatioPercent:         20,
			MaxLowSignalRecordsPerWindow: 10,
			DedupeKeyFields:              []string{"source", "kind", "evidence_ref"},
			ConfigEvidenceField:          "evidence_noise_budget",
			BreachBlocksArchive:          true,
		},
	}
}

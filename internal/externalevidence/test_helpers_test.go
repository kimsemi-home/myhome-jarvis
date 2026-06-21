package externalevidence

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func writePolicy(t *testing.T, root string, policy Policy) {
	t.Helper()
	body, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, filepath.FromSlash(PolicyRelativePath))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, append(body, '\n'), 0o600); err != nil {
		t.Fatal(err)
	}
}

func fixturePolicy(sourceURL string) Policy {
	return Policy{
		SchemaVersion:                    "external_evidence/v1",
		PublicSafe:                       true,
		ExternalNetworkCollectionAllowed: true,
		PrivateRoot:                      "data/private/external-evidence",
		ManifestPath:                     "data/private/external-evidence/manifest.jsonl",
		RawLayerPath:                     "data/private/external-evidence/raw",
		BronzeLayerPath:                  "data/private/external-evidence/bronze.jsonl",
		SilverLayerPath:                  "data/private/external-evidence/silver.jsonl",
		GoldLayerPath:                    "data/private/external-evidence/gold.jsonl",
		ArchiveSourceKey:                 "external_evidence",
		StorageArchiveSourcePath:         "data/private/external-evidence/manifest.jsonl",
		CollectionMaxBytes:               4096,
		SourceClasses:                    []string{"github"},
		SourceDescriptors: []SourceDescriptor{{
			Key: "github_fixture", Class: "github", Method: "GET", URL: sourceURL,
			FreshnessHours: 24, Preprocess: "json_public_metadata",
		}},
		PreprocessingRules: []string{"store_raw_payload_private"},
		RepoSplitAssessment: RepoSplitAssessment{
			Recommendation:      "keep_contract_in_myhome_jarvis_defer_repo_creation",
			FutureRepoCandidate: "kimsemi-home/myhome-external-evidence-lake",
			CreationGate:        "authority_review_required",
			PublicRepoRules: []string{
				"no_raw_payloads", "no_credentials", "no_cookies",
				"no_local_absolute_paths", "private_data_stays_private",
			},
		},
		Commands: []string{
			"mhj external-evidence status",
			"mhj external-evidence repo-split-decision",
			"mhj external-evidence repo-bootstrap",
			"mhj external-evidence collect",
		},
	}
}

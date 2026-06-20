package storagearchive

import (
	"fmt"
	"strings"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

func ValidatePolicy(policy domain.StoragePolicy) error {
	if len(policy.PrivateLogSources) == 0 {
		return fmt.Errorf("storage archive requires private log sources")
	}
	for _, source := range policy.PrivateLogSources {
		if !privateJSONL(source.Path) || source.Format != "jsonl" {
			return fmt.Errorf("storage archive source %q must be private jsonl", source.Key)
		}
	}
	archive := policy.LogArchive
	if archive.Mode != "compress_then_archive" ||
		archive.Compression != "gzip" ||
		archive.ArchiveExtension != ".jsonl.gz" {
		return fmt.Errorf("storage archive compression policy is invalid")
	}
	if !privatePath(archive.ArchiveRoot) || !privateJSONL(archive.ManifestPath) {
		return fmt.Errorf("storage archive paths must stay under data/private")
	}
	if archive.RawPayloadPublicAllowed || !archive.ConfigIsEvidence {
		return fmt.Errorf("storage archive redaction policy is invalid")
	}
	if !hasHashInputs(archive.ConfigHashInputs) {
		return fmt.Errorf("storage archive config hash inputs are invalid")
	}
	noise := policy.EvidenceNoiseBudget
	if !noise.Enabled ||
		!noise.BreachBlocksArchive ||
		noise.MaxNoiseRatioPercent > 25 ||
		noise.MaxLowSignalRecordsPerWindow <= 0 ||
		noise.ConfigEvidenceField == "" {
		return fmt.Errorf("storage archive noise budget is invalid")
	}
	return nil
}

func hasHashInputs(inputs []string) bool {
	seen := map[string]bool{}
	for _, input := range inputs {
		seen[input] = true
	}
	return seen["log_archive"] && seen["evidence_noise_budget"]
}

func privatePath(path string) bool {
	return strings.HasPrefix(path, "data/private/")
}

func privateJSONL(path string) bool {
	return privatePath(path) && strings.HasSuffix(path, ".jsonl")
}

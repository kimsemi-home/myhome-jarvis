package storagearchive

import (
	"path/filepath"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

const PolicyRelativePath = "generated/storage.generated.json"

func StatusForRoot(root string) (Status, error) {
	policy, err := domain.ReadStoragePolicy(
		filepath.Join(root, filepath.FromSlash(PolicyRelativePath)),
	)
	if err != nil {
		return Status{}, err
	}
	if err := ValidatePolicy(policy); err != nil {
		return Status{}, err
	}
	return statusFromPolicy(policy), nil
}

func statusFromPolicy(policy domain.StoragePolicy) Status {
	noise := policy.EvidenceNoiseBudget
	archive := policy.LogArchive
	publicSafe := !archive.RawPayloadPublicAllowed &&
		archive.ConfigIsEvidence && noise.Enabled &&
		noise.BreachBlocksArchive
	return Status{
		PolicyPath:                   PolicyRelativePath,
		Compression:                  archive.Compression,
		ArchiveRoot:                  archive.ArchiveRoot,
		ArchiveExtension:             archive.ArchiveExtension,
		ManifestPath:                 archive.ManifestPath,
		PrivateLogSourceCount:        len(policy.PrivateLogSources),
		PrivateLogSourceKeys:         privateLogSourceKeys(policy),
		Lifecycle:                    append([]string{}, archive.Lifecycle...),
		NoiseBudgetEnabled:           noise.Enabled,
		MaxNoiseRatioPercent:         noise.MaxNoiseRatioPercent,
		MaxLowSignalRecordsPerWindow: noise.MaxLowSignalRecordsPerWindow,
		NoiseBudgetWindow:            noise.Window,
		DedupeKeyFields:              append([]string{}, noise.DedupeKeyFields...),
		ConfigEvidenceField:          noise.ConfigEvidenceField,
		BreachBlocksArchive:          noise.BreachBlocksArchive,
		ConfigIsEvidence:             archive.ConfigIsEvidence,
		RawPayloadPublicAllowed:      archive.RawPayloadPublicAllowed,
		PublicSafe:                   publicSafe,
		ArchiveReady:                 publicSafe && len(policy.PrivateLogSources) > 0,
		NoiseBudgetReady:             noise.Enabled && noise.MaxNoiseRatioPercent <= 25,
		CompressionArchivePattern:    archive.Mode,
		CheckedAt:                    time.Now().UTC().Format(time.RFC3339),
	}
}

func privateLogSourceKeys(policy domain.StoragePolicy) []string {
	keys := make([]string, 0, len(policy.PrivateLogSources))
	for _, source := range policy.PrivateLogSources {
		if source.Key != "" {
			keys = append(keys, source.Key)
		}
	}
	return keys
}

package storagearchive

import (
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

func statusFromPolicy(
	policy domain.StoragePolicy,
	manifest manifestSummary,
) Status {
	noise := policy.EvidenceNoiseBudget
	archive := policy.LogArchive
	evidence := configEvidenceRefForPolicy(policy)
	publicSafe := !archive.RawPayloadPublicAllowed &&
		archive.ConfigIsEvidence && noise.Enabled &&
		noise.BreachBlocksArchive
	status := Status{
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
		ConfigHashInputs:             append([]string{}, evidence.Inputs...),
		ConfigEvidenceSHA256:         evidence.SHA256,
		BreachBlocksArchive:          noise.BreachBlocksArchive,
		ConfigIsEvidence:             archive.ConfigIsEvidence,
		RawPayloadPublicAllowed:      archive.RawPayloadPublicAllowed,
		PublicSafe:                   publicSafe,
		ArchiveReady:                 publicSafe && len(policy.PrivateLogSources) > 0,
		NoiseBudgetReady:             noise.Enabled && noise.MaxNoiseRatioPercent <= 25,
		CompressionArchivePattern:    archive.Mode,
		CheckedAt:                    time.Now().UTC().Format(time.RFC3339),
	}
	return withManifestSummary(status, manifest)
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

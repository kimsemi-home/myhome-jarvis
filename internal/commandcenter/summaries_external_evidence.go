package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/externalevidence"

func summarizeExternalEvidence(status externalevidence.Status) ExternalEvidenceSummary {
	return ExternalEvidenceSummary{
		PublicSafe:                  status.PublicSafe,
		EvidenceReady:               status.ManifestPresent && status.ManifestRecordCount > 0,
		ExternalCollectionAllowed:   status.ExternalCollectionAllowed,
		ForbiddenCollectionDisabled: !status.CredentialsAllowed && !status.CookiesAllowed,
		RawPayloadPublicAllowed:     status.RawPayloadPublicAllowed,
		SourceCount:                 status.SourceCount,
		SourceClassCount:            len(status.SourceClasses),
		ManifestPresent:             status.ManifestPresent,
		ManifestRecordCount:         status.ManifestRecordCount,
		ArchiveSourceKey:            status.ArchiveSourceKey,
		RepoSplitRecommendation:     status.RepoSplitRecommendation,
		RepoCreationGate:            status.RepoCreationGate,
		SplitTriggerCount:           status.SplitTriggerCount,
	}
}

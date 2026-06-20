package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/contextpack"

func summarizeContextPack(status contextpack.Status) ContextPackSummary {
	return ContextPackSummary{
		PackID:                        status.PackID,
		Version:                       status.Version,
		UpstreamCompatibilityVersion:  status.UpstreamCompatibilityVersion,
		OntologyVersion:               status.OntologyVersion,
		PublicSafe:                    status.PublicSafe,
		SplitCriteriaCount:            status.SplitCriteriaCount,
		ExportedArtifactCount:         status.ExportedArtifactCount,
		AuthorityContractVersion:      status.AuthorityContractVersion,
		SecurityContractVersion:       status.SecurityContractVersion,
		VerificationProfile:           status.VerificationProfile,
		VerificationRequiredUnitCount: status.VerificationRequiredUnitCount,
	}
}

package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/repofactory"

func summarizeContextPackEvidence(
	evidence repofactory.ContextPackEvidence,
) ContextPackEvidenceSummary {
	return ContextPackEvidenceSummary{
		DeclarationPath:              evidence.DeclarationPath,
		Valid:                        evidence.Valid,
		EvidenceState:                evidence.EvidenceState,
		DriftCount:                   evidence.DriftCount,
		MissingFieldCount:            evidence.MissingFieldCount,
		MissingArtifactCount:         evidence.MissingArtifactCount,
		StaleVersionCount:            evidence.StaleVersionCount,
		ForbiddenValueCount:          evidence.ForbiddenValueCount,
		PackID:                       evidence.PackID,
		ContextPackVersion:           evidence.ContextPackVersion,
		UpstreamCompatibilityVersion: evidence.UpstreamCompatibilityVersion,
		OntologyVersion:              evidence.OntologyVersion,
		AuthorityContractVersion:     evidence.AuthorityContractVersion,
		SecurityContractVersion:      evidence.SecurityContractVersion,
		VerificationProfile:          evidence.VerificationProfile,
		ExportedArtifactCount:        evidence.ExportedArtifactCount,
		RawDetailsPublicAllowed:      evidence.RawDetailsPublicAllowed,
	}
}

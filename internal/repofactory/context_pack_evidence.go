package repofactory

import "github.com/kimsemi-home/myhome-jarvis/internal/contextpack"

func contextPackEvidenceUnknown() ContextPackEvidence {
	return ContextPackEvidence{
		EvidenceState:           "required_before_creation",
		RawDetailsPublicAllowed: false,
	}
}

func contextPackEvidenceFromStatus(
	status contextpack.Status,
	verify contextpack.VerifyResult,
) ContextPackEvidence {
	state := "declaration_drift"
	if verify.Valid && status.PublicSafe {
		state = "ready"
	}
	return ContextPackEvidence{
		DeclarationPath:              status.DeclarationPath,
		Valid:                        verify.Valid,
		EvidenceState:                state,
		DriftCount:                   verify.DriftCount,
		MissingFieldCount:            verify.MissingFieldCount,
		MissingArtifactCount:         verify.MissingArtifactCount,
		StaleVersionCount:            verify.StaleVersionCount,
		ForbiddenValueCount:          verify.ForbiddenValueCount,
		PackID:                       status.PackID,
		ContextPackVersion:           status.Version,
		UpstreamCompatibilityVersion: status.UpstreamCompatibilityVersion,
		OntologyVersion:              status.OntologyVersion,
		AuthorityContractVersion:     status.AuthorityContractVersion,
		SecurityContractVersion:      status.SecurityContractVersion,
		VerificationProfile:          status.VerificationProfile,
		ExportedArtifactCount:        status.ExportedArtifactCount,
		RawDetailsPublicAllowed:      false,
	}
}

package externalbootstrap

import "github.com/kimsemi-home/myhome-jarvis/internal/repofactory"

func contextHandoff(factory repofactory.DecisionPacket) ContextHandoff {
	evidence := factory.ContextPackEvidence
	return ContextHandoff{
		PackID:                       evidence.PackID,
		ContextPackVersion:           evidence.ContextPackVersion,
		UpstreamCompatibilityVersion: evidence.UpstreamCompatibilityVersion,
		OntologyVersion:              evidence.OntologyVersion,
		AuthorityContractVersion:     evidence.AuthorityContractVersion,
		SecurityContractVersion:      evidence.SecurityContractVersion,
		VerificationProfile:          evidence.VerificationProfile,
		ExportedArtifactCount:        evidence.ExportedArtifactCount,
	}
}

func skeletonFiles(factory repofactory.DecisionPacket) []SkeletonFile {
	files := []SkeletonFile{}
	for _, evidence := range factory.TemplateEvidence {
		files = append(files, SkeletonFile{
			Role:           evidence.Role,
			Path:           evidence.PublicPath,
			SourceArtifact: evidence.SourceArtifact,
			Purpose:        "public_safe_repo_bootstrap",
			State:          evidence.State,
		})
	}
	return files
}

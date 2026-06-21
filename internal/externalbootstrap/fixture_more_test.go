package externalbootstrap

import "github.com/kimsemi-home/myhome-jarvis/internal/repofactory"

func contextPackEvidenceFixture() repofactory.ContextPackEvidence {
	return repofactory.ContextPackEvidence{
		Valid:                        true,
		EvidenceState:                "ready",
		PackID:                       "myhome-jarvis/context-pack",
		ContextPackVersion:           "v1",
		UpstreamCompatibilityVersion: "myhome-jarvis/context-pack/v1",
		OntologyVersion:              "concept-registry/v1",
		AuthorityContractVersion:     "authority/v1",
		SecurityContractVersion:      "security/v1",
		VerificationProfile:          "quality",
		ExportedArtifactCount:        6,
	}
}

func templateEvidenceFixture() []repofactory.TemplateEvidence {
	return []repofactory.TemplateEvidence{
		{
			Role: "generated_ci", PublicPath: ".github/workflows/quality.yml",
			SourceArtifact: "generated/github_quality_workflow.generated.yml",
			State:          "ready",
		},
		{
			Role: "context_pack_declaration", PublicPath: ".mhj/context-pack.json",
			SourceArtifact: "generated/context_pack.generated.json",
			State:          "ready",
		},
	}
}

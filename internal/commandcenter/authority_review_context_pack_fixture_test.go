package commandcenter

func authorityReviewContextPackFixture() ContextPackSummary {
	return ContextPackSummary{
		PackID:                        "myhome-jarvis/context-pack",
		Version:                       "v1",
		UpstreamCompatibilityVersion:  "myhome-jarvis/context-pack/v1",
		OntologyVersion:               "concept-registry/v1",
		PublicSafe:                    true,
		SplitCriteriaCount:            5,
		ExportedArtifactCount:         6,
		AuthorityContractVersion:      "authority/v1",
		SecurityContractVersion:       "security/v1",
		VerificationProfile:           "quality",
		VerificationRequiredUnitCount: 5,
	}
}

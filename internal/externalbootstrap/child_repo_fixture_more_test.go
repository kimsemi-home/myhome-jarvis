package externalbootstrap

func childContextFixture(packet Packet) childContextPack {
	handoff := packet.ContextHandoff
	return childContextPack{
		PackID:                       handoff.PackID,
		UpstreamRepo:                 "kimsemi-home/myhome-jarvis",
		Context:                      packet.Context,
		ContextPackVersion:           handoff.ContextPackVersion,
		UpstreamCompatibilityVersion: handoff.UpstreamCompatibilityVersion,
		OntologyVersion:              handoff.OntologyVersion,
		AuthorityContractVersion:     handoff.AuthorityContractVersion,
		SecurityContractVersion:      handoff.SecurityContractVersion,
		VerificationProfile:          handoff.VerificationProfile,
		ExportedArtifactCount:        handoff.ExportedArtifactCount,
		PrivateLakeStaysPrivate:      true,
		RawPayloadPublicAllowed:      false,
		ExternalWritesAllowed:        false,
		RepoCreationScope:            "repo_creation",
		CandidateRepo:                packet.CandidateRepo,
	}
}

func childHashFixture(packet Packet) childHashCacheDocument {
	return childHashCacheDocument{
		Context:                   packet.Context,
		Version:                   packet.Version,
		CandidateRepo:             packet.CandidateRepo,
		SourcePolicy:              "generated/external_evidence.generated.json",
		GeneratedContractVerified: true,
		HashCacheInputs:           append([]HashCacheInput{}, packet.HashCacheInputs...),
	}
}

package commandcenter

func authorityReviewContractEvidenceChecks(
	brief AuthorityReviewBrief,
	storage StorageArchiveSummary,
) []AuthorityReviewContractEvidenceCheck {
	return []AuthorityReviewContractEvidenceCheck{
		contractCheck("authority_review_request", brief.RequestState, true, brief.PublicSafe),
		contractCheck("public_safety_posture", "ready", true, brief.PublicSafe),
		contractCheck("repo_factory_preflight", brief.RepoFactoryPreflight.CreationDecision, true,
			brief.RepoFactoryPreflight.PublicSafe),
		contractCheck("storage_evidence", storage.CompressionArchivePattern, true,
			storage.PublicSafe && storage.ConfigIsEvidence),
		contractCheck("local_runtime", brief.LocalRuntime.State, true,
			brief.LocalRuntime.PublicSafe && !brief.LocalRuntime.RawRuntimePublicAllowed),
		contractCheck("merge_evidence", publicState(brief.MergeEvidence.PublicSafe), true,
			brief.MergeEvidence.PublicSafe),
		contractCheck("codex_sustainability", brief.CodexSustainability.SustainabilityPosture, true,
			brief.CodexSustainability.PublicSafe),
		contractCheck("external_evidence", externalEvidenceState(brief.ExternalEvidence), true,
			brief.ExternalEvidence.PublicSafe && !brief.ExternalEvidence.RawPayloadPublicAllowed),
		contractCheck("context_pack", brief.ContextPack.Version, true, brief.ContextPack.PublicSafe),
		contractCheck("capability_readiness", publicState(brief.CapabilityReadiness.PublicSafe), true,
			brief.CapabilityReadiness.PublicSafe),
	}
}

func externalEvidenceState(summary ExternalEvidenceSummary) string {
	if summary.EvidenceReady {
		return "ready"
	}
	return "missing_private_manifest"
}

func contractCheck(
	key string,
	state string,
	required bool,
	publicSafe bool,
) AuthorityReviewContractEvidenceCheck {
	if state == "" {
		state = "present"
	}
	return AuthorityReviewContractEvidenceCheck{
		Key:        key,
		State:      state,
		Required:   required,
		PublicSafe: publicSafe,
	}
}

func publicState(ok bool) string {
	if ok {
		return "public_safe"
	}
	return "not_public_safe"
}

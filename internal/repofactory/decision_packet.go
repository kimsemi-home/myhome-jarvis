package repofactory

func DecisionPacketForRoot(root string) (DecisionPacket, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return DecisionPacket{}, err
	}
	return decisionPacketFromPolicy(policy), nil
}

func decisionPacketFromPolicy(policy Policy) DecisionPacket {
	status := statusFromPolicy(policy)
	templates := packetTemplates(policy.TemplateFiles, policy.ForbiddenPublicFragments)
	gates := packetGates(status, policy.CreationGates)
	blockers := packetBlockingGates(gates)
	missing := packetMissingEvidence(gates)
	return DecisionPacket{
		Context:                        "RepoFactoryPreflightDecisionPacket",
		Version:                        "v1",
		PolicyPath:                     PolicyRelativePath,
		PublicSafe:                     status.PublicSafe,
		CreationDecision:               packetDecision(status, blockers),
		CreationAllowed:                false,
		RepoCreationBlockedUntilReview: true,
		SelfApprovalAllowed:            false,
		HumanReviewRequired:            policy.AuthorityReviewRequired,
		PublicSafetyEvidenceRequired:   policy.PublicSafetyEvidenceRequired,
		CodexProjectRequired:           policy.CodexProjectRequired,
		TemplateReadyCount:             packetTemplateReadyCount(templates),
		TemplateFileCount:              status.TemplateFileCount,
		GateReadyCount:                 packetGateReadyCount(gates),
		CreationGateCount:              status.CreationGateCount,
		BlockingGateCount:              blockers,
		MissingEvidenceKeys:            missing,
		NextSafeAction:                 packetNextAction(status, missing),
		TemplateEvidence:               templates,
		CreationGateEvidence:           gates,
		CheckedAt:                      status.CheckedAt,
	}
}

func packetDecision(status Status, blockers int) string {
	if !status.PublicSafe {
		return "repair_repo_factory_policy"
	}
	if blockers > 0 {
		return "blocked_pending_review_evidence"
	}
	return "ready_for_human_review"
}

func packetNextAction(status Status, missing []string) string {
	if !status.PublicSafe {
		return "repair_repo_factory_policy"
	}
	if contains(missing, "authority_review") {
		return "await_human_authority_review"
	}
	if contains(missing, "public_safety_evidence") {
		return "collect_public_safety_evidence"
	}
	return "request_final_repo_creation_review"
}

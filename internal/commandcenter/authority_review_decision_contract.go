package commandcenter

func authorityReviewDecisionContract(
	brief AuthorityReviewBrief,
	status Status,
) AuthorityReviewDecisionContract {
	items := authorityReviewDecisionContractItems(brief.GatedCapabilityKeys)
	checks := authorityReviewContractEvidenceChecks(brief, status.StorageArchive)
	grants := AuthorityReviewContractForbiddenGrants{}
	return AuthorityReviewDecisionContract{
		Context:                      "AuthorityReviewDecisionContract",
		Version:                      "v1",
		PublicSafe:                   authorityReviewDecisionContractPublicSafe(brief, items, checks),
		ReviewerPosture:              "human_only_non_delegable",
		ReviewOnly:                   true,
		CanApplyDecision:             false,
		ReadyCapabilitiesNonBlocking: append([]string{}, brief.CapabilityReadiness.ReadyCapabilityKeys...),
		ContractItems:                items,
		RequiredEvidenceChecks:       checks,
		ForbiddenGrantFlags:          grants,
	}
}

func authorityReviewDecisionContractItems(
	gatedKeys []string,
) []AuthorityReviewDecisionContractItem {
	items := make([]AuthorityReviewDecisionContractItem, 0, len(gatedKeys))
	for _, key := range gatedKeys {
		items = append(items, authorityReviewDecisionContractItem(key))
	}
	return items
}

func authorityReviewDecisionContractItem(
	key string,
) AuthorityReviewDecisionContractItem {
	item := AuthorityReviewDecisionContractItem{
		CapabilityKey:               key,
		DecisionKey:                 key + "_human_review",
		Scope:                       "capability_gate",
		RequiredReviewClass:         "human_review",
		RequiredEvidenceKeys:        commonAuthorityReviewEvidenceKeys(),
		HumanDecisionRecordRequired: true,
	}
	if key == "shorts_factory_control_plane" {
		item.Scope = "public_repo_and_workflow_control"
		item.RequiredReviewClass = "public_repo_change_review"
	}
	if key == "self_improvement_loop" {
		item.Scope = "closed_loop_authority_and_external_write_boundary"
		item.RequiredReviewClass = "workflow_change_review"
	}
	return item
}

func commonAuthorityReviewEvidenceKeys() []string {
	return []string{
		"authority_review_request",
		"public_safety_posture",
		"repo_factory_preflight",
		"storage_evidence",
		"local_runtime",
		"merge_evidence",
		"codex_sustainability",
		"context_pack",
		"capability_readiness",
	}
}

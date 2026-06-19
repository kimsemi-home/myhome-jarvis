package authority

func testPolicy() Policy {
	return Policy{
		Context:                     "AgentCluster",
		Version:                     "v1",
		GeneratedArtifact:           "generated/authority.generated.json",
		PublicStatusRedacted:        true,
		SelfAuthorityAllowed:        false,
		ReasoningTierGrantsApproval: false,
		PublicRepoHighRiskBlocked:   true,
		RequiredInputs:              []string{"confidence_assessor", "evidence_quality", "incident_lifecycle", "control_plane", "translation", "human_review", "public_safety"},
		ReasoningTiers: []ReasoningTier{
			{Key: "r0_compiler", May: []string{"deterministic_transform"}, MustNot: []string{"approve"}},
			{Key: "r1_low", May: []string{"small_candidate"}, MustNot: []string{"approve"}},
			{Key: "r2_medium", May: []string{"multi_file_candidate"}, MustNot: []string{"approve"}},
			{Key: "r3_high", May: []string{"root_cause_candidate"}, MustNot: []string{"approve"}},
			{Key: "r4_governance", May: []string{"policy_review"}, MustNot: []string{"solo_approve"}},
		},
		RolePermissions: []RolePermission{
			{Role: "producer", May: []string{"propose"}, MustNot: []string{"self_approve"}},
			{Role: "independent_reviewer", May: []string{"review_mapping"}, MustNot: []string{"self_approve"}},
			{Role: "adversarial_reviewer", May: []string{"challenge_evidence"}, MustNot: []string{"self_approve"}},
			{Role: "deterministic_verifier", May: []string{"run_checks"}, MustNot: []string{"approve_semantics"}},
			{Role: "governance_steward", May: []string{"gate_authority"}, MustNot: []string{"solo_major_ontology_change"}},
		},
		DomainAttributes:     []string{"agent_reliability", "reasoning_tier", "ontology_maturity", "evidence_quality", "security_impact", "data_sensitivity", "change_risk", "verification_scope", "lease_status", "quarantine_state", "human_review_capacity"},
		Decisions:            testDecisions(),
		Outcomes:             []string{"limited", "review_required", "blocked"},
		AuthorityDebtClasses: []string{"public_safety", "confidence_cap", "evidence_quality", "incident", "control_plane", "translation", "human_review"},
		PublicSummaryFields:  []string{"policy_path", "outcome", "active_rule", "input_count", "decision_count", "allowed_decision_count", "blocked_decision_count", "authority_debt_count", "public_repo_mode", "reasoning_tier_grants_approval", "self_authority_allowed", "public_safety_ok", "confidence_cap", "evidence_quality_debt_count", "incident_debt_count", "control_plane_debt_count", "translation_debt_count", "human_review_debt_count", "human_review_capacity_state", "allowed_decisions", "blocked_decisions", "by_risk", "checked_at"},
		Commands:             []string{"mhj authority status"},
	}
}

func testDecisions() []Decision {
	return []Decision{
		{Key: "read_status", Risk: "low", PublicRepoAllowed: true, AllowedWhenBlocked: true},
		{Key: "evidence_collection", Risk: "low", PublicRepoAllowed: true, AllowedWhenBlocked: true},
		{Key: "deterministic_verification", Risk: "low", PublicRepoAllowed: true, AllowedWhenBlocked: true},
		{Key: "revalidation", Risk: "low", PublicRepoAllowed: true, AllowedWhenBlocked: true},
		{Key: "low_risk_fixture_change", Risk: "medium", PublicRepoAllowed: true},
		{Key: "incident_response", Risk: "medium", PublicRepoAllowed: true, RequiresHumanReview: true, AllowedWhenBlocked: true},
		{Key: "major_ontology_change", Risk: "high", RequiresHumanReview: true},
		{Key: "security_boundary_change", Risk: "high", RequiresHumanReview: true},
		{Key: "production_change", Risk: "high", RequiresHumanReview: true},
		{Key: "evidence_pruning", Risk: "high", RequiresHumanReview: true},
		{Key: "quarantine_release", Risk: "high", RequiresHumanReview: true},
		{Key: "high_risk_automation", Risk: "high", RequiresHumanReview: true},
	}
}

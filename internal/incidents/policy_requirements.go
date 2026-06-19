package incidents

func requiredIncidentKinds() []string {
	return []string{
		"quality_regression",
		"public_safety",
		"evidence_gap",
		"authority_violation",
		"quarantine",
		"feedback_loop_gap",
	}
}

func requiredLifecycleStages() []string {
	return []string{
		"observed",
		"evidence_recorded",
		"classified",
		"owner_assigned",
		"fix_planned",
		"verified",
		"knowledge_updated",
	}
}

func requiredStatuses() []string {
	return []string{"open", "mitigating", "verified", "closed", "quarantined"}
}

func requiredOwnerRoles() []string {
	return []string{
		"producer",
		"independent_reviewer",
		"adversarial_reviewer",
		"deterministic_verifier",
		"governance_steward",
	}
}

func requiredQuarantineStates() []string {
	return []string{"none", "quarantined", "release_requested", "released"}
}

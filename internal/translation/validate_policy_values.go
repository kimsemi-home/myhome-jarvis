package translation

func validatePolicyRequiredValues(policy Policy) error {
	checks := []struct {
		label    string
		values   []string
		required []string
	}{
		{"translation policy context", policy.AllowedContexts, []string{
			"AgentCluster", "KnowledgeIndex", "AgentOps", "SecurityPolicy",
		}},
		{"translation manifest required field", normalizeList(policy.RequiredManifestFields), []string{
			"source_context", "target_context", "source_version", "target_version",
			"preserved_rules", "known_losses", "owner", "evidence_refs",
		}},
		{"translation loss level", normalizeList(policy.LossLevels), []string{
			"l0_none", "l1_note", "l2_degraded", "l3_review_required", "l4_forbidden",
		}},
		{"translation loss category", normalizeList(policy.AllowedLossCategories), []string{
			"mapping_gap", "authority", "security_boundary", "financial_commitment",
		}},
	}
	for _, check := range checks {
		if err := requireAll(check.label, check.values, check.required); err != nil {
			return err
		}
	}
	return nil
}

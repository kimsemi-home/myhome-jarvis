package controlplane

func validatePolicyLists(policy Policy) error {
	checks := []struct {
		label    string
		values   []string
		required []string
	}{
		{"control-plane decision kind", normalizeList(policy.AllowedDecisionKinds), []string{
			"loop_once", "loop_worker_cycle", "checkpoint_write",
		}},
		{"control-plane authority profile", normalizeList(policy.AllowedAuthorityProfiles), []string{
			"local_readonly", "external_write_gated",
		}},
		{"control-plane lease status", normalizeList(policy.AllowedLeaseStatuses), []string{
			"issued", "active", "finished", "aborted", "quarantined",
		}},
		{"control-plane required field", normalizeList(policy.RequiredFields), []string{
			"decision_kind", "policy_version", "ontology_version",
			"authority_profile", "selected_route", "reviewer_role",
			"verifier_role", "lease_seconds", "lease_status",
			"evidence_refs", "output_ref",
		}},
	}
	for _, check := range checks {
		if err := requireAll(check.label, check.values, check.required); err != nil {
			return err
		}
	}
	return nil
}

package controlplane

func testPolicy() Policy {
	return Policy{
		Context:                    "AgentOps",
		Version:                    "v1",
		GeneratedArtifact:          "generated/control_plane.generated.json",
		PrivateManifestLedger:      "data/private/control-plane/manifests.jsonl",
		ManifestRequired:           true,
		AppendOnly:                 true,
		PublicStatusRedacted:       true,
		RawRationalePublicAllowed:  false,
		VerifierSeparationRequired: true,
		MinLeaseSeconds:            1,
		MaxLeaseSeconds:            3600,
		AllowedDecisionKinds:       []string{"loop_once", "loop_worker_cycle", "linear_next_observation", "checkpoint_write"},
		AllowedAuthorityProfiles:   []string{"local_readonly", "external_write_gated", "review_only"},
		AllowedLeaseStatuses:       []string{"issued", "active", "expiring", "expired", "finished", "aborted", "escalated", "quarantined"},
		RequiredFields:             []string{"decision_kind", "policy_version", "ontology_version", "authority_profile", "selected_route", "reviewer_role", "verifier_role", "lease_seconds", "lease_status", "evidence_refs", "output_ref"},
		AllowedEvidencePrefixes:    []string{"data/private/", "generated/", "docs/", "cmd/", "internal/", "apps/flutter/", "lisp/", "crates/", "fixtures/", "harness/", ".github/"},
		PublicSummaryFields:        []string{"policy_path", "manifest_path", "exists", "count", "invalid_manifest_count", "manifest_debt_count", "verifier_separation_required", "verifier_violation_count", "min_lease_seconds", "max_lease_seconds", "by_decision_kind", "by_authority_profile", "by_lease_status", "last_observed_at", "checked_at"},
		Commands:                   []string{"mhj control-plane status", "mhj control-plane verify"},
	}
}

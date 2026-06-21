package externalevidence

func repoSplitDecisionOptions(
	assessment RepoSplitAssessment,
) []RepoSplitDecisionOption {
	return []RepoSplitDecisionOption{
		{
			Key:                      "keep_contract_in_myhome_jarvis",
			Label:                    "Keep public contract in myhome-jarvis",
			Summary:                  "retain redacted status, SSOT, policy validation, and archive contract here",
			PrivacyRisk:              "lowest_public_surface_no_raw_payloads",
			MaintenanceBurden:        "single_repo_contract_and_quality_run",
			SSOTHandoff:              "native_ssot_no_cross_repo_sync",
			ContextPackHandoff:       "uses_existing_context_pack_export",
			GitHubActionsCost:        "uses_existing_quality_cache",
			ArchiveCacheBehavior:     "uses_existing_external_evidence_archive_source",
			OntologyVersionDiscovery: "single_repo_generated_context_pack",
		},
		{
			Key:                      "split_new_external_evidence_lake_repo",
			Label:                    "Create " + assessment.FutureRepoCandidate + " after review",
			Summary:                  "move collectors and lake experiments after explicit authority approval",
			PrivacyRisk:              "higher_public_surface_requires_private_payload_boundary",
			MaintenanceBurden:        "separate_release_ci_context_pack_sync",
			SSOTHandoff:              "requires_context_pack_import_and_version_pin",
			ContextPackHandoff:       "must_declare_context_ontology_authority_security",
			GitHubActionsCost:        "adds_quality_workflow_minutes_and_cache_keys",
			ArchiveCacheBehavior:     "must_keep_private_lake_hash_archives_local",
			OntologyVersionDiscovery: "requires_cross_repo_ontology_version_discovery",
			HumanApprovalRequired:    true,
		},
	}
}

package externalevidence

func repoRuleState(assessment RepoSplitAssessment) string {
	if repoRulesPublicSafe(assessment) {
		return "ready"
	}
	return "missing_public_repo_rules"
}

func repoRulesPublicSafe(assessment RepoSplitAssessment) bool {
	return contains(assessment.PublicRepoRules, "no_raw_payloads") &&
		contains(assessment.PublicRepoRules, "no_credentials") &&
		contains(assessment.PublicRepoRules, "no_cookies") &&
		contains(assessment.PublicRepoRules, "no_local_absolute_paths") &&
		contains(assessment.PublicRepoRules, "private_data_stays_private")
}

func privatePayloadState(policy Policy) string {
	if !policy.RawPayloadPublicAllowed &&
		!policy.CredentialsAllowed &&
		!policy.CookiesAllowed {
		return "private_only"
	}
	return "unsafe_public_payload_boundary"
}

package financeconsent

func missingActiveKinds(activeKinds map[string]bool) []string {
	missing := []string{}
	for _, kind := range requiredConsentKinds {
		if !activeKinds[kind] {
			missing = append(missing, kind)
		}
	}
	return missing
}

func readinessState(status Status) string {
	if status.ForbiddenActionEnabledCount > 0 {
		return "blocked"
	}
	if !readOnlyReviewOnly(status.FinanceMode) {
		return "blocked"
	}
	if !status.Exists || status.MissingRequiredConsentCount > 0 {
		return "blocked"
	}
	if status.InvalidRecordCount > 0 ||
		status.MissingEvidenceCount > 0 ||
		status.ReviewRequiredCount > 0 {
		return "review_required"
	}
	return "ready_read_only"
}

func readOnlyReviewOnly(mode string) bool {
	return mode == "read_only_review_only"
}

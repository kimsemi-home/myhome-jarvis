package externalbootstrap

func finalizeChildRepoStatus(status ChildRepoStatus) ChildRepoStatus {
	status.MissingFileCount = countChildFindings(status, "missing_file")
	status.DriftCount = countChildFindings(status, "drift")
	status.InvalidHashCacheCount = countChildFindings(status, "invalid_hash_cache")
	status.PublicSafetyFindingCount = countChildFindings(status, "public_safety")
	status.ContextPackValid = status.ContextPackValid && status.DriftCount == 0
	status.HashCacheValid = status.HashCacheValid && status.InvalidHashCacheCount == 0
	status.Valid = status.CheckoutState == "present" && status.ContextPackValid &&
		status.HashCacheValid && status.PublicSafetyOK && status.PrivateDataAbsent &&
		status.MissingFileCount == 0 && status.DriftCount == 0 &&
		status.InvalidHashCacheCount == 0 && status.PublicSafetyFindingCount == 0
	status.EvidenceState = childEvidenceState(status)
	status.NextSafeAction = childNextSafeAction(status)
	return status
}

func countChildFindings(status ChildRepoStatus, prefix string) int {
	count := 0
	for _, finding := range status.Findings {
		if len(finding.Code) >= len(prefix) && finding.Code[:len(prefix)] == prefix {
			count++
		}
	}
	return count
}

func childEvidenceState(status ChildRepoStatus) string {
	if status.Valid {
		return "ready"
	}
	if status.CheckoutState != "present" || status.MissingFileCount > 0 {
		return "missing"
	}
	return "drifted"
}

func childNextSafeAction(status ChildRepoStatus) string {
	if status.Valid {
		return "use_child_repo_as_cross_repo_evidence"
	}
	if status.CheckoutState != "present" {
		return "provide_child_repo_checkout_path"
	}
	if status.MissingFileCount > 0 {
		return "restore_minimal_public_skeleton"
	}
	return "reconcile_child_repo_context_and_hash_cache"
}

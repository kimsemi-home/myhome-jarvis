package review

func capacityState(policy Policy, status Status) (string, string) {
	if status.HighRiskOpenCount > policy.MaxHighRiskOpenReviews {
		return "overloaded", "high_risk_review_overload"
	}
	if status.OpenCount > policy.MaxOpenReviews {
		return "overloaded", "review_queue_overload"
	}
	if status.MissingReviewerCount > 0 {
		return "constrained", "missing_reviewer"
	}
	if status.MissingEvidenceCount > 0 {
		return "constrained", "missing_evidence_ref"
	}
	if status.InvalidReviewCount > 0 {
		return "constrained", "invalid_review_record"
	}
	if status.OpenCount > 0 && status.BackupAvailableCount < policy.MinBackupReviewers {
		return "constrained", "backup_reviewer_unavailable"
	}
	if status.OpenCount > 0 {
		return "constrained", "open_review_queue"
	}
	return "available", "no_open_reviews"
}

func finalizeStatus(policy Policy, status Status) Status {
	status.ReviewDebtCount = status.InvalidReviewCount + status.MissingEvidenceCount
	status.ReviewDebtCount += status.MissingReviewerCount + status.HighRiskOpenCount
	if status.OpenCount > policy.MaxOpenReviews {
		status.ReviewDebtCount += status.OpenCount - policy.MaxOpenReviews
	}
	status.CapacityState, status.ActiveRule = capacityState(policy, status)
	return status
}

func reviewOpen(status string) bool {
	switch status {
	case "approved", "rejected":
		return false
	default:
		return true
	}
}

package commandcenter

func workItemRef(status Status) string {
	if status.AuthorityReview.RequestID != "" {
		return "universal_work_item:" + status.AuthorityReview.RequestID
	}
	if status.NextSafeAction != "" {
		return "universal_work_item:" + status.NextSafeAction
	}
	return "universal_work_item:closed_loop"
}

func authorityRef(status Status) string {
	if status.AuthorityReview.RequestID == "" {
		return ""
	}
	return "authority_review_queue:" + status.AuthorityReview.RequestID
}

func workItemState(status Status) string {
	if status.AuthorityReview.QueueReady {
		return "pending_authority_review"
	}
	if status.BlockedGateCount > 0 {
		return "gated"
	}
	return "ready"
}

func decisionKey(action string) string {
	if action == "request_authority_review" {
		return "authority_review_request"
	}
	if action == "" {
		return "closed_loop_review"
	}
	return action
}

func mergeEligibilityHint(status Status) string {
	if status.MergeEvidence.MergeReady {
		return "merge_when_checks_pass"
	}
	return "collect_merge_evidence"
}

package main

func routeAuthorityReviewStatus(root string, args []string) (bool, error) {
	if len(args) != 2 || args[0] != "authority-review" {
		return false, nil
	}
	switch args[1] {
	case "approval-status":
		return true, authorityReviewApprovalStatus(root)
	case "brief":
		return true, authorityReviewBrief(root)
	case "decision-packet":
		return true, authorityReviewDecisionPacket(root)
	case "request":
		return true, authorityReviewRequest(root)
	case "evidence":
		return true, authorityReviewEvidence(root)
	case "queue":
		return true, authorityReviewQueue(root)
	default:
		return false, nil
	}
}

func routeAuthorityReview(root string, args []string) error {
	if len(args) == 2 && args[0] == "record" {
		return authorityReviewRecord(root, []byte(args[1]))
	}
	if len(args) == 2 && args[0] == "approve" {
		return authorityReviewApprove(root, []byte(args[1]))
	}
	if len(args) >= 1 && args[0] == "refresh" {
		return authorityReviewRefresh(root, args[1:])
	}
	if len(args) == 1 {
		switch args[0] {
		case "approval-status":
			return authorityReviewApprovalStatus(root)
		case "brief":
			return authorityReviewBrief(root)
		case "decision-packet":
			return authorityReviewDecisionPacket(root)
		case "status":
			return authorityReviewStatus(root)
		case "request":
			return authorityReviewRequest(root)
		case "evidence":
			return authorityReviewEvidence(root)
		case "queue":
			return authorityReviewQueue(root)
		}
	}
	return usage()
}

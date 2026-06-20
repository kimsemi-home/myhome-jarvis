package main

func routeAuthorityReviewStatus(root string, args []string) (bool, error) {
	if len(args) != 2 || args[0] != "authority-review" {
		return false, nil
	}
	switch args[1] {
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
	if len(args) == 1 {
		switch args[0] {
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

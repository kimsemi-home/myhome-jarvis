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

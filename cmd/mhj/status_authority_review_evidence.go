package main

import "github.com/kimsemi-home/myhome-jarvis/internal/authority"

func authorityReviewEvidence(root string) error {
	status, err := authority.ReviewRequestEvidenceForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

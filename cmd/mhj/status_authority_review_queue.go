package main

import "github.com/kimsemi-home/myhome-jarvis/internal/authority"

func authorityReviewQueue(root string) error {
	status, err := authority.ReviewQueueStatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

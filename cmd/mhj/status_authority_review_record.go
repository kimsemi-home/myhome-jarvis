package main

import "github.com/kimsemi-home/myhome-jarvis/internal/authority"

func authorityReviewRecord(root string, payload []byte) error {
	result, err := authority.RecordReviewRequest(root, payload)
	if err != nil {
		return err
	}
	return writeJSON(result)
}

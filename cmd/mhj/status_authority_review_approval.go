package main

import "github.com/kimsemi-home/myhome-jarvis/internal/authority"

func authorityReviewApprovalStatus(root string) error {
	status, err := authority.ApprovalDecisionStatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func authorityReviewApprove(root string, payload []byte) error {
	result, err := authority.RecordApprovalDecision(root, payload)
	if err != nil {
		return err
	}
	return writeJSON(result)
}

package main

import "github.com/kimsemi-home/myhome-jarvis/internal/commandcenter"

func authorityReviewBrief(root string) error {
	brief, err := commandcenter.AuthorityReviewBriefForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(brief)
}

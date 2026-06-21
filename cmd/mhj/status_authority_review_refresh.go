package main

import "github.com/kimsemi-home/myhome-jarvis/internal/authority"

func authorityReviewRefresh(root string, args []string) error {
	if len(args) > 1 {
		return usage()
	}
	linearRef := ""
	if len(args) == 1 {
		linearRef = args[0]
	}
	result, err := authority.RefreshReviewRequest(root, linearRef)
	if err != nil {
		return err
	}
	return writeJSON(result)
}

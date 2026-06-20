package main

import "github.com/kimsemi-home/myhome-jarvis/internal/mergeevidence"

func mergeEvidenceStatus(root string) error {
	status, err := mergeevidence.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

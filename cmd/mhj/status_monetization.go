package main

import "github.com/kimsemi-home/myhome-jarvis/internal/monetization"

func monetizationStatus(root string) error {
	status, err := monetization.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

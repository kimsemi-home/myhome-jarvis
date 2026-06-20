package main

import "github.com/kimsemi-home/myhome-jarvis/internal/mediareadiness"

func mediaReadinessStatus(root string) error {
	status, err := mediareadiness.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

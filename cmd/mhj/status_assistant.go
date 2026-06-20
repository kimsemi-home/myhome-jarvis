package main

import "github.com/kimsemi-home/myhome-jarvis/internal/commandcenter"

func assistantStatus(root string) error {
	status, err := commandcenter.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

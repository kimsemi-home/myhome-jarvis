package main

import "github.com/kimsemi-home/myhome-jarvis/internal/repofactory"

func repoFactoryStatus(root string) error {
	status, err := repofactory.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

package main

import "github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"

func storageArchiveStatus(root string) error {
	status, err := storagearchive.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

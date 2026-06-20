package main

import "github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"

func storageArchiveRun(root string) error {
	report, err := storagearchive.RunForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(report)
}

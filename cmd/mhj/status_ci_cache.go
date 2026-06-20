package main

import "github.com/kimsemi-home/myhome-jarvis/internal/cicache"

func ciCacheStatus(root string) error {
	status, err := cicache.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

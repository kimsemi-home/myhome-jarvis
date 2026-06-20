package main

import "github.com/kimsemi-home/myhome-jarvis/internal/security"

func securityStatus(root string) error {
	status, err := security.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

package main

import "github.com/kimsemi-home/myhome-jarvis/internal/financeconsent"

func financeConsentStatus(root string) error {
	status, err := financeconsent.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

package main

import "github.com/kimsemi-home/myhome-jarvis/internal/financeconsent"

func routeFinanceConsent(root string, args []string) error {
	if len(args) == 2 && args[0] == "record" {
		return financeConsentRecord(root, []byte(args[1]))
	}
	if len(args) == 1 && args[0] == "status" {
		return financeConsentStatus(root)
	}
	return usage()
}

func financeConsentRecord(root string, payload []byte) error {
	result, err := financeconsent.RecordConsent(root, payload)
	if err != nil {
		return err
	}
	return writeJSON(result)
}

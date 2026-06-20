package main

import "github.com/kimsemi-home/myhome-jarvis/internal/monetization"

func routeMonetization(root string, args []string) error {
	if len(args) == 2 && args[0] == "record" {
		return monetizationRecord(root, []byte(args[1]))
	}
	if len(args) == 1 && args[0] == "status" {
		return monetizationStatus(root)
	}
	return usage()
}

func monetizationRecord(root string, payload []byte) error {
	result, err := monetization.RecordExperiment(root, payload)
	if err != nil {
		return err
	}
	return writeJSON(result)
}

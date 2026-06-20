package main

import "github.com/kimsemi-home/myhome-jarvis/internal/codexsustainability"

func routeCodexSustainability(root string, args []string) error {
	if len(args) == 1 && args[0] == "record-quality" {
		return codexSustainabilityRecordQuality(root)
	}
	if len(args) == 2 && args[0] == "record-proposal" {
		return codexSustainabilityRecordProposal(root, []byte(args[1]))
	}
	return usage()
}

func codexSustainabilityRecordQuality(root string) error {
	status, err := codexsustainability.CaptureQualityRun(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func codexSustainabilityRecordProposal(root string, payload []byte) error {
	result, err := codexsustainability.RecordProposal(root, payload)
	if err != nil {
		return err
	}
	return writeJSON(result)
}

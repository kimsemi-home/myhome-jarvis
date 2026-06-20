package main

import "github.com/kimsemi-home/myhome-jarvis/internal/codexcost"

func routeCodexCost(root string, args []string) error {
	if len(args) == 2 && args[0] == "record" {
		return codexCostRecord(root, []byte(args[1]))
	}
	return usage()
}

func codexCostRecord(root string, payload []byte) error {
	result, err := codexcost.RecordUsage(root, payload)
	if err != nil {
		return err
	}
	return writeJSON(result)
}

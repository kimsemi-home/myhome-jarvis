package main

import "github.com/kimsemi-home/myhome-jarvis/internal/codexcost"

func routeCodexCost(root string, args []string) error {
	if len(args) == 1 && args[0] == "scaling-packet" {
		return codexCostScalingPacket(root)
	}
	if len(args) == 1 && args[0] == "brief" {
		return codexCostBrief(root)
	}
	if len(args) == 1 && args[0] == "roi" {
		return codexCostROI(root)
	}
	if len(args) == 2 && args[0] == "attribute" {
		return codexCostAttribute(root, []byte(args[1]))
	}
	if len(args) == 2 && args[0] == "record" {
		return codexCostRecord(root, []byte(args[1]))
	}
	if len(args) == 2 && args[0] == "guard" {
		return codexCostGuard(root, []byte(args[1]))
	}
	return usage()
}

func codexCostScalingPacket(root string) error {
	packet, err := codexcost.ScalingPacketForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(packet)
}

func codexCostBrief(root string) error {
	brief, err := codexcost.BriefForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(brief)
}

func codexCostROI(root string) error {
	summary, err := codexcost.ROISummaryForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(summary)
}

func codexCostAttribute(root string, payload []byte) error {
	result, err := codexcost.AttributeCost(root, payload)
	if err != nil {
		return err
	}
	return writeJSON(result)
}

func codexCostRecord(root string, payload []byte) error {
	result, err := codexcost.RecordUsage(root, payload)
	if err != nil {
		return err
	}
	return writeJSON(result)
}

func codexCostGuard(root string, payload []byte) error {
	result, err := codexcost.GuardLoop(root, payload)
	if err != nil {
		return err
	}
	return writeJSON(result)
}

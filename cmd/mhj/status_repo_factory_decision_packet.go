package main

import "github.com/kimsemi-home/myhome-jarvis/internal/repofactory"

func repoFactoryDecisionPacket(root string) error {
	packet, err := repofactory.DecisionPacketForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(packet)
}

package main

import "github.com/kimsemi-home/myhome-jarvis/internal/commandcenter"

func authorityReviewDecisionPacket(root string) error {
	packet, err := commandcenter.AuthorityReviewDecisionPacketForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(packet)
}

package main

import "github.com/kimsemi-home/myhome-jarvis/internal/authority"

func authorityReviewRequest(root string) error {
	packet, err := authority.ReviewRequestPacketForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(packet)
}

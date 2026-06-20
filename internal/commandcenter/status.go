package commandcenter

import "time"

func StatusForRoot(root string) (Status, error) {
	in, err := collectInputs(root)
	if err != nil {
		return Status{}, err
	}
	status := Status{
		Context:    "AssistantCommandCenter",
		Version:    "v1",
		PublicSafe: publicSafe(in),
		Redaction:  "public-summary-only",
		Vision:     summarizeVision(in.Vision),
		PDCA:       summarizePDCA(in.PDCA),
		Evidence:   summarizeEvidence(in.Evidence),
		Incidents:  summarizeIncidents(in.Incidents),
		Authority:  summarizeAuthority(in.Authority),
		Review:     summarizeReview(in.Review),
		Cost:       summarizeCost(in.Cost),
		CheckedAt:  time.Now().UTC().Format(time.RFC3339),
	}
	status.BlockedGates = blockedGates(in)
	status.BlockedGateCount = len(status.BlockedGates)
	status.CompactState = compactState(status)
	status.NextSafeAction = nextSafeAction(status)
	return status, nil
}

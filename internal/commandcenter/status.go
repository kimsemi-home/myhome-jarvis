package commandcenter

import "time"

func StatusForRoot(root string) (Status, error) {
	in, err := collectInputs(root)
	if err != nil {
		return Status{}, err
	}
	status := Status{
		Context:             "AssistantCommandCenter",
		Version:             "v1",
		PublicSafe:          publicSafe(in),
		Redaction:           "public-summary-only",
		Vision:              summarizeVision(in.Vision),
		PDCA:                summarizePDCA(in.PDCA),
		Evidence:            summarizeEvidence(in.Evidence),
		EvidenceIntegrity:   summarizeEvidenceIntegrity(in.EvidenceIntegrity),
		Incidents:           summarizeIncidents(in.Incidents),
		Authority:           summarizeAuthority(in.Authority),
		AuthorityReview:     summarizeAuthorityReview(in.AuthorityReview),
		Review:              summarizeReview(in.Review),
		FinanceConsent:      summarizeFinanceConsent(in.FinanceConsent),
		Cost:                summarizeCost(in.Cost),
		CodexSustainability: summarizeCodexSustainability(in.CodexSustainability),
		ContextPack:         summarizeContextPack(in.ContextPack),
		MediaReadiness:      summarizeMediaReadiness(in.MediaReadiness),
		MergeEvidence:       summarizeMergeEvidence(in.MergeEvidence),
		Monetization:        summarizeMonetization(in.Monetization),
		RepoFactory:         summarizeRepoFactory(in.RepoFactory),
		CheckedAt:           time.Now().UTC().Format(time.RFC3339),
	}
	status.BlockedGates = blockedGates(in)
	status.BlockedGateCount = len(status.BlockedGates)
	status.CompactState = compactState(status)
	status.NextSafeAction = nextSafeAction(status)
	return status, nil
}

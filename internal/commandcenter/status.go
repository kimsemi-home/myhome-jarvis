package commandcenter

import "time"

func StatusForRoot(root string) (Status, error) {
	in, err := collectInputs(root)
	if err != nil {
		return Status{}, err
	}
	status := Status{
		Context:              "AssistantCommandCenter",
		Version:              "v1",
		PublicSafe:           publicSafe(in),
		Redaction:            "public-summary-only",
		Vision:               summarizeVision(in.Vision),
		PDCA:                 summarizePDCA(in.PDCA),
		Evidence:             summarizeEvidence(in.Evidence),
		EvidenceIntegrity:    summarizeEvidenceIntegrity(in.EvidenceIntegrity),
		Incidents:            summarizeIncidents(in.Incidents),
		Authority:            summarizeAuthority(in.Authority),
		AuthorityReview:      summarizeAuthorityReview(in.AuthorityReview),
		Review:               summarizeReview(in.Review),
		FinanceConsent:       summarizeFinanceConsent(in.FinanceConsent),
		Cost:                 summarizeCost(in.Cost),
		CodexCostBrief:       summarizeCodexCostBrief(in.CostBrief),
		CodexCostScaling:     summarizeCodexCostScaling(in.CostScaling),
		CodexSustainability:  summarizeCodexSustainability(in.CodexSustainability),
		ExternalEvidence:     summarizeExternalEvidence(in.ExternalEvidence),
		StorageArchive:       summarizeStorageArchive(in.StorageArchive),
		Supervisor:           summarizeSupervisor(in.Supervisor),
		ContextPack:          summarizeContextPack(in.ContextPack),
		MediaReadiness:       summarizeMediaReadiness(in.MediaReadiness),
		MergeEvidence:        summarizeMergeEvidence(in.MergeEvidence),
		Monetization:         summarizeMonetization(in.Monetization),
		RepoFactory:          summarizeRepoFactory(in.RepoFactory),
		RepoFactoryPreflight: summarizeRepoFactoryPreflight(in.RepoFactoryPreflight),
		CheckedAt:            time.Now().UTC().Format(time.RFC3339),
	}
	status.LocalRuntime = summarizeLocalRuntime(status.Supervisor)
	status.BlockedGates = blockedGates(in)
	status.BlockedGateCount = len(status.BlockedGates)
	status.CompactState = compactState(status)
	status.NextSafeAction = nextSafeAction(status)
	applyVisionReadiness(&status)
	status.WorkItem = summarizeWorkItem(status)
	return status, nil
}

package commandcenter

import "github.com/kimsemi-home/myhome-jarvis/internal/repofactory"

func summarizeRepoFactoryPreflight(
	packet repofactory.DecisionPacket,
) RepoFactoryPreflightSummary {
	return RepoFactoryPreflightSummary{
		PublicSafe:                     packet.PublicSafe,
		CreationDecision:               packet.CreationDecision,
		CreationAllowed:                packet.CreationAllowed,
		RepoCreationBlockedUntilReview: packet.RepoCreationBlockedUntilReview,
		SelfApprovalAllowed:            packet.SelfApprovalAllowed,
		TemplateReadyCount:             packet.TemplateReadyCount,
		TemplateFileCount:              packet.TemplateFileCount,
		GateReadyCount:                 packet.GateReadyCount,
		CreationGateCount:              packet.CreationGateCount,
		BlockingGateCount:              packet.BlockingGateCount,
		MissingEvidenceKeys:            append([]string{}, packet.MissingEvidenceKeys...),
		NextSafeAction:                 packet.NextSafeAction,
	}
}

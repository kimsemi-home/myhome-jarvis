package commandcenter

import "time"

const visionCompletionRule = "all capabilities ready, evidence retention ready, and no open gates"

func VisionAuditForRoot(root string) (VisionAudit, error) {
	status, err := StatusForRoot(root)
	if err != nil {
		return VisionAudit{}, err
	}
	policy, err := readVisionPolicy(root)
	if err != nil {
		return VisionAudit{}, err
	}
	return visionAuditFromStatus(policy, status), nil
}

func visionAuditFromStatus(
	policy visionPolicy,
	status Status,
) VisionAudit {
	return VisionAudit{
		Context:                 "AssistantVisionCompletionAudit",
		Version:                 "v1",
		PublicSafe:              status.PublicSafe,
		Redaction:               "public-summary-only",
		PolicyPath:              visionPolicyPath,
		Mission:                 policy.Mission,
		OperatingMode:           policy.OperatingMode,
		UniversalTermCount:      len(policy.UniversalTerms),
		LinearEpicCount:         len(policy.LinearEpics),
		RequirementCount:        status.Vision.CapabilityCount,
		ReadyRequirementCount:   status.Vision.ReadyPillarCount,
		GatedRequirementCount:   status.Vision.GatedPillarCount,
		BlockedRequirementCount: status.Vision.BlockedPillarCount,
		OpenGateCount:           status.BlockedGateCount,
		GoalComplete:            visionGoalComplete(status),
		CompletionRule:          visionCompletionRule,
		NextSafeAction:          status.NextSafeAction,
		EvidenceRetention:       visionEvidenceRetention(status.StorageArchive),
		Requirements:            visionAuditRequirements(status),
		CheckedAt:               time.Now().UTC().Format(time.RFC3339),
	}
}

func visionGoalComplete(status Status) bool {
	return status.PublicSafe &&
		status.Vision.CapabilityCount > 0 &&
		status.Vision.ReadyPillarCount == status.Vision.CapabilityCount &&
		visionEvidenceRetentionReady(status.StorageArchive) &&
		status.BlockedGateCount == 0 &&
		status.NextSafeAction == "continue_closed_loop_planning"
}

func visionEvidenceRetentionReady(summary StorageArchiveSummary) bool {
	return summary.PublicSafe &&
		summary.ArchiveReady &&
		summary.NoiseBudgetReady &&
		summary.CompressionArchivePattern == "compress_then_archive" &&
		summary.ConfigIsEvidence &&
		summary.BreachBlocksArchive &&
		summary.ManifestBudgetBreachCount == 0 &&
		summary.ManifestInvalidEntryCount == 0
}

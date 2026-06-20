package commandcenter

func visionAuditRequirements(status Status) []VisionRequirementAudit {
	rows := make([]VisionRequirementAudit, 0, len(status.Vision.PillarKeys))
	for _, key := range status.Vision.PillarKeys {
		state := visionPillarReadiness(key, status)
		rows = append(rows, VisionRequirementAudit{
			CapabilityKey:  key,
			State:          state,
			EvidenceRefs:   visionEvidenceRefs(key),
			GateRefs:       visionGateRefs(key, state, status),
			NextSafeAction: visionRequirementNextAction(key, state, status),
		})
	}
	return rows
}

func visionRequirementNextAction(
	key string,
	state string,
	status Status,
) string {
	if state == "ready" {
		return "none"
	}
	switch key {
	case "local_media_concierge":
		return "repair_media_readiness"
	case "household_finance_copilot":
		return "record_finance_consent"
	case "monetization_console":
		return "review_monetization_experiments"
	case "codex_cost_governor":
		return "review_codex_cost_budget"
	default:
		return status.NextSafeAction
	}
}

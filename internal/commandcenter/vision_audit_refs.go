package commandcenter

func visionEvidenceRefs(key string) []string {
	switch key {
	case "local_media_concierge":
		return []string{"media_readiness"}
	case "household_finance_copilot":
		return []string{"finance_consent"}
	case "shorts_factory_control_plane":
		return []string{"repo_factory", "authority_review", "merge_evidence"}
	case "monetization_console":
		return []string{"monetization", "codex_cost"}
	case "codex_cost_governor":
		return []string{"codex_cost", "codex_sustainability", "storage_archive"}
	case "self_improvement_loop":
		return []string{
			"pdca", "evidence", "storage_archive", "incidents", "review",
			"authority",
		}
	default:
		return []string{"assistant_status"}
	}
}

func visionGateRefs(key string, state string, status Status) []string {
	if state == "ready" {
		return []string{}
	}
	allowed := visionGateCandidates(key)
	refs := []string{}
	for _, gate := range status.BlockedGates {
		if allowed[gate.Key] {
			refs = append(refs, gate.Key)
		}
	}
	return refs
}

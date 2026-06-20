package commandcenter

func visionGateCandidates(key string) map[string]bool {
	switch key {
	case "local_media_concierge":
		return gateSet("media_readiness")
	case "household_finance_copilot":
		return gateSet("finance_consent")
	case "shorts_factory_control_plane":
		return gateSet("repo_factory", "authority", "merge_evidence")
	case "monetization_console":
		return gateSet("monetization", "cost", "authority")
	case "codex_cost_governor":
		return gateSet("cost", "codex_sustainability", "storage_archive")
	case "self_improvement_loop":
		return gateSet(
			"pdca", "evidence", "incidents", "review", "authority",
			"storage_archive", "local_runtime",
		)
	default:
		return map[string]bool{}
	}
}

func gateSet(keys ...string) map[string]bool {
	set := map[string]bool{}
	for _, key := range keys {
		set[key] = true
	}
	return set
}

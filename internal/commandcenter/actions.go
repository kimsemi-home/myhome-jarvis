package commandcenter

func nextSafeAction(status Status) string {
	if !status.PublicSafe {
		return "run_public_safety_review"
	}
	for _, gate := range status.BlockedGates {
		switch gate.Key {
		case "authority":
			return "resolve_authority_gate"
		case "review":
			return "assign_or_reduce_human_review"
		case "cost":
			return "review_codex_cost_budget"
		case "incidents":
			return "close_or_quarantine_incidents"
		case "evidence":
			return "record_missing_evidence"
		case "pdca":
			return "complete_pdca_artifacts"
		}
	}
	return "continue_closed_loop_planning"
}

func compactState(status Status) string {
	if !status.PublicSafe {
		return "unsafe"
	}
	if status.Authority.Outcome == "blocked" {
		return "blocked"
	}
	if status.BlockedGateCount > 0 {
		return "gated"
	}
	return "ready"
}

func publicSafe(in inputs) bool {
	return in.Authority.PublicRepoMode &&
		in.Authority.PublicSafetyOK &&
		!in.Authority.SelfAuthorityAllowed
}

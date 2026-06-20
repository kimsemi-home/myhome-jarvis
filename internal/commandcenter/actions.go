package commandcenter

func nextSafeAction(status Status) string {
	if !status.PublicSafe {
		return "run_public_safety_review"
	}
	for _, gate := range status.BlockedGates {
		switch gate.Key {
		case "authority":
			if status.AuthorityReview.NextSafeAction != "" &&
				status.AuthorityReview.NextSafeAction != "none" {
				return status.AuthorityReview.NextSafeAction
			}
			if status.AuthorityReview.ReviewRequestable {
				return "request_authority_review"
			}
			return "resolve_authority_gate"
		case "review":
			return "assign_or_reduce_human_review"
		case "finance_consent":
			return "record_finance_consent"
		case "cost":
			return "review_codex_cost_budget"
		case "codex_sustainability":
			return "review_codex_sustainability_evidence"
		case "context_pack":
			return "review_context_pack_handoff"
		case "monetization":
			return "review_monetization_experiments"
		case "repo_factory":
			return "review_repo_factory_bootstrap"
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
		!in.Authority.SelfAuthorityAllowed &&
		in.RepoFactory.PublicSafe
}

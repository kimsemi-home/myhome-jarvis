package codexcost

var requiredUnitKinds = []string{
	"codex_tokens",
	"codex_coin",
	"github_actions_minutes",
	"external_tool_cost",
}

var requiredLoopScopes = []string{
	"assistant_loop",
	"linear_project",
	"repo",
	"monetization_experiment",
}

var requiredRecordFields = []string{
	"at",
	"scope",
	"unit_kind",
	"amount",
	"status",
	"evidence_refs",
}

var requiredSummaryFields = []string{
	"record_count",
	"invalid_record_count",
	"total_units",
	"review_required_count",
	"missing_evidence_count",
	"budget_state",
	"by_unit_kind",
	"by_scope",
	"checked_at",
}

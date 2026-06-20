package monetization

var requiredStates = []string{"backlog", "review_required", "running", "closed"}
var requiredDecisions = []string{"hypothesis_created", "scale_requested", "close_experiment"}
var requiredReviews = []string{"requested", "approved", "rejected", "not_required"}
var requiredBands = []string{"unknown", "low", "medium", "high"}
var requiredCostUnits = []string{"codex_tokens", "codex_coin", "github_actions_minutes", "external_tool_cost"}

var requiredFields = []string{
	"at", "experiment_id", "hypothesis_key", "state", "decision_kind",
	"review_status", "expected_value_band", "cost_estimate_units",
	"cost_unit_kind", "evidence_refs",
}

var requiredSummaryFields = []string{
	"experiment_count", "decision_count", "invalid_record_count",
	"review_required_count", "missing_evidence_count",
	"missing_cost_estimate_count", "expected_value_unknown_count",
	"by_state", "by_decision_kind", "by_review_status",
	"by_expected_value_band", "checked_at",
}

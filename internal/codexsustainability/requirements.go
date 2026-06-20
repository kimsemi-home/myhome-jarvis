package codexsustainability

var requiredRecordKinds = []string{
	"usage_sample", "cycle_sample", "trend_baseline", "feature_proposal",
}

var requiredMetrics = []string{
	"codex_tokens", "codex_coin", "github_actions_minutes",
	"elapsed_cycle_minutes", "rework_count", "cache_hit_count",
	"cache_miss_count", "validation_failure_count", "human_review_debt",
	"accepted_change_count", "cache_savings_units",
}

var requiredFields = []string{"at", "record_kind", "evidence_refs"}

var requiredProposalFields = []string{
	"proposal_id", "evidence_refs", "cost_per_accepted_change",
	"median_cycle_minutes", "cache_savings_units", "defect_rework_rate",
	"monetization_ref",
}

var requiredSummaryFields = []string{
	"record_count", "trend_posture", "sustainability_posture",
	"evidence_freshness", "review_gate_count", "cost_per_accepted_change",
	"cache_savings_units", "rework_count",
}

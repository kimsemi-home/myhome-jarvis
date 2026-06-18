package review

var requiredOverloadRules = []string{
	"low_risk_only",
	"deterministic_verification",
	"evidence_collection",
	"incident_response",
	"revalidation",
	"high_risk_change",
	"major_ontology_change",
	"security_boundary_change",
	"production_change",
}

var frozenOverloadRules = []string{
	"high_risk_change",
	"major_ontology_change",
	"security_boundary_change",
	"production_change",
}

var requiredSummaryFields = []string{
	"count",
	"open_count",
	"high_risk_open_count",
	"invalid_review_count",
	"missing_evidence_count",
	"missing_reviewer_count",
	"backup_available_count",
	"review_debt_count",
	"capacity_state",
	"active_rule",
	"by_risk",
	"by_status",
	"by_reviewer_role",
	"by_queue_class",
	"checked_at",
}

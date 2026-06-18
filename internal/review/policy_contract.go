package review

var requiredRisks = []string{"low", "medium", "high"}

var requiredQueueClasses = []string{
	"security_incident",
	"production_incident",
	"ssot_defect",
	"user_impact_regression",
	"authority_boundary_change",
	"release_blocking",
	"major_ontology_change",
	"spec_change",
	"documentation",
}

var requiredStatuses = []string{
	"requested",
	"assigned",
	"in_review",
	"approved",
	"rejected",
	"deferred",
	"escalated",
}

var requiredRequesterRoles = []string{
	"producer",
	"independent_reviewer",
	"adversarial_reviewer",
	"deterministic_verifier",
	"governance_steward",
}

var requiredReviewerRoles = []string{
	"independent_reviewer",
	"adversarial_reviewer",
	"deterministic_verifier",
	"governance_steward",
	"backup_steward",
}

var requiredReviewFields = []string{
	"at",
	"item_key",
	"queue_class",
	"risk",
	"status",
	"requester_role",
	"reviewer_role",
	"backup_available",
	"evidence_refs",
}

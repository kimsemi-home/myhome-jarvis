package pdca

var requiredSteps = []string{"plan", "do", "check", "act"}

var requiredFields = []string{
	"cycle_id", "at", "status", "owner",
	"plan_ref", "do_ref", "check_ref", "act_ref",
}

var requiredStatuses = []string{"open", "checking", "acting", "closed"}

var requiredSources = []string{
	"LearningLedger", "EvidenceGraph", "HumanReviewCapacity",
	"AuthorityGate", "VerificationEvidenceOps", "QualityLedger",
}

var requiredSummary = []string{
	"cycle_count", "open_count", "closed_count",
	"invalid_cycle_count", "ready_step_count", "ready",
}

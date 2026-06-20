package financeconsent

var requiredConsentKinds = []string{
	"finance_connector",
	"spouse_scope",
	"household_scope",
}

var requiredConsentStatuses = []string{
	"requested",
	"granted",
	"revoked",
	"expired",
	"denied",
}

var requiredReviewStatuses = []string{
	"requested",
	"approved",
	"rejected",
}

var requiredFields = []string{
	"at",
	"consent_kind",
	"subject_scope",
	"status",
	"review_status",
	"authority_profile",
	"evidence_refs",
}

var requiredSummaryFields = []string{
	"readiness_state",
	"finance_mode",
	"exists",
	"record_count",
	"active_consent_count",
	"missing_required_consent_count",
	"review_required_count",
	"missing_evidence_count",
	"invalid_record_count",
	"revoked_or_expired_count",
	"forbidden_action_enabled_count",
	"consent_debt_count",
	"checked_at",
}

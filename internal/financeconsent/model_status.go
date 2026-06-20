package financeconsent

type Status struct {
	ReadinessState              string `json:"readiness_state"`
	FinanceMode                 string `json:"finance_mode"`
	Exists                      bool   `json:"exists"`
	RecordCount                 int    `json:"record_count"`
	ActiveConsentCount          int    `json:"active_consent_count"`
	MissingRequiredConsentCount int    `json:"missing_required_consent_count"`
	ReviewRequiredCount         int    `json:"review_required_count"`
	MissingEvidenceCount        int    `json:"missing_evidence_count"`
	InvalidRecordCount          int    `json:"invalid_record_count"`
	RevokedOrExpiredCount       int    `json:"revoked_or_expired_count"`
	ForbiddenActionEnabledCount int    `json:"forbidden_action_enabled_count"`
	ConsentDebtCount            int    `json:"consent_debt_count"`
	CheckedAt                   string `json:"checked_at"`
}

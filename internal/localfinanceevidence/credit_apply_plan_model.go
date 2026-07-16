package localfinanceevidence

type CreditBatchApplyPlan struct {
	SchemaVersion           string   `json:"schema_version"`
	ExecutionMode           string   `json:"execution_mode"`
	CredentialsRead         bool     `json:"credentials_read"`
	ExternalNetwork         bool     `json:"external_network"`
	ExternalWrites          bool     `json:"external_writes"`
	ManifestSHA256          string   `json:"manifest_sha256"`
	PreviewSetHash          string   `json:"preview_set_hash"`
	BatchHash               string   `json:"batch_hash"`
	StatementCount          int      `json:"statement_count"`
	ApprovalChallengeSHA256 string   `json:"approval_challenge_sha256"`
	Checks                  []string `json:"checks"`
	PlanHash                string   `json:"plan_hash"`
}

type CreditBatchApproval struct {
	SchemaVersion           string `json:"schema_version"`
	Decision                string `json:"decision"`
	PlanHash                string `json:"plan_hash"`
	ApprovalChallengeSHA256 string `json:"approval_challenge_sha256"`
	ManifestSHA256          string `json:"manifest_sha256"`
	PreviewSetHash          string `json:"preview_set_hash"`
	BatchHash               string `json:"batch_hash"`
	ApprovalHash            string `json:"approval_hash"`
}

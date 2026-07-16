package localfinanceevidence

type CreditBatchRollbackState struct {
	Accounts            int `json:"accounts"`
	Transactions        int `json:"transactions"`
	ImportRuns          int `json:"import_runs"`
	Templates           int `json:"templates"`
	CategorySuggestions int `json:"category_suggestions"`
}

type CreditBatchApplyRehearsal struct {
	SchemaVersion                string                   `json:"schema_version"`
	ExecutionMode                string                   `json:"execution_mode"`
	CredentialsRead              bool                     `json:"credentials_read"`
	ExternalNetwork              bool                     `json:"external_network"`
	ExternalWrites               bool                     `json:"external_writes"`
	EphemeralLocalDatabaseWrites bool                     `json:"ephemeral_local_database_writes"`
	PersistentDatabaseWrites     bool                     `json:"persistent_database_writes"`
	Plan                         CreditBatchApplyPlan     `json:"apply_plan"`
	Approval                     CreditBatchApproval      `json:"approval"`
	FirstApply                   CreditBatchApplyReport   `json:"first_apply"`
	Replay                       CreditBatchApplyReport   `json:"replay"`
	StaleApprovalRejected        bool                     `json:"stale_approval_rejected"`
	DeniedApprovalRejected       bool                     `json:"denied_approval_rejected"`
	MidBatchFailureRejected      bool                     `json:"mid_batch_failure_rejected"`
	RollbackState                CreditBatchRollbackState `json:"rollback_state"`
	Checks                       []string                 `json:"checks"`
	ReportHash                   string                   `json:"report_hash"`
}

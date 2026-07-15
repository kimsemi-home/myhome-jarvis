package shortsfactory

type Privacy struct {
	ContainsCredentials        bool `json:"contains_credentials"`
	ContainsAccountIdentifiers bool `json:"contains_account_identifiers"`
	ContainsRawPrivateEvidence bool `json:"contains_raw_private_evidence"`
	ContainsRevenueDetails     bool `json:"contains_revenue_details"`
}

type YouTubeConsent struct {
	Required   bool   `json:"required"`
	ReceiptRef string `json:"receipt_ref"`
	Visibility string `json:"visibility"`
}

type InputHashes struct {
	Content        string `json:"content"`
	Evidence       string `json:"evidence"`
	Assets         string `json:"assets"`
	ExecutionPlan  string `json:"execution_plan"`
	ChannelBinding string `json:"channel_binding"`
	ConsentReceipt string `json:"consent_receipt"`
}

type GateResult struct {
	SchemaVersion         string      `json:"schema_version"`
	Decision              string      `json:"decision"`
	CriteriaVersion       string      `json:"criteria_version"`
	InputHash             string      `json:"input_hash"`
	ReceiptRef            string      `json:"receipt_ref,omitempty"`
	ReleasedOpenLoopSteps []string    `json:"released_open_loop_steps,omitempty"`
	Criteria              []Criterion `json:"criteria"`
}

type Criterion struct {
	ID     string `json:"id"`
	Passed bool   `json:"passed"`
}

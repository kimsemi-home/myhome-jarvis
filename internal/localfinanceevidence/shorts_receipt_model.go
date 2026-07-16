package localfinanceevidence

type ShortsReceipt struct {
	SchemaVersion     string `json:"schema_version"`
	ExecutionMode     string `json:"execution_mode"`
	LoopbackOnly      bool   `json:"loopback_only"`
	ExternalWrites    bool   `json:"external_writes"`
	PlanHash          string `json:"plan_hash"`
	ContentHash       string `json:"content_hash"`
	SessionRequests   int    `json:"session_requests"`
	ProbeRequests     int    `json:"probe_requests"`
	ChunkRequests     int    `json:"chunk_requests"`
	RecoveryAttempts  int    `json:"recovery_attempts"`
	AcceptedBytes     int64  `json:"accepted_bytes"`
	Complete          bool   `json:"complete"`
	PrivacyStatus     string `json:"privacy_status"`
	VideoIdentityHash string `json:"video_identity_hash"`
	ReceiptHash       string `json:"receipt_hash"`
}

type ShortsEmulator struct {
	SessionCreates      int `json:"session_creates"`
	VideoCreates        int `json:"video_creates"`
	InterruptedRequests int `json:"interrupted_requests"`
	TokenRequests       int `json:"token_requests"`
	ChannelRequests     int `json:"channel_requests"`
	RedirectRequests    int `json:"redirect_requests"`
	OversizedRequests   int `json:"oversized_requests"`
}

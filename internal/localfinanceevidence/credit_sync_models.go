package localfinanceevidence

type CreditOAuth struct {
	SchemaVersion             string `json:"schema_version"`
	ExecutionMode             string `json:"execution_mode"`
	LoopbackOnly              bool   `json:"loopback_only"`
	CredentialsRead           bool   `json:"credentials_read"`
	ExternalNetwork           bool   `json:"external_network"`
	ExternalWrites            bool   `json:"external_writes"`
	AuthorizationCodeExchange bool   `json:"authorization_code_exchange"`
	RefreshTokenExchange      bool   `json:"refresh_token_exchange"`
	OfficialOriginPinned      bool   `json:"official_origin_pinned"`
	RedirectRejected          bool   `json:"redirect_rejected"`
	OversizedResponseRejected bool   `json:"oversized_response_rejected"`
	TokenRequests             int    `json:"token_requests"`
	RedirectRequests          int    `json:"redirect_requests"`
	OversizedRequests         int    `json:"oversized_requests"`
}

type CreditSync struct {
	SchemaVersion       string `json:"schema_version"`
	Sources             int    `json:"sources"`
	MessagesInspected   int    `json:"messages_inspected"`
	AllowedMessages     int    `json:"allowed_messages"`
	AttachmentsWritten  int    `json:"attachments_written"`
	Duplicates          int    `json:"duplicates"`
	PreviouslyProcessed int    `json:"previously_processed"`
	ReceiptsWritten     int    `json:"receipts_written"`
	Retries             int    `json:"retries"`
}

type CreditWatch struct {
	Files      int `json:"files"`
	Inserted   int `json:"inserted"`
	Duplicates int `json:"duplicates"`
}

type CreditMonthly struct {
	TransactionCount  int   `json:"transaction_count"`
	CardSpendMinor    int64 `json:"card_spend_minor"`
	CardRefundMinor   int64 `json:"card_refund_minor"`
	NetCardSpendMinor int64 `json:"net_card_spend_minor"`
	Reconciled        bool  `json:"reconciled"`
}

type CreditEmulator struct {
	ListRequests       int `json:"list_requests"`
	MessageRequests    int `json:"message_requests"`
	AttachmentRequests int `json:"attachment_requests"`
	InjectedFailures   int `json:"injected_failures"`
}

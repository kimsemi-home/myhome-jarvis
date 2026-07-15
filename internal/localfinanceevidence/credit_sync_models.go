package localfinanceevidence

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

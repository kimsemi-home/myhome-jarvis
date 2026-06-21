package externalevidence

type manifestRecord struct {
	At                      string `json:"at"`
	Source                  string `json:"source"`
	Kind                    string `json:"kind"`
	EvidenceRef             string `json:"evidence_ref"`
	SourceClass             string `json:"source_class"`
	Status                  string `json:"status"`
	HTTPStatus              int    `json:"http_status"`
	PayloadBytes            int    `json:"payload_bytes"`
	RawSHA256               string `json:"raw_sha256"`
	RawPrivatePath          string `json:"raw_private_path"`
	Preprocess              string `json:"preprocess"`
	FreshnessHours          int    `json:"freshness_hours"`
	RawPayloadPublicAllowed bool   `json:"raw_payload_public_allowed"`
}

type normalizedRecord struct {
	At             string `json:"at"`
	Source         string `json:"source"`
	Kind           string `json:"kind"`
	EvidenceRef    string `json:"evidence_ref"`
	SourceClass    string `json:"source_class"`
	PayloadBytes   int    `json:"payload_bytes"`
	HTTPStatus     int    `json:"http_status"`
	Preprocess     string `json:"preprocess"`
	FreshnessHours int    `json:"freshness_hours"`
}

type goldSummaryRecord struct {
	At             string `json:"at"`
	Source         string `json:"source"`
	Kind           string `json:"kind"`
	EvidenceRef    string `json:"evidence_ref"`
	SourceCount    int    `json:"source_count"`
	CollectedCount int    `json:"collected_count"`
	CachedCount    int    `json:"cached_count"`
	FailedCount    int    `json:"failed_count"`
}

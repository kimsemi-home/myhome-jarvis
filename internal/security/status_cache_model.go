package security

const (
	statusCachePath       = "data/private/security/status-cache.json"
	statusCacheVersion    = "security-status-cache/v1"
	statusCacheValidation = "mhj security history"
)

type statusCacheKey struct {
	Head      string
	InputHash string
	Key       string
}

type statusCacheRecord struct {
	Version             string `json:"version"`
	Key                 string `json:"key"`
	Head                string `json:"head"`
	InputHash           string `json:"input_hash"`
	HistoryOK           bool   `json:"history_ok"`
	HistoryFindingCount int    `json:"history_finding_count"`
	CheckedAt           string `json:"checked_at"`
	ValidationCommand   string `json:"validation_command"`
}

type historyAggregate struct {
	OK           bool
	FindingCount int
}

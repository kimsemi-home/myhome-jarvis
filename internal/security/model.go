package security

type Finding struct {
	Path    string `json:"path"`
	Line    int    `json:"line,omitempty"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Report struct {
	Root     string    `json:"root"`
	OK       bool      `json:"ok"`
	Findings []Finding `json:"findings"`
}

type HistoryFinding struct {
	Commit  string `json:"commit,omitempty"`
	Path    string `json:"path"`
	Line    int    `json:"line,omitempty"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type HistoryReport struct {
	Root     string           `json:"root"`
	OK       bool             `json:"ok"`
	Findings []HistoryFinding `json:"findings"`
}

type Status struct {
	OK                  bool          `json:"ok"`
	CurrentOK           bool          `json:"current_ok"`
	CurrentFindingCount int           `json:"current_finding_count"`
	HistoryOK           bool          `json:"history_ok"`
	HistoryFindingCount int           `json:"history_finding_count"`
	Cache               CacheEvidence `json:"cache"`
	CheckedAt           string        `json:"checked_at"`
}

type CacheEvidence struct {
	Path                    string `json:"path"`
	State                   string `json:"state"`
	Key                     string `json:"key"`
	InputHash               string `json:"input_hash"`
	Head                    string `json:"head"`
	EvidenceRef             string `json:"evidence_ref"`
	ValidationCommand       string `json:"validation_command"`
	PublicSafe              bool   `json:"public_safe"`
	RawDetailsPublicAllowed bool   `json:"raw_details_public_allowed"`
}

type historyPattern struct {
	Code    string
	Pattern string
	Message string
}

type gitCommandError struct {
	err     error
	message string
}

func (err gitCommandError) Error() string {
	return err.message
}

func (err gitCommandError) Unwrap() error {
	return err.err
}

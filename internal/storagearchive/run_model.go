package storagearchive

type RunReport struct {
	PolicyPath              string      `json:"policy_path"`
	ArchiveRoot             string      `json:"archive_root"`
	ManifestPath            string      `json:"manifest_path"`
	SourceCount             int         `json:"source_count"`
	ArchivedCount           int         `json:"archived_count"`
	SkippedCount            int         `json:"skipped_count"`
	BudgetBreachCount       int         `json:"budget_breach_count"`
	PublicSafe              bool        `json:"public_safe"`
	RawPayloadPublicAllowed bool        `json:"raw_payload_public_allowed"`
	Results                 []RunResult `json:"results"`
	CheckedAt               string      `json:"checked_at"`
}

type RunResult struct {
	SourceKey         string `json:"source_key"`
	SourcePath        string `json:"source_path"`
	State             string `json:"state"`
	ArchivePath       string `json:"archive_path,omitempty"`
	InputBytes        int64  `json:"input_bytes"`
	OutputBytes       int64  `json:"output_bytes"`
	InputSHA256       string `json:"input_sha256,omitempty"`
	RecordCount       int    `json:"record_count"`
	NoiseCount        int    `json:"noise_count"`
	NoiseRatioPercent int    `json:"noise_ratio_percent"`
	BudgetOK          bool   `json:"budget_ok"`
}

type manifestEntry struct {
	At                string `json:"at"`
	SourceKey         string `json:"source_key"`
	SourcePath        string `json:"source_path"`
	ArchivePath       string `json:"archive_path,omitempty"`
	State             string `json:"state"`
	InputBytes        int64  `json:"input_bytes"`
	OutputBytes       int64  `json:"output_bytes"`
	InputSHA256       string `json:"input_sha256,omitempty"`
	RecordCount       int    `json:"record_count"`
	NoiseCount        int    `json:"noise_count"`
	NoiseRatioPercent int    `json:"noise_ratio_percent"`
	BudgetVerdict     string `json:"budget_verdict"`
	RawPayloadStored  bool   `json:"raw_payload_stored"`
}

type sourceScan struct {
	Content           []byte
	InputSHA256       string
	RecordCount       int
	NoiseCount        int
	NoiseRatioPercent int
	BudgetOK          bool
}

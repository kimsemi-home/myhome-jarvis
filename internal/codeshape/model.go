package codeshape

const PolicyRelativePath = "generated/code_shape.generated.json"

type Policy struct {
	Context               string           `json:"context"`
	Version               string           `json:"version"`
	GeneratedArtifact     string           `json:"generated_artifact"`
	MaxFileLines          int              `json:"max_file_lines"`
	PublicStatusRedacted  bool             `json:"public_status_redacted"`
	SourceRoots           []string         `json:"source_roots"`
	Extensions            []string         `json:"extensions"`
	ExcludedPrefixes      []string         `json:"excluded_prefixes"`
	LegacyDebtFiles       []LegacyDebtFile `json:"legacy_debt_files"`
	PublicSummaryFields   []string         `json:"public_summary_fields"`
	ForbiddenPublicFields []string         `json:"forbidden_public_fields"`
	Commands              []string         `json:"commands"`
}

type LegacyDebtFile struct {
	Path     string `json:"path"`
	MaxLines int    `json:"max_lines"`
}

type FileFinding struct {
	Path           string `json:"path"`
	Lines          int    `json:"lines"`
	MaxLines       int    `json:"max_lines"`
	LegacyMaxLines int    `json:"legacy_max_lines,omitempty"`
}

type Status struct {
	PolicyPath            string        `json:"policy_path"`
	MaxFileLines          int           `json:"max_file_lines"`
	FileCount             int           `json:"file_count"`
	OverBudgetCount       int           `json:"over_budget_count"`
	LegacyDebtCount       int           `json:"legacy_debt_count"`
	BudgetRegressionCount int           `json:"budget_regression_count"`
	MaxObservedPath       string        `json:"max_observed_path"`
	MaxObservedLines      int           `json:"max_observed_lines"`
	TopDebt               []FileFinding `json:"top_debt"`
	Regressions           []FileFinding `json:"regressions"`
	OK                    bool          `json:"ok"`
	CheckedAt             string        `json:"checked_at"`
}

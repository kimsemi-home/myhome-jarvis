package translation

type Status struct {
	PolicyPath           string         `json:"policy_path"`
	LedgerPath           string         `json:"ledger_path"`
	ManifestRoot         string         `json:"manifest_root"`
	LedgerExists         bool           `json:"ledger_exists"`
	ManifestRootExists   bool           `json:"manifest_root_exists"`
	ManifestCount        int            `json:"manifest_count"`
	InvalidManifestCount int            `json:"invalid_manifest_count"`
	MissingManifestCount int            `json:"missing_manifest_count"`
	LossCount            int            `json:"loss_count"`
	OpenLossCount        int            `json:"open_loss_count"`
	ClosedLossCount      int            `json:"closed_loss_count"`
	InvalidLossCount     int            `json:"invalid_loss_count"`
	OpenDebtCount        int            `json:"open_debt_count"`
	ForbiddenLossCount   int            `json:"forbidden_loss_count"`
	ByLevel              map[string]int `json:"by_level"`
	BySourceContext      map[string]int `json:"by_source_context"`
	ByTargetContext      map[string]int `json:"by_target_context"`
	LastObservedAt       string         `json:"last_observed_at,omitempty"`
	CheckedAt            string         `json:"checked_at"`
}

package controlplane

type Status struct {
	PolicyPath                 string         `json:"policy_path"`
	ManifestPath               string         `json:"manifest_path"`
	Exists                     bool           `json:"exists"`
	Count                      int            `json:"count"`
	InvalidManifestCount       int            `json:"invalid_manifest_count"`
	ManifestDebtCount          int            `json:"manifest_debt_count"`
	VerifierSeparationRequired bool           `json:"verifier_separation_required"`
	VerifierViolationCount     int            `json:"verifier_violation_count"`
	MinLeaseSeconds            int            `json:"min_lease_seconds"`
	MaxLeaseSeconds            int            `json:"max_lease_seconds"`
	ByDecisionKind             map[string]int `json:"by_decision_kind"`
	ByAuthorityProfile         map[string]int `json:"by_authority_profile"`
	ByLeaseStatus              map[string]int `json:"by_lease_status"`
	LastObservedAt             string         `json:"last_observed_at,omitempty"`
	CheckedAt                  string         `json:"checked_at"`
}

package contextpack

type AuthorityContract struct {
	Path                     string `json:"path"`
	Version                  string `json:"version"`
	SelfApprovalAllowed      bool   `json:"self_approval_allowed"`
	PublicSafetyGateRequired bool   `json:"public_safety_gate_required"`
}

type SecurityContract struct {
	Path                      string `json:"path"`
	Version                   string `json:"version"`
	PrivatePathsPublicAllowed bool   `json:"private_paths_public_allowed"`
	LocalPathsPublicAllowed   bool   `json:"local_paths_public_allowed"`
}

type VerificationProfile struct {
	Name          string   `json:"name"`
	Graph         string   `json:"graph"`
	RequiredUnits []string `json:"required_units"`
}

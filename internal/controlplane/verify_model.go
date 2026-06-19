package controlplane

const VerificationRelativePath = "generated/control_plane_verification.generated.json"

type VerificationPolicy struct {
	SchemaVersion              string              `json:"schema_version"`
	Source                     string              `json:"source"`
	PolicyArtifact             string              `json:"policy_artifact"`
	Command                    string              `json:"command"`
	StatusCommand              string              `json:"status_command"`
	VerifierSeparationRequired bool                `json:"verifier_separation_required"`
	Checks                     []VerificationCheck `json:"checks"`
}

type VerificationCheck struct {
	ID       string `json:"id"`
	Evidence string `json:"evidence"`
}

type VerificationStatus struct {
	OK                     bool   `json:"ok"`
	PolicyPath             string `json:"policy_path"`
	VerifierPath           string `json:"verifier_path"`
	CheckCount             int    `json:"check_count"`
	ManifestDebtCount      int    `json:"manifest_debt_count"`
	VerifierViolationCount int    `json:"verifier_violation_count"`
	CheckedAt              string `json:"checked_at"`
}

package codexcost

type GuardRequest struct {
	Scope            string   `json:"scope"`
	UnitKind         string   `json:"unit_kind"`
	EstimatedUnits   int64    `json:"estimated_units"`
	EstimatedMinutes int64    `json:"estimated_minutes"`
	EvidenceRefs     []string `json:"evidence_refs"`
}

type GuardResult struct {
	Decision              string   `json:"decision"`
	Reasons               []string `json:"reasons,omitempty"`
	Scope                 string   `json:"scope"`
	UnitKind              string   `json:"unit_kind"`
	EstimatedUnits        int64    `json:"estimated_units"`
	EstimatedMinutes      int64    `json:"estimated_minutes"`
	CurrentBudgetState    string   `json:"current_budget_state"`
	ProjectedBudgetState  string   `json:"projected_budget_state"`
	SustainabilityPosture string   `json:"sustainability_posture"`
	ReviewGateCount       int      `json:"review_gate_count"`
	EvidenceRefCount      int      `json:"evidence_ref_count"`
}

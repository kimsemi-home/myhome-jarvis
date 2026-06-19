package incidents

type Policy struct {
	Context                   string   `json:"context"`
	Version                   string   `json:"version"`
	GeneratedArtifact         string   `json:"generated_artifact"`
	PrivateIncidentLedger     string   `json:"private_incident_ledger"`
	AppendOnly                bool     `json:"append_only"`
	PublicStatusRedacted      bool     `json:"public_status_redacted"`
	RawIncidentPublicAllowed  bool     `json:"raw_incident_public_allowed"`
	QuarantineStaleAfterHours int      `json:"quarantine_stale_after_hours"`
	AllowedKinds              []string `json:"allowed_kinds"`
	Lifecycle                 []string `json:"lifecycle"`
	AllowedStatuses           []string `json:"allowed_statuses"`
	OwnerRoles                []string `json:"owner_roles"`
	QuarantineStates          []string `json:"quarantine_states"`
	RequiredFields            []string `json:"required_fields"`
	AllowedEvidencePrefixes   []string `json:"allowed_evidence_prefixes"`
	PublicSummaryFields       []string `json:"public_summary_fields"`
	ForbiddenPublicFields     []string `json:"forbidden_public_fields"`
	Commands                  []string `json:"commands"`
}

type Incident struct {
	ID              string   `json:"id"`
	At              string   `json:"at"`
	Kind            string   `json:"kind"`
	Stage           string   `json:"stage"`
	Status          string   `json:"status"`
	OwnerRole       string   `json:"owner_role"`
	QuarantineState string   `json:"quarantine_state"`
	EvidenceRefs    []string `json:"evidence_refs"`
}

type Status struct {
	PolicyPath                string         `json:"policy_path"`
	LedgerPath                string         `json:"ledger_path"`
	Exists                    bool           `json:"exists"`
	Count                     int            `json:"count"`
	OpenCount                 int            `json:"open_count"`
	ClosedCount               int            `json:"closed_count"`
	InvalidIncidentCount      int            `json:"invalid_incident_count"`
	IncidentDebtCount         int            `json:"incident_debt_count"`
	MissingOwnerCount         int            `json:"missing_owner_count"`
	MissingEvidenceRefCount   int            `json:"missing_evidence_ref_count"`
	StaleQuarantineCount      int            `json:"stale_quarantine_count"`
	QuarantineStaleAfterHours int            `json:"quarantine_stale_after_hours"`
	ByKind                    map[string]int `json:"by_kind"`
	ByStage                   map[string]int `json:"by_stage"`
	ByStatus                  map[string]int `json:"by_status"`
	ByOwnerRole               map[string]int `json:"by_owner_role"`
	ByQuarantineState         map[string]int `json:"by_quarantine_state"`
	LastObservedAt            string         `json:"last_observed_at,omitempty"`
	CheckedAt                 string         `json:"checked_at"`
}

package pdca

const PolicyRelativePath = "generated/pdca.generated.json"

type Policy struct {
	Context               string   `json:"context"`
	Version               string   `json:"version"`
	GeneratedArtifact     string   `json:"generated_artifact"`
	PrivateCycleLedger    string   `json:"private_cycle_ledger"`
	AppendOnly            bool     `json:"append_only"`
	PublicStatusRedacted  bool     `json:"public_status_redacted"`
	RawCyclePublicAllowed bool     `json:"raw_cycle_public_allowed"`
	Steps                 []Step   `json:"steps"`
	RequiredFields        []string `json:"required_fields"`
	AllowedStatuses       []string `json:"allowed_statuses"`
	EvidenceSources       []string `json:"evidence_sources"`
	PublicSummaryFields   []string `json:"public_summary_fields"`
	ForbiddenPublicFields []string `json:"forbidden_public_fields"`
	Commands              []string `json:"commands"`
}

type Step struct {
	ID       string `json:"id"`
	Role     string `json:"role"`
	Artifact string `json:"artifact"`
	Command  string `json:"command"`
}

type Cycle struct {
	CycleID  string `json:"cycle_id"`
	At       string `json:"at"`
	Status   string `json:"status"`
	Owner    string `json:"owner"`
	PlanRef  string `json:"plan_ref"`
	DoRef    string `json:"do_ref"`
	CheckRef string `json:"check_ref"`
	ActRef   string `json:"act_ref"`
}

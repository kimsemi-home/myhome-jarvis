package linear

const WriteEvidenceRelativePath = "data/private/linear-write-evidence.jsonl"

type WriteEvidence struct {
	At       string `json:"at"`
	Action   string `json:"action"`
	IssueKey string `json:"issue_key,omitempty"`
	Synced   bool   `json:"synced"`
}

type WriteEvidenceStatus struct {
	EvidencePath         string         `json:"evidence_path"`
	SyncedMutationCount  int            `json:"synced_mutation_count"`
	HasSyncedMutation    bool           `json:"has_synced_mutation"`
	LatestSyncedMutation *WriteEvidence `json:"latest_synced_mutation,omitempty"`
}

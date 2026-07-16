package localfinanceevidence

const (
	ReceiptSchema  = "myhome.indirect-evidence/v1"
	ManifestSchema = "myhome.local-finance-evidence-manifest/v10"
)

type Receipt struct {
	SchemaVersion         string   `json:"schema_version"`
	Component             string   `json:"component"`
	Capability            string   `json:"capability"`
	ExecutionMode         string   `json:"execution_mode"`
	ExternalWritesEnabled bool     `json:"external_writes_enabled"`
	ArtifactHash          string   `json:"artifact_hash"`
	Checks                []string `json:"checks"`
	ReceiptHash           string   `json:"receipt_hash"`
}

type Manifest struct {
	SchemaVersion         string     `json:"schema_version"`
	Month                 string     `json:"month"`
	ExternalWritesEnabled bool       `json:"external_writes_enabled"`
	Receipts              []Receipt  `json:"receipts"`
	ExecutionProofs       []ProofRef `json:"execution_proofs"`
	AggregateHash         string     `json:"aggregate_hash"`
}

type ProofRef struct {
	Component      string `json:"component"`
	Capability     string `json:"capability"`
	ProofSchema    string `json:"proof_schema"`
	Path           string `json:"path"`
	ArtifactSHA256 string `json:"artifact_sha256"`
	ReportHash     string `json:"report_hash"`
}

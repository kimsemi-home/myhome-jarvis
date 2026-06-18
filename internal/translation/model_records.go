package translation

type lossRecord struct {
	At            string   `json:"at"`
	SourceContext string   `json:"source_context"`
	TargetContext string   `json:"target_context"`
	Level         string   `json:"level"`
	Category      string   `json:"category"`
	Status        string   `json:"status"`
	ManifestPath  string   `json:"manifest_path"`
	EvidenceRefs  []string `json:"evidence_refs"`
}

type manifest struct {
	SourceContext  string      `json:"source_context"`
	TargetContext  string      `json:"target_context"`
	SourceVersion  string      `json:"source_version"`
	TargetVersion  string      `json:"target_version"`
	PreservedRules []string    `json:"preserved_rules"`
	KnownLosses    []knownLoss `json:"known_losses"`
	Owner          string      `json:"owner"`
	EvidenceRefs   []string    `json:"evidence_refs"`
}

type knownLoss struct {
	Level    string `json:"level"`
	Category string `json:"category"`
}

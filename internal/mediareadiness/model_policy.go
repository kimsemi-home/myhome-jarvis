package mediareadiness

type Policy struct {
	Context                 string          `json:"context"`
	Version                 string          `json:"version"`
	GeneratedArtifact       string          `json:"generated_artifact"`
	BenchmarkKind           string          `json:"benchmark_kind"`
	PublicStatusRedacted    bool            `json:"public_status_redacted"`
	ExecuteCommands         bool            `json:"execute_commands"`
	PersistPayloads         bool            `json:"persist_payloads"`
	PersistURLs             bool            `json:"persist_urls"`
	TargetPlanningLatencyMS int64           `json:"target_planning_latency_ms"`
	Cases                   []BenchmarkCase `json:"cases"`
	PublicSummaryFields     []string        `json:"public_summary_fields"`
	ForbiddenPublicFields   []string        `json:"forbidden_public_fields"`
	Commands                []string        `json:"commands"`
}

type BenchmarkCase struct {
	ID           string `json:"id"`
	Capability   string `json:"capability"`
	Command      string `json:"command"`
	PayloadKind  string `json:"payload_kind"`
	ExpectedHost string `json:"expected_host"`
}

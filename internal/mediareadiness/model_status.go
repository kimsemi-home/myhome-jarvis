package mediareadiness

type Status struct {
	Context                 string       `json:"context"`
	Version                 string       `json:"version"`
	PolicyPath              string       `json:"policy_path"`
	BenchmarkKind           string       `json:"benchmark_kind"`
	PublicSafe              bool         `json:"public_safe"`
	Redaction               string       `json:"redaction"`
	CaseCount               int          `json:"case_count"`
	AvailableCount          int          `json:"available_count"`
	DegradedCount           int          `json:"degraded_count"`
	MaxPlanningLatencyMS    int64        `json:"max_planning_latency_ms"`
	TargetPlanningLatencyMS int64        `json:"target_planning_latency_ms"`
	LocalLauncherAvailable  bool         `json:"local_launcher_available"`
	LocalLauncherProbe      string       `json:"local_launcher_probe"`
	Cases                   []CaseStatus `json:"cases"`
	CheckedAt               string       `json:"checked_at"`
}

type CaseStatus struct {
	ID                string `json:"id"`
	Capability        string `json:"capability"`
	Command           string `json:"command"`
	Available         bool   `json:"available"`
	Availability      string `json:"availability"`
	PlanningLatencyMS int64  `json:"planning_latency_ms"`
	InvocationCount   int    `json:"invocation_count"`
	ExpectedHost      string `json:"expected_host"`
}

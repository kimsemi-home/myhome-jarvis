package commandcenter

type MediaReadinessSummary struct {
	PublicSafe              bool  `json:"public_safe"`
	CaseCount               int   `json:"case_count"`
	AvailableCount          int   `json:"available_count"`
	DegradedCount           int   `json:"degraded_count"`
	MaxPlanningLatencyMS    int64 `json:"max_planning_latency_ms"`
	TargetPlanningLatencyMS int64 `json:"target_planning_latency_ms"`
	LocalLauncherAvailable  bool  `json:"local_launcher_available"`
}

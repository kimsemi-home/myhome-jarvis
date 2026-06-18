package qualitylog

const RelativePath = "data/private/quality/runs.jsonl"

type Step struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Run struct {
	At             string `json:"at"`
	OK             bool   `json:"ok"`
	DurationMillis int64  `json:"duration_millis"`
	StepCount      int    `json:"step_count"`
	PassCount      int    `json:"pass_count"`
	FailCount      int    `json:"fail_count"`
	SkipCount      int    `json:"skip_count"`
	Steps          []Step `json:"steps"`
}

type Status struct {
	Path      string `json:"path"`
	Exists    bool   `json:"exists"`
	Count     int    `json:"count"`
	Last      *Run   `json:"last,omitempty"`
	CheckedAt string `json:"checked_at"`
}

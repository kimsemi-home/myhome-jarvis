package learning

import "time"

type Status struct {
	Path           string         `json:"path"`
	PolicyPath     string         `json:"policy_path"`
	Exists         bool           `json:"exists"`
	Count          int            `json:"count"`
	OpenCount      int            `json:"open_count"`
	ClosedCount    int            `json:"closed_count"`
	ByKind         map[string]int `json:"by_kind"`
	ByStage        map[string]int `json:"by_stage"`
	LastKind       string         `json:"last_kind,omitempty"`
	LastStage      string         `json:"last_stage,omitempty"`
	LastStatus     string         `json:"last_status,omitempty"`
	LastObservedAt string         `json:"last_observed_at,omitempty"`
	CheckedAt      string         `json:"checked_at"`
}

func newStatus(policy Policy) Status {
	return Status{
		Path:       policy.PrivateLedger,
		PolicyPath: PolicyRelativePath,
		ByKind:     map[string]int{},
		ByStage:    map[string]int{},
		CheckedAt:  time.Now().UTC().Format(time.RFC3339),
	}
}

func applyObservation(status *Status, observation Observation) {
	status.Exists = true
	status.Count++
	status.ByKind[observation.Kind]++
	status.ByStage[observation.Stage]++
	if observation.Status == "closed" {
		status.ClosedCount++
	} else {
		status.OpenCount++
	}
	status.LastKind = observation.Kind
	status.LastStage = observation.Stage
	status.LastStatus = observation.Status
	status.LastObservedAt = observation.At
}

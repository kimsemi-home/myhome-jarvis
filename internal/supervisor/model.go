package supervisor

type DaemonState struct {
	Name           string `json:"name"`
	PID            int    `json:"pid"`
	Host           string `json:"host"`
	Port           int    `json:"port"`
	Address        string `json:"address"`
	Version        string `json:"version"`
	ExecuteEnabled bool   `json:"execute_enabled"`
	LANBindAllowed bool   `json:"lan_bind_allowed"`
	StartedAt      string `json:"started_at"`
	UpdatedAt      string `json:"updated_at"`
}

type DaemonStatus struct {
	Name           string `json:"name"`
	Recorded       bool   `json:"recorded"`
	StatePath      string `json:"state_path"`
	PID            int    `json:"pid,omitempty"`
	Address        string `json:"address,omitempty"`
	Version        string `json:"version,omitempty"`
	StartedAt      string `json:"started_at,omitempty"`
	UpdatedAt      string `json:"updated_at,omitempty"`
	ProcessRunning bool   `json:"process_running"`
	ProbeOK        bool   `json:"probe_ok"`
	ProbeStatus    int    `json:"probe_status,omitempty"`
	ProbeURL       string `json:"probe_url,omitempty"`
	Stale          bool   `json:"stale"`
	Message        string `json:"message"`
	CheckedAt      string `json:"checked_at"`
}

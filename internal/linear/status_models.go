package linear

type Status struct {
	Mode               string        `json:"mode"`
	TokenConfigured    bool          `json:"token_configured"`
	TokenSource        string        `json:"token_source,omitempty"`
	Synced             bool          `json:"synced"`
	QueuePath          string        `json:"queue_path"`
	Endpoint           string        `json:"endpoint,omitempty"`
	HTTPStatus         int           `json:"http_status,omitempty"`
	RateLimitRemaining int           `json:"rate_limit_remaining,omitempty"`
	Viewer             *ViewerStatus `json:"viewer,omitempty"`
	Teams              []TeamStatus  `json:"teams,omitempty"`
	Message            string        `json:"message"`
}

type StatusSummary struct {
	Mode               string `json:"mode"`
	TokenConfigured    bool   `json:"token_configured"`
	Synced             bool   `json:"synced"`
	QueuePath          string `json:"queue_path"`
	Endpoint           string `json:"endpoint,omitempty"`
	HTTPStatus         int    `json:"http_status,omitempty"`
	RateLimitRemaining int    `json:"rate_limit_remaining,omitempty"`
	ViewerConfigured   bool   `json:"viewer_configured"`
	TeamCount          int    `json:"team_count"`
	Message            string `json:"message"`
}

type ViewerStatus struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

type TeamStatus struct {
	ID   string `json:"id"`
	Key  string `json:"key,omitempty"`
	Name string `json:"name"`
}

type OfflineEvent struct {
	At      string `json:"at"`
	Kind    string `json:"kind"`
	Message string `json:"message"`
	Synced  bool   `json:"synced"`
}

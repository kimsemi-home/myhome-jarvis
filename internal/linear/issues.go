package linear

type StateStatus struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
}

type Issue struct {
	ID          string      `json:"id"`
	Identifier  string      `json:"identifier"`
	Title       string      `json:"title"`
	Description string      `json:"description,omitempty"`
	URL         string      `json:"url,omitempty"`
	UpdatedAt   string      `json:"updatedAt,omitempty"`
	Team        TeamStatus  `json:"team"`
	State       StateStatus `json:"state"`
}

type Comment struct {
	ID        string `json:"id"`
	Body      string `json:"body,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

type OperationResult struct {
	Mode               string       `json:"mode"`
	Synced             bool         `json:"synced"`
	QueuePath          string       `json:"queue_path"`
	HTTPStatus         int          `json:"http_status,omitempty"`
	RateLimitRemaining int          `json:"rate_limit_remaining,omitempty"`
	Message            string       `json:"message"`
	Issues             []Issue      `json:"issues,omitempty"`
	Issue              *Issue       `json:"issue,omitempty"`
	Comment            *Comment     `json:"comment,omitempty"`
	State              *StateStatus `json:"state,omitempty"`
}

type IssueSummary struct {
	Identifier string `json:"identifier"`
	Title      string `json:"title"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	StateType  string `json:"state_type,omitempty"`
}

type OperationSummary struct {
	Mode               string         `json:"mode"`
	Synced             bool           `json:"synced"`
	QueuePath          string         `json:"queue_path"`
	HTTPStatus         int            `json:"http_status,omitempty"`
	RateLimitRemaining int            `json:"rate_limit_remaining,omitempty"`
	Message            string         `json:"message"`
	IssueCount         int            `json:"issue_count,omitempty"`
	Issues             []IssueSummary `json:"issues,omitempty"`
	Issue              *IssueSummary  `json:"issue,omitempty"`
	CommentCreated     bool           `json:"comment_created,omitempty"`
	CommentCreatedAt   string         `json:"comment_created_at,omitempty"`
	StateType          string         `json:"state_type,omitempty"`
}

type backlogSeed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

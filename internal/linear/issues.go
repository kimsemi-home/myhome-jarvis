package linear

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
)

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

const projectIssueTitlePrefix = "[myhome-jarvis]"

type issueScope struct {
	TeamID  string
	TeamKey string
}

func PullIssues(ctx context.Context, root string, client *http.Client) OperationResult {
	result := baseOperationResult(root)
	token, err := loadToken(root)
	if err != nil {
		result.Message = "No Linear token found. Pull skipped and offline mode remains active."
		return result
	}

	var response struct {
		Issues struct {
			Nodes []Issue `json:"nodes"`
		} `json:"issues"`
	}
	query := `query PullIssues { issues(first: 50) { nodes { id identifier title updatedAt team { id key } state { id name type } } } }`
	httpStatus, remaining, err := doGraphQL(ctx, client, token.Value, query, nil, &response)
	result.HTTPStatus = httpStatus
	result.RateLimitRemaining = remaining
	if err != nil {
		result.Message = "Linear pull failed; offline mode remains active: " + err.Error()
		return result
	}
	result.Mode = "online"
	result.Synced = true
	result.Issues = filterActiveIssues(response.Issues.Nodes, configuredIssueScope())
	result.Message = fmt.Sprintf("Pulled %d active Linear issues.", len(result.Issues))
	return result
}

func NextIssue(ctx context.Context, root string, client *http.Client) OperationResult {
	result := PullIssues(ctx, root, client)
	if !result.Synced || len(result.Issues) == 0 {
		if result.Message == "" {
			result.Message = "No Linear issue is available."
		}
		return result
	}
	for index := range result.Issues {
		issue := result.Issues[index]
		if isOpenState(issue.State) && isProjectIssue(issue) {
			result.Issue = &issue
			result.Message = "Selected next project Linear issue."
			return result
		}
	}
	result.Message = "Pulled active Linear issues, but none matched the project issue prefix."
	return result
}

func AddComment(ctx context.Context, root string, client *http.Client, issueID string, body string) OperationResult {
	issueID = strings.TrimSpace(issueID)
	body = strings.TrimSpace(body)
	payload := map[string]string{"issue_id": issueID, "body": body}
	result := baseOperationResult(root)
	if issueID == "" || body == "" {
		result.Message = "Issue id and comment body are required."
		_ = AppendOfflineAction(root, "linear_comment_invalid", result.Message, payload)
		return result
	}
	token, err := loadToken(root)
	if err != nil {
		result.Message = "No Linear token found. Comment queued offline with synced=false."
		_ = AppendOfflineAction(root, "linear_comment", result.Message, payload)
		return result
	}

	var response struct {
		CommentCreate struct {
			Success bool     `json:"success"`
			Comment *Comment `json:"comment"`
		} `json:"commentCreate"`
	}
	query := `mutation AddComment($issueId: String!, $body: String!) { commentCreate(input: { issueId: $issueId, body: $body }) { success comment { id createdAt } } }`
	httpStatus, remaining, err := doGraphQL(ctx, client, token.Value, query, map[string]string{"issueId": issueID, "body": body}, &response)
	result.HTTPStatus = httpStatus
	result.RateLimitRemaining = remaining
	if err != nil || !response.CommentCreate.Success {
		result.Message = "Linear comment failed; queued offline with synced=false."
		if err != nil {
			result.Message += " " + err.Error()
		}
		_ = AppendOfflineAction(root, "linear_comment", result.Message, payload)
		return result
	}
	result.Mode = "online"
	result.Synced = true
	result.Comment = response.CommentCreate.Comment
	result.Message = "Linear comment created."
	return result
}

func TransitionIssue(ctx context.Context, root string, client *http.Client, issueID string, stateName string) OperationResult {
	issueID = strings.TrimSpace(issueID)
	stateName = strings.TrimSpace(stateName)
	payload := map[string]string{"issue_id": issueID, "state": stateName}
	result := baseOperationResult(root)
	if issueID == "" || stateName == "" {
		result.Message = "Issue id and target state are required."
		_ = AppendOfflineAction(root, "linear_transition_invalid", result.Message, payload)
		return result
	}
	token, err := loadToken(root)
	if err != nil {
		result.Message = "No Linear token found. Transition queued offline with synced=false."
		_ = AppendOfflineAction(root, "linear_transition", result.Message, payload)
		return result
	}

	stateID, status, remaining, err := findWorkflowStateID(ctx, client, token.Value, stateName)
	result.HTTPStatus = status
	result.RateLimitRemaining = remaining
	if err != nil {
		result.Message = "Linear transition lookup failed; queued offline with synced=false. " + err.Error()
		_ = AppendOfflineAction(root, "linear_transition", result.Message, payload)
		return result
	}

	var response struct {
		IssueUpdate struct {
			Success bool   `json:"success"`
			Issue   *Issue `json:"issue"`
		} `json:"issueUpdate"`
	}
	query := `mutation TransitionIssue($issueId: String!, $stateId: String!) { issueUpdate(id: $issueId, input: { stateId: $stateId }) { success issue { id identifier title state { id name type } } } }`
	httpStatus, remaining, err := doGraphQL(ctx, client, token.Value, query, map[string]string{"issueId": issueID, "stateId": stateID}, &response)
	result.HTTPStatus = httpStatus
	result.RateLimitRemaining = remaining
	if err != nil || !response.IssueUpdate.Success {
		result.Message = "Linear transition failed; queued offline with synced=false."
		if err != nil {
			result.Message += " " + err.Error()
		}
		_ = AppendOfflineAction(root, "linear_transition", result.Message, payload)
		return result
	}
	result.Mode = "online"
	result.Synced = true
	result.Issue = response.IssueUpdate.Issue
	if response.IssueUpdate.Issue != nil {
		result.State = &response.IssueUpdate.Issue.State
	}
	result.Message = "Linear issue transitioned."
	return result
}

func CreateFromBacklog(ctx context.Context, root string, client *http.Client) OperationResult {
	result := baseOperationResult(root)
	seeds := backlogSeeds()
	payload := map[string]any{"issues": seeds}
	token, err := loadToken(root)
	if err != nil {
		result.Message = "No Linear token found. Backlog seed queued offline with synced=false."
		_ = AppendOfflineAction(root, "linear_create_from_backlog", result.Message, payload)
		return result
	}
	teamID := strings.TrimSpace(os.Getenv("LINEAR_TEAM_ID"))
	if teamID == "" {
		teams, status, remaining, err := listTeams(ctx, client, token.Value)
		result.HTTPStatus = status
		result.RateLimitRemaining = remaining
		if err != nil || len(teams) == 0 {
			result.Message = "Linear team lookup failed; backlog seed queued offline with synced=false."
			if err != nil {
				result.Message += " " + err.Error()
			}
			_ = AppendOfflineAction(root, "linear_create_from_backlog", result.Message, payload)
			return result
		}
		teamID = teams[0].ID
	}

	for _, seed := range seeds {
		var response struct {
			IssueCreate struct {
				Success bool   `json:"success"`
				Issue   *Issue `json:"issue"`
			} `json:"issueCreate"`
		}
		variables := map[string]any{
			"title":       seed.Title,
			"description": seed.Description,
			"teamId":      teamID,
			"priority":    seed.Priority,
		}
		query := `mutation IssueCreate($teamId: String!, $title: String!, $description: String!, $priority: Int) { issueCreate(input: { teamId: $teamId, title: $title, description: $description, priority: $priority }) { success issue { id identifier title state { id name type } } } }`
		httpStatus, remaining, err := doGraphQL(ctx, client, token.Value, query, variables, &response)
		result.HTTPStatus = httpStatus
		result.RateLimitRemaining = remaining
		if err != nil || !response.IssueCreate.Success {
			result.Message = "Linear issue creation failed; remaining seed queued offline with synced=false."
			if err != nil {
				result.Message += " " + err.Error()
			}
			_ = AppendOfflineAction(root, "linear_create_from_backlog", result.Message, payload)
			return result
		}
		if response.IssueCreate.Issue != nil {
			result.Issues = append(result.Issues, *response.IssueCreate.Issue)
		}
	}
	result.Mode = "online"
	result.Synced = true
	result.Message = fmt.Sprintf("Created %d Linear backlog seed issues.", len(result.Issues))
	return result
}

func baseOperationResult(root string) OperationResult {
	return OperationResult{
		Mode:      "offline",
		Synced:    false,
		QueuePath: privateRelativePath(filepathJoinSlash(root, "data", "private", "linear-offline-queue.jsonl")),
	}
}

func SummarizeOperation(result OperationResult) OperationSummary {
	summary := OperationSummary{
		Mode:               result.Mode,
		Synced:             result.Synced,
		QueuePath:          privateRelativePath(result.QueuePath),
		HTTPStatus:         result.HTTPStatus,
		RateLimitRemaining: result.RateLimitRemaining,
		Message:            result.Message,
		IssueCount:         len(result.Issues),
	}
	if len(result.Issues) > 0 {
		summary.Issues = make([]IssueSummary, 0, len(result.Issues))
		for _, issue := range result.Issues {
			summary.Issues = append(summary.Issues, summarizeIssue(issue))
		}
	}
	if result.Issue != nil {
		issue := summarizeIssue(*result.Issue)
		summary.Issue = &issue
	}
	if result.Comment != nil {
		summary.CommentCreated = true
		summary.CommentCreatedAt = result.Comment.CreatedAt
	}
	if result.State != nil {
		summary.StateType = strings.TrimSpace(result.State.Type)
	}
	return summary
}

func summarizeIssue(issue Issue) IssueSummary {
	return IssueSummary{
		Identifier: strings.TrimSpace(issue.Identifier),
		Title:      strings.TrimSpace(issue.Title),
		UpdatedAt:  strings.TrimSpace(issue.UpdatedAt),
		StateType:  strings.TrimSpace(issue.State.Type),
	}
}

func findWorkflowStateID(ctx context.Context, client *http.Client, token string, wanted string) (string, int, int, error) {
	var response struct {
		WorkflowStates struct {
			Nodes []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"nodes"`
		} `json:"workflowStates"`
	}
	query := `query WorkflowStates { workflowStates { nodes { id name type } } }`
	httpStatus, remaining, err := doGraphQL(ctx, client, token, query, nil, &response)
	if err != nil {
		return "", httpStatus, remaining, err
	}
	for _, state := range response.WorkflowStates.Nodes {
		if strings.EqualFold(state.Name, wanted) || strings.EqualFold(state.Type, wanted) {
			return state.ID, httpStatus, remaining, nil
		}
	}
	return "", httpStatus, remaining, fmt.Errorf("workflow state %q not found", wanted)
}

func listTeams(ctx context.Context, client *http.Client, token string) ([]TeamStatus, int, int, error) {
	var response struct {
		Teams struct {
			Nodes []TeamStatus `json:"nodes"`
		} `json:"teams"`
	}
	query := `query Teams { teams { nodes { id name } } }`
	httpStatus, remaining, err := doGraphQL(ctx, client, token, query, nil, &response)
	return response.Teams.Nodes, httpStatus, remaining, err
}

func isOpenState(state StateStatus) bool {
	name := strings.ToLower(strings.TrimSpace(state.Name))
	stateType := strings.ToLower(strings.TrimSpace(state.Type))
	switch name {
	case "done", "canceled", "cancelled":
		return false
	}
	switch stateType {
	case "completed", "canceled", "cancelled":
		return false
	}
	return true
}

func configuredIssueScope() issueScope {
	return issueScope{
		TeamID:  strings.TrimSpace(os.Getenv("LINEAR_TEAM_ID")),
		TeamKey: strings.TrimSpace(os.Getenv("LINEAR_TEAM_KEY")),
	}
}

func filterActiveIssues(issues []Issue, scope issueScope) []Issue {
	filtered := make([]Issue, 0, len(issues))
	for _, issue := range issues {
		if !isOpenState(issue.State) {
			continue
		}
		if scope.TeamID != "" && !strings.EqualFold(strings.TrimSpace(issue.Team.ID), scope.TeamID) {
			continue
		}
		if scope.TeamKey != "" && !strings.EqualFold(strings.TrimSpace(issue.Team.Key), scope.TeamKey) {
			continue
		}
		filtered = append(filtered, issue)
	}
	return filtered
}

func isProjectIssue(issue Issue) bool {
	return strings.HasPrefix(strings.TrimSpace(issue.Title), projectIssueTitlePrefix)
}

func backlogSeeds() []backlogSeed {
	return []backlogSeed{
		{
			Title:       "[myhome-jarvis] P0: Enforce no-Python language policy",
			Description: "Acceptance: `go run ./cmd/mhj security check` rejects Python, Node.js, TypeScript, secret, and private-data risks.",
			Priority:    1,
		},
		{
			Title:       "[myhome-jarvis] P0: Add Go mhj CLI skeleton",
			Description: "Acceptance: `version`, `security check`, `command`, `harness home`, `linear status`, `linear pull`, `linear next`, `linear comment`, `linear transition`, `loop once`, and `quality` commands exist.",
			Priority:    1,
		},
		{
			Title:       "[myhome-jarvis] P0: Add Common Lisp executable SSOT",
			Description: "Acceptance: SBCL validation and deterministic codegen both pass.",
			Priority:    1,
		},
		{
			Title:       "[myhome-jarvis] P0: Implement Rust command validation core",
			Description: "Acceptance: Rust tests cover YouTube, OTT, volume, display, and unsafe URL cases.",
			Priority:    1,
		},
		{
			Title:       "[myhome-jarvis] P1: Add Linear GraphQL client in Go",
			Description: "Acceptance: status, pull, next, comment, transition, and create-from-backlog use direct GraphQL HTTP calls with offline fallback.",
			Priority:    2,
		},
	}
}

func filepathJoinSlash(parts ...string) string {
	if len(parts) == 0 {
		return ""
	}
	joined := parts[0]
	for _, part := range parts[1:] {
		joined = strings.TrimRight(joined, "/") + "/" + strings.TrimLeft(part, "/")
	}
	return joined
}

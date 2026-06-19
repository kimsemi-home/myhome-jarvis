package linear

import (
	"context"
	"fmt"
	"net/http"
)

func createBacklogIssue(ctx context.Context, client *http.Client, token string, teamID string, seed backlogSeed) (*Issue, int, int, error) {
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
	status, remaining, err := doGraphQL(ctx, client, token, query, variables, &response)
	if err != nil {
		return nil, status, remaining, err
	}
	if !response.IssueCreate.Success {
		return nil, status, remaining, fmt.Errorf("issueCreate returned success=false")
	}
	return response.IssueCreate.Issue, status, remaining, nil
}

func recordBacklogWriteEvidence(root string, issues []Issue) bool {
	evidenceRecorded := true
	for _, issue := range issues {
		if err := AppendWriteEvidence(root, "linear_create_from_backlog", issue.Identifier); err != nil {
			evidenceRecorded = false
		}
	}
	return evidenceRecorded
}

package linear

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

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
	status, remaining, err := doGraphQL(ctx, client, token, query, nil, &response)
	if err != nil {
		return "", status, remaining, err
	}
	for _, state := range response.WorkflowStates.Nodes {
		if strings.EqualFold(state.Name, wanted) || strings.EqualFold(state.Type, wanted) {
			return state.ID, status, remaining, nil
		}
	}
	return "", status, remaining, fmt.Errorf("workflow state %q not found", wanted)
}

func listTeams(ctx context.Context, client *http.Client, token string) ([]TeamStatus, int, int, error) {
	var response struct {
		Teams struct {
			Nodes []TeamStatus `json:"nodes"`
		} `json:"teams"`
	}
	query := `query Teams { teams { nodes { id name } } }`
	status, remaining, err := doGraphQL(ctx, client, token, query, nil, &response)
	return response.Teams.Nodes, status, remaining, err
}

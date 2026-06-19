package linear

import (
	"context"
	"fmt"
	"net/http"
)

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
	status, remaining, err := doGraphQL(ctx, client, token.Value, query, nil, &response)
	result.HTTPStatus = status
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

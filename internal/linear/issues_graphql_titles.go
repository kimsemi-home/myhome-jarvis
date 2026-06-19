package linear

import (
	"context"
	"net/http"
)

func listIssueTitles(ctx context.Context, client *http.Client, token string) (map[string]struct{}, int, int, error) {
	var response struct {
		Issues struct {
			Nodes []struct {
				Title string `json:"title"`
			} `json:"nodes"`
		} `json:"issues"`
	}
	query := `query ExistingIssueTitles { issues(first: 250) { nodes { title } } }`
	status, remaining, err := doGraphQL(ctx, client, token, query, nil, &response)
	if err != nil {
		return nil, status, remaining, err
	}
	titles := make(map[string]struct{}, len(response.Issues.Nodes))
	for _, node := range response.Issues.Nodes {
		title := normalizedIssueTitle(node.Title)
		if title != "" {
			titles[title] = struct{}{}
		}
	}
	return titles, status, remaining, nil
}

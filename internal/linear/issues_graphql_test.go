package linear

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

type linearGraphQLBody struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables"`
}

func linearGraphQLClient(t *testing.T, remaining string, handle func(linearGraphQLBody) string) *http.Client {
	t.Helper()
	return &http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
		var body linearGraphQLBody
		if err := json.NewDecoder(request.Body).Decode(&body); err != nil {
			t.Fatal(err)
		}
		return linearGraphQLResponse(http.StatusOK, remaining, handle(body)), nil
	})}
}

func linearGraphQLResponse(status int, remaining string, body string) *http.Response {
	header := http.Header{}
	if remaining != "" {
		header.Set("X-RateLimit-Remaining", remaining)
	}
	return &http.Response{StatusCode: status, Header: header, Body: io.NopCloser(strings.NewReader(body))}
}

func issueNode(id, identifier, title, teamID, teamKey, stateID, stateName, stateType string) string {
	return fmt.Sprintf(
		`{"id":%q,"identifier":%q,"title":%q,"updatedAt":"2026-06-14T00:00:00.000Z",`+
			`"team":{"id":%q,"key":%q},"state":{"id":%q,"name":%q,"type":%q}}`,
		id, identifier, title, teamID, teamKey, stateID, stateName, stateType,
	)
}

func issuesBody(nodes ...string) string {
	return `{"data":{"issues":{"nodes":[` + strings.Join(nodes, ",") + `]}}}`
}

func existingTitlesBody(titles ...string) string {
	nodes := make([]string, 0, len(titles))
	for _, title := range titles {
		nodes = append(nodes, fmt.Sprintf(`{"title":%q}`, title))
	}
	return issuesBody(nodes...)
}

func workflowStatesBody(id, name, stateType string) string {
	return fmt.Sprintf(
		`{"data":{"workflowStates":{"nodes":[{"id":%q,"name":%q,"type":%q}]}}}`,
		id, name, stateType,
	)
}

func issueMutationBody(field, identifier, title, stateName, stateType string) string {
	issue := issueNode("issue-id", identifier, title, "team-id", "KIM", "state-id", stateName, stateType)
	return fmt.Sprintf(`{"data":{%q:{"success":true,"issue":%s}}}`, field, issue)
}

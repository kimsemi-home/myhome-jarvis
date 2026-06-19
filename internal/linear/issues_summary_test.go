package linear

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"testing"
)

func TestOperationSummaryRedactsLinearIssueDetails(t *testing.T) {
	result := OperationResult{
		Mode:               "online",
		Synced:             true,
		QueuePath:          filepath.Join(t.TempDir(), "data", "private", "linear-offline-queue.jsonl"),
		HTTPStatus:         http.StatusOK,
		RateLimitRemaining: 42,
		Message:            "Selected next open Linear issue.",
		Issues: []Issue{{
			ID:          "issue-id",
			Identifier:  "MHJ-1",
			Title:       "Build local daemon",
			Description: "raw acceptance text",
			URL:         "https://linear.app/example/issue/MHJ-1",
			UpdatedAt:   "2026-06-14T00:00:00.000Z",
			Team:        TeamStatus{ID: "team-id", Name: "Private Team"},
			State:       StateStatus{ID: "state-id", Name: "Todo", Type: "unstarted"},
		}},
	}
	result.Issue = &result.Issues[0]

	payload, err := json.Marshal(SummarizeOperation(result))
	if err != nil {
		t.Fatal(err)
	}
	body := string(payload)
	requireBodyContains(
		t,
		body,
		`"queue_path":"data/private/linear-offline-queue.jsonl"`,
		`"identifier":"MHJ-1"`,
		`"state_type":"unstarted"`,
	)
	requireBodyOmits(
		t,
		body,
		`description`, `url`, `team`, `issue-id`, `team-id`, `state-id`,
		`Private Team`, `raw acceptance text`, `"queue_path":"/`, `"queue_path":"\\`,
	)
}

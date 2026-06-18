package linear

import (
	"net/http"
	"testing"
)

func TestSummarizeStatusRedactsViewerTeamAndAbsoluteQueuePath(t *testing.T) {
	status := Status{
		Mode:               "online",
		TokenConfigured:    true,
		TokenSource:        "file:data/private/linear-token.txt",
		Synced:             true,
		QueuePath:          "/tmp/work/data/private/linear-offline-queue.jsonl",
		Endpoint:           Endpoint,
		HTTPStatus:         http.StatusOK,
		RateLimitRemaining: 42,
		Viewer:             &ViewerStatus{ID: "viewer-id", Name: "Example User"},
		Teams:              []TeamStatus{{ID: "team-id", Name: "Home"}},
		Message:            "ok",
	}

	summary := SummarizeStatus(status)

	if summary.QueuePath != "data/private/linear-offline-queue.jsonl" {
		t.Fatalf("queue path = %q", summary.QueuePath)
	}
	if !summary.ViewerConfigured || summary.TeamCount != 1 {
		t.Fatalf("unexpected viewer/team summary: %#v", summary)
	}

	fallbackSummary := SummarizeStatus(Status{QueuePath: "/tmp/local/linear-offline-queue.jsonl"})
	if fallbackSummary.QueuePath != "linear-offline-queue.jsonl" {
		t.Fatalf("fallback queue path = %q", fallbackSummary.QueuePath)
	}
}

package linear

import (
	"context"
	"strings"
	"testing"
)

func TestAddCommentUsesVariables(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	commentBody := "Line one\nLine two with \"quotes\""
	root := t.TempDir()
	client := linearGraphQLClient(t, "", func(body linearGraphQLBody) string {
		if !strings.Contains(body.Query, "commentCreate") {
			t.Fatalf("expected commentCreate mutation, got %s", body.Query)
		}
		if body.Variables["body"] != commentBody || body.Variables["issueId"] != "MHJ-1" {
			t.Fatalf("unexpected variables: %#v", body.Variables)
		}
		return `{"data":{"commentCreate":{"success":true,"comment":` +
			`{"id":"comment-id","body":"ok","createdAt":"2026-06-14T00:00:00.000Z"}}}}`
	})

	result := AddComment(context.Background(), root, client, "MHJ-1", commentBody)
	if !result.Synced || result.Comment == nil || result.Comment.ID != "comment-id" {
		t.Fatalf("unexpected result: %#v", result)
	}
	status, err := WriteEvidenceStatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.SyncedMutationCount != 1 || status.LatestSyncedMutation == nil {
		t.Fatalf("write evidence status = %#v", status)
	}
	if status.LatestSyncedMutation.Action != "linear_comment" ||
		status.LatestSyncedMutation.IssueKey != "MHJ-1" ||
		!status.LatestSyncedMutation.Synced {
		t.Fatalf("write evidence = %#v", status.LatestSyncedMutation)
	}
}

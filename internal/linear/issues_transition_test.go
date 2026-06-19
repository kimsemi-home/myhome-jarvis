package linear

import (
	"context"
	"strings"
	"testing"
)

func TestTransitionRecordsApprovedWriteEvidence(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	root := t.TempDir()
	requests := 0
	client := linearGraphQLClient(t, "", func(body linearGraphQLBody) string {
		requests++
		switch {
		case strings.Contains(body.Query, "query WorkflowStates"):
			return workflowStatesBody("state-id", "Done", "completed")
		case strings.Contains(body.Query, "mutation TransitionIssue"):
			if body.Variables["issueId"] != "MHJ-1" || body.Variables["stateId"] != "state-id" {
				t.Fatalf("unexpected variables: %#v", body.Variables)
			}
			return issueMutationBody("issueUpdate", "MHJ-1", "Done issue", "Done", "completed")
		default:
			t.Fatalf("unexpected GraphQL request: %s", body.Query)
			return ""
		}
	})

	result := TransitionIssue(context.Background(), root, client, "MHJ-1", "Done")
	if !result.Synced || result.State == nil || result.State.Type != "completed" {
		t.Fatalf("unexpected result: %#v", result)
	}
	if requests != 2 {
		t.Fatalf("requests = %d", requests)
	}
	status, err := WriteEvidenceStatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.SyncedMutationCount != 1 || status.LatestSyncedMutation == nil {
		t.Fatalf("write evidence status = %#v", status)
	}
	if status.LatestSyncedMutation.Action != "linear_transition" ||
		status.LatestSyncedMutation.IssueKey != "MHJ-1" {
		t.Fatalf("write evidence = %#v", status.LatestSyncedMutation)
	}
}

package linear

import (
	"context"
	"strings"
	"testing"
)

func TestCreateFromBacklogSkipsExistingSeedTitles(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	t.Setenv("LINEAR_TEAM_ID", "team-id")
	root := t.TempDir()
	seeds := backlogSeeds()
	missingTitle := seeds[1].Title
	requests := 0
	createdTitles := []string{}
	client := linearGraphQLClient(t, "4000", func(body linearGraphQLBody) string {
		requests++
		switch {
		case strings.Contains(body.Query, "query ExistingIssueTitles"):
			return existingTitlesBody(seeds[0].Title, seeds[2].Title)
		case strings.Contains(body.Query, "mutation IssueCreate"):
			title, ok := body.Variables["title"].(string)
			if !ok || title != missingTitle {
				t.Fatalf("created title = %#v, expected %q", body.Variables["title"], missingTitle)
			}
			createdTitles = append(createdTitles, title)
			if body.Variables["teamId"] != "team-id" || body.Variables["priority"] != float64(3) {
				t.Fatalf("unexpected variables: %#v", body.Variables)
			}
			return issueMutationBody("issueCreate", "KIM-11", missingTitle, "Todo", "unstarted")
		default:
			t.Fatalf("unexpected GraphQL request: %s", body.Query)
			return ""
		}
	})

	result := CreateFromBacklog(context.Background(), root, client)
	if !result.Synced || len(result.Issues) != 1 {
		t.Fatalf("unexpected result: %#v", result)
	}
	if requests != 2 || len(createdTitles) != 1 {
		t.Fatalf("requests=%d created=%#v", requests, createdTitles)
	}
	if !strings.Contains(result.Message, "Created 1") || !strings.Contains(result.Message, "skipped 2") {
		t.Fatalf("message did not include created/skipped counts: %q", result.Message)
	}
	status, err := WriteEvidenceStatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.SyncedMutationCount != 1 || status.LatestSyncedMutation == nil {
		t.Fatalf("write evidence status = %#v", status)
	}
	if status.LatestSyncedMutation.Action != "linear_create_from_backlog" || status.LatestSyncedMutation.IssueKey != "KIM-11" {
		t.Fatalf("write evidence = %#v", status.LatestSyncedMutation)
	}
}

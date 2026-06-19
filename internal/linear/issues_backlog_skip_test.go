package linear

import (
	"context"
	"strings"
	"testing"
)

func TestCreateFromBacklogSkipsAllExistingSeedTitles(t *testing.T) {
	t.Setenv("LINEAR_API_KEY", "linear-example-key")
	t.Setenv("LINEAR_TEAM_ID", "team-id")
	seeds := backlogSeeds()
	client := linearGraphQLClient(t, "4000", func(body linearGraphQLBody) string {
		if !strings.Contains(body.Query, "query ExistingIssueTitles") {
			t.Fatalf("unexpected GraphQL request: %s", body.Query)
		}
		return existingTitlesBody(seeds[0].Title, seeds[1].Title, seeds[2].Title)
	})

	result := CreateFromBacklog(context.Background(), t.TempDir(), client)
	if !result.Synced || len(result.Issues) != 0 {
		t.Fatalf("unexpected result: %#v", result)
	}
	if result.RateLimitRemaining != 4000 {
		t.Fatalf("rate remaining = %d", result.RateLimitRemaining)
	}
	if result.Message != "Created 0 Linear backlog seed issues; skipped 3 existing seeds." {
		t.Fatalf("message = %q", result.Message)
	}
}

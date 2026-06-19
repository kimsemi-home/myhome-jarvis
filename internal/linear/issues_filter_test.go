package linear

import "testing"

func TestFilterActiveIssuesFiltersConfiguredTeamID(t *testing.T) {
	issues := []Issue{
		{
			Identifier: "KIM-6",
			Team:       TeamStatus{ID: "team-kim", Key: "KIM"},
			State:      StateStatus{Name: "Todo", Type: "unstarted"},
		},
		{
			Identifier: "OPS-1",
			Team:       TeamStatus{ID: "team-ops", Key: "OPS"},
			State:      StateStatus{Name: "Todo", Type: "unstarted"},
		},
	}

	filtered := filterActiveIssues(issues, issueScope{TeamID: "team-kim"})
	if len(filtered) != 1 || filtered[0].Identifier != "KIM-6" {
		t.Fatalf("filtered issues = %#v", filtered)
	}
}

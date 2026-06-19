package linear

import (
	"os"
	"strings"
)

const projectIssueTitlePrefix = "[myhome-jarvis]"

type issueScope struct {
	TeamID  string
	TeamKey string
}

func isOpenState(state StateStatus) bool {
	name := strings.ToLower(strings.TrimSpace(state.Name))
	stateType := strings.ToLower(strings.TrimSpace(state.Type))
	switch name {
	case "done", "canceled", "cancelled":
		return false
	}
	switch stateType {
	case "completed", "canceled", "cancelled":
		return false
	}
	return true
}

func isStartedState(state StateStatus) bool {
	return strings.EqualFold(strings.TrimSpace(state.Type), "started")
}

func configuredIssueScope() issueScope {
	return issueScope{
		TeamID:  strings.TrimSpace(envLinearTeamID()),
		TeamKey: strings.TrimSpace(os.Getenv("LINEAR_TEAM_KEY")),
	}
}

func filterActiveIssues(issues []Issue, scope issueScope) []Issue {
	filtered := make([]Issue, 0, len(issues))
	for _, issue := range issues {
		if !isOpenState(issue.State) {
			continue
		}
		if scope.TeamID != "" && !strings.EqualFold(strings.TrimSpace(issue.Team.ID), scope.TeamID) {
			continue
		}
		if scope.TeamKey != "" && !strings.EqualFold(strings.TrimSpace(issue.Team.Key), scope.TeamKey) {
			continue
		}
		filtered = append(filtered, issue)
	}
	return filtered
}

func isProjectIssue(issue Issue) bool {
	return strings.HasPrefix(strings.TrimSpace(issue.Title), projectIssueTitlePrefix)
}

func normalizedIssueTitle(title string) string {
	return strings.ToLower(strings.TrimSpace(title))
}

func envLinearTeamID() string {
	return os.Getenv("LINEAR_TEAM_ID")
}

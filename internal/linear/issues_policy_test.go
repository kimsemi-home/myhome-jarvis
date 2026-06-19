package linear

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestProjectIssueTitlePrefixMatchesGeneratedPolicy(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("..", "..", "generated", "linear.generated.json"))
	if err != nil {
		t.Fatal(err)
	}
	var policy struct {
		ProjectIssueTitlePrefix    string   `json:"project_issue_title_prefix"`
		NextPrefersProjectIssue    bool     `json:"next_prefers_project_issues"`
		NextRequiresProjectIssue   bool     `json:"next_requires_project_issue"`
		BacklogSeedDedupesByTitle  bool     `json:"backlog_seed_dedupes_by_title"`
		BacklogSeedCurrentProject  bool     `json:"backlog_seed_current_project_only"`
		BacklogSeedQueriesExisting bool     `json:"backlog_seed_queries_existing_titles"`
		OfflineReplayEvidence      string   `json:"offline_replay_evidence"`
		OfflineReplayRateFloor     int      `json:"offline_replay_rate_limit_floor"`
		Commands                   []string `json:"commands"`
		OfflineReplaySafeKinds     []string `json:"offline_replay_safe_action_kinds"`
	}
	if err := json.Unmarshal(data, &policy); err != nil {
		t.Fatal(err)
	}
	if policy.ProjectIssueTitlePrefix != projectIssueTitlePrefix {
		t.Fatalf("project prefix = %q, expected %q", policy.ProjectIssueTitlePrefix, projectIssueTitlePrefix)
	}
	if !policy.NextPrefersProjectIssue || !policy.NextRequiresProjectIssue {
		t.Fatal("generated policy must keep next project issue selection enabled")
	}
	if !policy.BacklogSeedDedupesByTitle || !policy.BacklogSeedCurrentProject || !policy.BacklogSeedQueriesExisting {
		t.Fatal("generated policy must keep backlog seed dedupe rules enabled")
	}
	if policy.OfflineReplayEvidence != OfflineReplayRelativePath {
		t.Fatalf("offline replay evidence path = %q", policy.OfflineReplayEvidence)
	}
	if policy.OfflineReplayRateFloor != defaultReplayRateLimitFloor {
		t.Fatalf("offline replay rate floor = %d", policy.OfflineReplayRateFloor)
	}
	if !containsString(policy.Commands, "mhj linear replay-offline") {
		t.Fatalf("generated commands missing replay-offline: %#v", policy.Commands)
	}
	for _, kind := range []string{offlineReplayCommentKind, offlineReplayTransitionKind} {
		if !containsString(policy.OfflineReplaySafeKinds, kind) {
			t.Fatalf("generated safe replay kinds missing %s: %#v", kind, policy.OfflineReplaySafeKinds)
		}
	}
}

package mediareadiness

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func testPolicy() Policy {
	return Policy{
		Context:                 "MediaReadinessBenchmark",
		Version:                 "v1",
		GeneratedArtifact:       PolicyRelativePath,
		BenchmarkKind:           "command_planning",
		PublicStatusRedacted:    true,
		TargetPlanningLatencyMS: 250,
		Cases: []BenchmarkCase{
			{ID: "youtube_launch", Capability: "youtube_launch", Command: "open_youtube", PayloadKind: "empty", ExpectedHost: "www.youtube.com"},
			{ID: "youtube_search", Capability: "youtube_search", Command: "open_youtube_search", PayloadKind: "fixture_query", ExpectedHost: "www.youtube.com"},
			{ID: "ott_netflix", Capability: "ott_launch", Command: "open_ott", PayloadKind: "fixture_service_netflix", ExpectedHost: "www.netflix.com"},
			{ID: "playback_readiness", Capability: "playback_readiness", Command: "open_youtube", PayloadKind: "empty", ExpectedHost: "www.youtube.com"},
		},
		Commands: []string{"mhj media-readiness status"},
	}
}

func writePolicy(t *testing.T, root string, policy Policy) {
	t.Helper()
	body, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, filepath.FromSlash(PolicyRelativePath))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, append(body, '\n'), 0o600); err != nil {
		t.Fatal(err)
	}
}

package externalevidence

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCollectWritesPrivateLayersAndCachesEvidence(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"items":[{"full_name":"owner/repo","stars":10}]}`))
	}))
	defer server.Close()
	root := t.TempDir()
	writePolicy(t, root, fixturePolicy(server.URL))

	first, err := collectForRoot(root, 0, server.Client())
	if err != nil {
		t.Fatal(err)
	}
	second, err := collectForRoot(root, 0, server.Client())
	if err != nil {
		t.Fatal(err)
	}
	if first.CollectedCount != 1 || second.CachedCount != 1 ||
		second.CollectedCount != 0 {
		t.Fatalf("collect cache reports = %#v %#v", first, second)
	}
	body := readPrivate(t, root, "data/private/external-evidence/manifest.jsonl")
	if strings.Count(body, `"kind":"external_signal"`) != 1 ||
		!strings.Contains(body, `"source":"github_fixture"`) {
		t.Fatalf("manifest body = %s", body)
	}
	entries, err := os.ReadDir(filepath.Join(root, filepath.FromSlash(first.RawLayerPath)))
	if err != nil || len(entries) != 1 {
		t.Fatalf("raw layer entries = %d err = %v", len(entries), err)
	}
}

func TestCollectReportDoesNotExposeRawPayload(t *testing.T) {
	report := CollectReport{PublicSafe: true, Results: []CollectResult{{
		SourceKey: "github_fixture", State: "collected", RawSHA256: strings.Repeat("a", 64),
	}}}
	body, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"owner/repo", "https://", "token", "cookie"} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("collect report leaked %s in %s", forbidden, body)
		}
	}
}

func readPrivate(t *testing.T, root string, rel string) string {
	t.Helper()
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		t.Fatal(err)
	}
	return string(body)
}

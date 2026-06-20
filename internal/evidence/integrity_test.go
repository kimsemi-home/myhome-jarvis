package evidence

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestIntegrityCountsDanglingRefsByPrefix(t *testing.T) {
	root := t.TempDir()
	writeEvidencePolicy(t, root, true)
	writeFile(t, root, "data/private/learning/observations.jsonl", `{"id":"learn_1","at":"2026-06-18T00:00:00Z","status":"closed","evidence_refs":["data/private/quality/runs.jsonl","generated/missing.generated.json","internal/missing.go"]}`+"\n")
	writeFile(t, root, "data/private/quality/runs.jsonl", `{"ok":true}`+"\n")

	status, err := IntegrityForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.CheckedEvidenceRefCount != 3 || status.PresentEvidenceRefCount != 1 || status.DanglingEvidenceRefCount != 2 {
		t.Fatalf("integrity counts = %#v", status)
	}
	if status.NextSafeAction != "repair_private_learning_refs" {
		t.Fatalf("next safe action = %q", status.NextSafeAction)
	}
	assertPrefixCount(t, status, "generated/", 1)
	assertPrefixCount(t, status, "internal/", 1)
}

func TestIntegrityStatusRedactsExactEvidenceRefs(t *testing.T) {
	root := t.TempDir()
	writeEvidencePolicy(t, root, true)
	writeFile(t, root, "data/private/learning/observations.jsonl", `{"id":"learn_2","at":"2026-06-18T00:00:00Z","status":"open","summary":"private summary","evidence_refs":["generated/missing.generated.json"],"next_action":"private next action"}`+"\n")

	status, err := IntegrityForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"generated/missing.generated.json", "private summary", "private next action", "evidence_refs"} {
		if strings.Contains(string(payload), forbidden) {
			t.Fatalf("integrity leaked %q in %s", forbidden, payload)
		}
	}
}

func assertPrefixCount(t *testing.T, status IntegrityStatus, prefix string, dangling int) {
	t.Helper()
	for _, item := range status.PrefixCounts {
		if item.Prefix == prefix && item.DanglingCount == dangling {
			return
		}
	}
	t.Fatalf("missing prefix %s with dangling=%d: %#v", prefix, dangling, status.PrefixCounts)
}

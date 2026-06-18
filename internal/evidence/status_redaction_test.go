package evidence

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestStatusCountsDanglingRefsWithoutLeakingRawObservation(t *testing.T) {
	root := t.TempDir()
	writeEvidencePolicy(t, root, true)
	writeFile(t, root, "data/private/learning/observations.jsonl", `{"id":"learn_2","at":"2026-06-18T00:00:00Z","kind":"evidence_debt","source":"review","stage":"evidence_recorded","status":"open","summary":"private summary should stay private","evidence_refs":["generated/missing.generated.json"],"owner":"go","next_action":"private next action"}`+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.DanglingEvidenceRefCount != 1 || status.EdgeCount != 1 {
		t.Fatalf("dangling/edge counts = %d/%d", status.DanglingEvidenceRefCount, status.EdgeCount)
	}
	payload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	assertEvidenceStatusRedacted(t, string(payload))
}

func assertEvidenceStatusRedacted(t *testing.T, body string) {
	t.Helper()
	for _, forbidden := range []string{
		"private summary",
		"private next action",
		"evidence_refs",
		"generated/missing.generated.json",
		"summary",
		"next_action",
	} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("status leaked %q in %s", forbidden, body)
		}
	}
}

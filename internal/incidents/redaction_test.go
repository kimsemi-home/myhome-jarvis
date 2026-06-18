package incidents

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func TestStatusJSONDoesNotLeakRawIncidentDetails(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/incidents/incidents.jsonl",
		`{"id":"inc_1","at":"2026-06-18T00:00:00Z","kind":"quality_regression","stage":"fix_planned","status":"mitigating","owner_role":"go","quarantine_state":"none","evidence_refs":["data/private/quality/runs.jsonl"],"summary":"private incident summary","root_cause_notes":"private root cause"}`+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.ToLower(string(payload))
	for _, forbidden := range forbiddenIncidentFragments() {
		if strings.Contains(body, forbidden) {
			t.Fatalf("incident status leaked %q in %s", forbidden, body)
		}
	}
}

func forbiddenIncidentFragments() []string {
	return []string{
		"summary", "root_cause", "private incident", "private root cause",
		"evidence_refs", "data/private/quality/runs.jsonl", "token", "secret",
		string(filepath.Separator) + "users" + string(filepath.Separator),
	}
}

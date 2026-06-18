package evidencequality

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func TestStatusJSONDoesNotLeakRawSnapshotDetails(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/evidence-quality/snapshots.jsonl",
		`{"id":"eq_1","at":"2026-06-18T00:00:00Z","evidence_ref":"data/private/quality/runs.jsonl","purpose":"release_gate","quality_level":"medium","schema_version":"quality:v1","ontology_version":"concepts:v1","mapping_confidence":"medium","assessed_by":"go","reassessment_reasons":["schema_version_change"],"raw_notes":"private quality notes","raw_evidence":"private evidence body"}`+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.ToLower(string(payload))
	for _, forbidden := range forbiddenSnapshotFragments() {
		if strings.Contains(body, forbidden) {
			t.Fatalf("evidence quality status leaked %q in %s", forbidden, body)
		}
	}
}

func forbiddenSnapshotFragments() []string {
	return []string{
		"raw_notes", "raw_evidence", "evidence_ref", "private quality",
		"private evidence", "data/private/quality/runs.jsonl", "token", "secret",
		string(filepath.Separator) + "users" + string(filepath.Separator),
	}
}

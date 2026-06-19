package controlplane

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func TestStatusJSONDoesNotLeakRawControlPlaneDetails(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/control-plane/manifests.jsonl", redactionFixture())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.InvalidManifestCount != 1 || status.ManifestDebtCount != 1 {
		t.Fatalf("expected raw rationale to count as debt, got %#v", status)
	}
	payload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.ToLower(string(payload))
	for _, forbidden := range []string{
		"raw_rationale",
		"selection_rationale",
		"candidate_agents",
		"evidence_refs",
		"private reasoning",
		"data/private/checkpoints/one.json",
		"token",
		"secret",
		"linear.app/",
		string(filepath.Separator) + "users" + string(filepath.Separator),
	} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("control-plane status leaked %q in %s", forbidden, body)
		}
	}
}

func redactionFixture() string {
	return `{"id":"cpm_1","at":"2026-06-18T00:00:00Z","decision_kind":"loop_once","policy_version":"control-plane:v1","ontology_version":"concepts:v1","authority_profile":"local_readonly","selected_route":"loop_once","reviewer_role":"go_review_gate","verifier_role":"deterministic_quality_gate","lease_seconds":120,"lease_status":"finished","evidence_refs":["generated/control_plane.generated.json"],"output_ref":"data/private/checkpoints/one.json","raw_rationale":"private reasoning"}` + "\n"
}

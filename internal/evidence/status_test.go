package evidence

import (
	"testing"
)

func TestStatusConnectsLearningObservationsToEvidenceArtifacts(t *testing.T) {
	root := t.TempDir()
	writeEvidencePolicy(t, root, true)
	writeFile(t, root, "data/private/learning/observations.jsonl", `{"id":"learn_1","at":"2026-06-18T00:00:00Z","kind":"loop_gap","source":"quality_gate","stage":"evidence_recorded","status":"open","summary":"private detail","evidence_refs":["data/private/quality/runs.jsonl","docs/working-log.md"],"owner":"go","next_action":"private action"}`+"\n")
	writeFile(t, root, "data/private/quality/runs.jsonl", `{"at":"2026-06-18T00:01:00Z","ok":true}`+"\n")
	writeFile(t, root, "docs/working-log.md", "# Working Log\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.SourceCount != 9 || status.PresentSourceCount != 2 {
		t.Fatalf("source counts = %d/%d", status.PresentSourceCount, status.SourceCount)
	}
	if status.NodeCount != 4 || status.EdgeCount != 2 {
		t.Fatalf("graph counts nodes=%d edges=%d", status.NodeCount, status.EdgeCount)
	}
	if status.OpenLearningCount != 1 || status.DanglingEvidenceRefCount != 0 {
		t.Fatalf("learning/dangling counts = %d/%d", status.OpenLearningCount, status.DanglingEvidenceRefCount)
	}
	if status.ByNodeKind["learning_observation"] != 1 || status.ByNodeKind["quality_run"] != 1 || status.ByNodeKind["evidence_artifact"] != 2 {
		t.Fatalf("node kinds = %#v", status.ByNodeKind)
	}
	if status.ByEdgeKind["supports"] != 2 {
		t.Fatalf("edge kinds = %#v", status.ByEdgeKind)
	}
	if status.LastObservedAt != "2026-06-18T00:01:00Z" {
		t.Fatalf("last observed = %q", status.LastObservedAt)
	}
}

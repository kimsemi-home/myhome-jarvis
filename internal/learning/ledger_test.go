package learning

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStatusForRootMissingJournal(t *testing.T) {
	root := copyPolicyFixture(t)
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.Exists || status.Count != 0 || status.OpenCount != 0 {
		t.Fatalf("status = %#v", status)
	}
	if status.Path != "data/private/learning/observations.jsonl" {
		t.Fatalf("path = %q", status.Path)
	}
	if status.PolicyPath != PolicyRelativePath {
		t.Fatalf("policy path = %q", status.PolicyPath)
	}
}

func TestRecordWritesPrivateObservationAndRedactedStatus(t *testing.T) {
	root := copyPolicyFixture(t)
	payload := []byte(`{"kind":"loop_gap","source":"quality_gate","summary":"Quality failed before a regression was captured.","evidence_refs":["data/private/quality/runs.jsonl","docs/working-log.md"],"owner":"go","next_action":"Add a focused regression test."}`)

	result, err := Record(root, payload)
	if err != nil {
		t.Fatal(err)
	}
	if result.Kind != "loop_gap" || result.Stage != "evidence_recorded" || result.Status != "open" || result.EvidenceRefCount != 2 {
		t.Fatalf("result = %#v", result)
	}
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.Exists || status.Count != 1 || status.OpenCount != 1 || status.ByKind["loop_gap"] != 1 || status.ByStage["evidence_recorded"] != 1 {
		t.Fatalf("status = %#v", status)
	}
	statusPayload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"Quality failed", "Add a focused", "evidence_refs"} {
		if strings.Contains(string(statusPayload), forbidden) {
			t.Fatalf("status leaked %q in %s", forbidden, statusPayload)
		}
	}
	journal, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(status.Path)))
	if err != nil {
		t.Fatal(err)
	}
	for _, expected := range []string{`"summary":"Quality failed before a regression was captured."`, `"next_action":"Add a focused regression test."`} {
		if !strings.Contains(string(journal), expected) {
			t.Fatalf("journal missing %s in %s", expected, journal)
		}
	}
	if strings.Contains(string(journal), root) {
		t.Fatalf("journal leaked root in %s", journal)
	}
}

func TestRecordRejectsMissingEvidence(t *testing.T) {
	root := copyPolicyFixture(t)
	payload := []byte(`{"kind":"loop_gap","source":"quality_gate","summary":"gap","owner":"go","next_action":"add test"}`)

	if _, err := Record(root, payload); err == nil {
		t.Fatal("expected missing evidence refs to fail")
	}
}

func TestRecordRejectsSensitiveText(t *testing.T) {
	root := copyPolicyFixture(t)
	payload := []byte(`{"kind":"loop_gap","source":"quality_gate","summary":"Bearer abc123 appeared in output.","evidence_refs":["data/private/quality/runs.jsonl"],"owner":"go","next_action":"add test"}`)

	if _, err := Record(root, payload); err == nil {
		t.Fatal("expected sensitive text to fail")
	}
}

func TestRecordRejectsAbsoluteEvidenceRef(t *testing.T) {
	root := copyPolicyFixture(t)
	payload := []byte(`{"kind":"loop_gap","source":"quality_gate","summary":"gap","evidence_refs":["/tmp/evidence.jsonl"],"owner":"go","next_action":"add test"}`)

	if _, err := Record(root, payload); err == nil {
		t.Fatal("expected absolute evidence ref to fail")
	}
}

func copyPolicyFixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	source := filepath.Join(repoRoot(t), filepath.FromSlash(PolicyRelativePath))
	body, err := os.ReadFile(source)
	if err != nil {
		t.Fatal(err)
	}
	target := filepath.Join(root, filepath.FromSlash(PolicyRelativePath))
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(target, body, 0o644); err != nil {
		t.Fatal(err)
	}
	return root
}

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		next := filepath.Dir(dir)
		if next == dir {
			t.Fatal("could not find repo root")
		}
		dir = next
	}
}

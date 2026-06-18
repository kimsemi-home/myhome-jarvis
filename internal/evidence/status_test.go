package evidence

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
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
	body := string(payload)
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

func TestReadPolicyRejectsRawPublicEvidence(t *testing.T) {
	root := t.TempDir()
	writeEvidencePolicy(t, root, false)

	_, err := StatusForRoot(root)
	if err == nil {
		t.Fatal("expected raw public evidence policy to fail")
	}
}

func writeEvidencePolicy(t *testing.T, root string, redacted bool) {
	t.Helper()
	rawPublic := "false"
	if !redacted {
		rawPublic = "true"
	}
	writeFile(t, root, PolicyRelativePath, `{
  "context": "AgentCluster",
  "version": "v1",
  "generated_artifact": "generated/evidence.generated.json",
  "private_root": "data/private",
  "private_graph_required": true,
  "public_status_redacted": true,
  "raw_evidence_public_allowed": `+rawPublic+`,
  "node_kinds": ["learning_observation", "evidence_artifact", "checkpoint", "control_plane_manifest", "incident", "evidence_quality_snapshot", "review_queue_item", "quality_run", "linear_write", "audit_event"],
  "edge_kinds": ["supports", "observed_in", "verified_by", "recorded_by"],
  "private_sources": [
    {"key": "learning", "path": "data/private/learning/observations.jsonl", "node_kind": "learning_observation", "format": "jsonl"},
    {"key": "checkpoints", "path": "data/private/checkpoints", "node_kind": "checkpoint", "format": "directory"},
    {"key": "control_plane", "path": "data/private/control-plane/manifests.jsonl", "node_kind": "control_plane_manifest", "format": "jsonl"},
    {"key": "incidents", "path": "data/private/incidents/incidents.jsonl", "node_kind": "incident", "format": "jsonl"},
    {"key": "evidence_quality", "path": "data/private/evidence-quality/snapshots.jsonl", "node_kind": "evidence_quality_snapshot", "format": "jsonl"},
    {"key": "review", "path": "data/private/review/queue.jsonl", "node_kind": "review_queue_item", "format": "jsonl"},
    {"key": "quality", "path": "data/private/quality/runs.jsonl", "node_kind": "quality_run", "format": "jsonl"},
    {"key": "linear_write", "path": "data/private/linear-write-evidence.jsonl", "node_kind": "linear_write", "format": "jsonl"},
    {"key": "audit", "path": "data/private/audit/command-intents.jsonl", "node_kind": "audit_event", "format": "jsonl"}
  ],
  "allowed_evidence_prefixes": ["data/private/", "generated/", "docs/", "cmd/", "internal/", "apps/flutter/", "lisp/", "crates/", "fixtures/", "harness/", ".github/"],
  "public_summary_fields": ["policy_path", "private_root", "source_count", "present_source_count", "node_count", "edge_count", "dangling_evidence_ref_count", "open_learning_count", "by_node_kind", "by_edge_kind", "sources", "last_observed_at", "checked_at"],
  "forbidden_public_fields": ["summary", "next_action", "evidence_refs", "token", "secret", "credential", "cookie", "raw_prompt", "raw_transcript", "account_id", "card_number", "local_absolute_path"],
  "commands": ["mhj evidence status"]
}`+"\n")
}

func writeFile(t *testing.T, root string, rel string, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}

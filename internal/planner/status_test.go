package planner

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
)

func TestStatusForRootReturnsGeneratedPlannerGraph(t *testing.T) {
	status, err := StatusForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}

	if status.LoopMode != "closed-loop" {
		t.Fatalf("loop mode = %q", status.LoopMode)
	}
	if status.TaskCount != 6 {
		t.Fatalf("task count = %d", status.TaskCount)
	}
	if status.ReadyCount != 0 {
		t.Fatalf("ready count = %d", status.ReadyCount)
	}
	if status.CompletedCount != 5 {
		t.Fatalf("completed count = %d", status.CompletedCount)
	}
	if status.BlockedExternalWriteCount != 1 {
		t.Fatalf("blocked external write count = %d", status.BlockedExternalWriteCount)
	}
	if !status.ExternalWriteGate.StandingBoundary || !status.ExternalWriteGate.ApprovalRequired || !status.ExternalWriteGate.MutationSuccessRequired {
		t.Fatalf("external write gate = %#v", status.ExternalWriteGate)
	}
	if status.ExternalWriteGate.BoundaryTaskID != "linear_sync" || !status.ExternalWriteGate.BoundaryTaskBlocked {
		t.Fatalf("external write gate boundary = %#v", status.ExternalWriteGate)
	}
	if status.LinearWriteEvidence.EvidencePath != linear.WriteEvidenceRelativePath {
		t.Fatalf("linear write evidence path = %q", status.LinearWriteEvidence.EvidencePath)
	}
	if len(status.BlockedExternalWriteTasks) != 1 || status.BlockedExternalWriteTasks[0].ID != "linear_sync" {
		t.Fatalf("blocked external write tasks = %#v", status.BlockedExternalWriteTasks)
	}
	if status.NextTask != nil {
		t.Fatalf("next task = %#v", status.NextTask)
	}
	if status.LinearTemplateCount != 2 {
		t.Fatalf("linear template count = %d", status.LinearTemplateCount)
	}
	if status.CheckpointRoot != "data/private/checkpoints" {
		t.Fatalf("checkpoint root = %q", status.CheckpointRoot)
	}
	if !status.QualityRequired {
		t.Fatal("quality should be required")
	}
	if !status.LinearOfflineFallback {
		t.Fatal("linear offline fallback should be enabled")
	}
	if !status.KnowledgeIndexRequired {
		t.Fatal("knowledge index should be required before planning")
	}
	if status.KnowledgeEvidence == nil {
		t.Fatal("expected knowledge evidence")
	}
	if status.KnowledgeEvidence.Query != "planner KnowledgeIndex Linear closed loop" {
		t.Fatalf("knowledge query = %q", status.KnowledgeEvidence.Query)
	}
	if status.KnowledgeEvidence.HitCount == 0 || len(status.KnowledgeEvidence.MustRead) == 0 {
		t.Fatalf("knowledge evidence = %#v", status.KnowledgeEvidence)
	}
}

func TestStatusForRootSeparatesExternalWriteGateFromSyncedEvidence(t *testing.T) {
	root := t.TempDir()
	generatedDir := filepath.Join(root, "generated")
	if err := os.MkdirAll(generatedDir, 0o755); err != nil {
		t.Fatal(err)
	}
	data := `{"loop_mode":"closed-loop","max_task_scope":"one file","checkpoint_root":"data/private/checkpoints","quality_required":true,"linear_offline_fallback":true,"knowledge_index_required_before_planning":false,"knowledge_index_default_query":"","external_write_gate":{"standing_boundary":true,"approval_required":true,"mutation_success_required":true,"boundary_task_id":"linear_sync","evidence_path":"data/private/linear-write-evidence.jsonl"},"default_next":"ready","task_graph":[{"id":"repo_safety","title":"Repo safety","owner":"go","status":"completed","depends_on":[]},{"id":"linear_sync","title":"Linear sync","owner":"go","status":"blocked_external_write","depends_on":["repo_safety"]}],"linear_templates":[]}`
	if err := os.WriteFile(filepath.Join(generatedDir, "planner.generated.json"), []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := linear.AppendWriteEvidence(root, "linear_transition", "KIM-12"); err != nil {
		t.Fatal(err)
	}

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}

	if !status.ExternalWriteGate.BoundaryTaskBlocked || status.BlockedExternalWriteCount != 1 {
		t.Fatalf("external write gate not preserved: %#v", status)
	}
	if status.LinearWriteEvidence.SyncedMutationCount != 1 || !status.LinearWriteEvidence.HasSyncedMutation {
		t.Fatalf("linear write evidence status = %#v", status.LinearWriteEvidence)
	}
	if status.LinearWriteEvidence.LatestSyncedMutation == nil {
		t.Fatalf("latest write evidence missing: %#v", status.LinearWriteEvidence)
	}
	if status.LinearWriteEvidence.LatestSyncedMutation.Action != "linear_transition" || status.LinearWriteEvidence.LatestSyncedMutation.IssueKey != "KIM-12" {
		t.Fatalf("latest write evidence = %#v", status.LinearWriteEvidence.LatestSyncedMutation)
	}
}

func TestReadPolicyRejectsUnknownTaskStatus(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "planner.json")
	data := `{"loop_mode":"closed-loop","max_task_scope":"one file","checkpoint_root":"data/private/checkpoints","quality_required":true,"linear_offline_fallback":true,"external_write_gate":{"standing_boundary":true,"approval_required":true,"mutation_success_required":true,"boundary_task_id":"linear_sync","evidence_path":"data/private/linear-write-evidence.jsonl"},"default_next":"ready","task_graph":[{"id":"repo_safety","title":"Repo safety","owner":"go","status":"maybe","depends_on":[]},{"id":"linear_sync","title":"Linear sync","owner":"go","status":"blocked_external_write","depends_on":["repo_safety"]}],"linear_templates":[]}`
	if err := os.WriteFile(path, []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}
	policy, err := ReadPolicy(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected unknown task status to be rejected")
	}
}

func TestReadPolicyRejectsAbsoluteCheckpointRoot(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "planner.json")
	data := `{"loop_mode":"closed-loop","max_task_scope":"one file","checkpoint_root":"/tmp/checkpoints","quality_required":true,"linear_offline_fallback":true,"external_write_gate":{"standing_boundary":true,"approval_required":true,"mutation_success_required":true,"boundary_task_id":"linear_sync","evidence_path":"data/private/linear-write-evidence.jsonl"},"default_next":"ready","task_graph":[{"id":"linear_sync","title":"Linear sync","owner":"go","status":"blocked_external_write","depends_on":[]}],"linear_templates":[]}`
	if err := os.WriteFile(path, []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}
	policy, err := ReadPolicy(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected absolute checkpoint root to be rejected")
	}
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

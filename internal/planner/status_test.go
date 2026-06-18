package planner

import (
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

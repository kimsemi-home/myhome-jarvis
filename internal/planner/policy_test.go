package planner

import (
	"path/filepath"
	"testing"
)

func TestReadPolicyRejectsUnknownTaskStatus(t *testing.T) {
	path := filepath.Join(t.TempDir(), "planner.json")
	writePlannerPolicyFixture(t, path, "data/private/checkpoints", "maybe")

	policy, err := ReadPolicy(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected unknown task status to be rejected")
	}
}

func TestReadPolicyRejectsAbsoluteCheckpointRoot(t *testing.T) {
	path := filepath.Join(t.TempDir(), "planner.json")
	writePlannerPolicyFixture(t, path, "/tmp/checkpoints", "completed")

	policy, err := ReadPolicy(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := validatePolicy(policy); err == nil {
		t.Fatal("expected absolute checkpoint root to be rejected")
	}
}

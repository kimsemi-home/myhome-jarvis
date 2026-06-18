package orchestrator

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteCheckpointUsesCollisionResistantNames(t *testing.T) {
	root := t.TempDir()
	first, err := WriteCheckpoint(root, Checkpoint{Task: "first"})
	if err != nil {
		t.Fatal(err)
	}
	second, err := WriteCheckpoint(root, Checkpoint{Task: "second"})
	if err != nil {
		t.Fatal(err)
	}
	if first == second {
		t.Fatalf("checkpoint paths collided: %s", first)
	}
	for _, path := range []string{first, second} {
		if !strings.Contains(filepath.Base(path), ".") {
			t.Fatalf("checkpoint filename should include sub-second precision: %s", path)
		}
	}
}

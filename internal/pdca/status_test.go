package pdca

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStatusForRootCountsCyclesWithoutLeakingRefs(t *testing.T) {
	root := t.TempDir()
	writePolicyFixture(t, root)
	if err := os.MkdirAll(filepath.Join(root, "generated"), 0o755); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"planner", "control_plane", "verification_evidence", "learning"} {
		path := filepath.Join(root, "generated", name+".generated.json")
		if err := os.WriteFile(path, []byte("{}"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	ledger := filepath.Join(root, "data", "private", "pdca")
	if err := os.MkdirAll(ledger, 0o755); err != nil {
		t.Fatal(err)
	}
	body := `{"cycle_id":"c1","at":"2026-06-19T00:00:00Z","status":"closed","owner":"go","plan_ref":"generated/planner.generated.json","do_ref":"data/private/checkpoints/c1.json","check_ref":"generated/verification_evidence.generated.json","act_ref":"data/private/learning/observations.jsonl"}` + "\n"
	if err := os.WriteFile(filepath.Join(ledger, "cycles.jsonl"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.Ready || status.CycleCount != 1 || status.ClosedCount != 1 {
		t.Fatalf("unexpected status: %#v", status)
	}
}

func writePolicyFixture(t *testing.T, root string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Join(root, "generated"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, PolicyRelativePath), []byte(policyFixture), 0o644); err != nil {
		t.Fatal(err)
	}
}

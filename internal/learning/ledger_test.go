package learning

import "testing"

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

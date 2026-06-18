package review

import "testing"

func TestStatusForRootAllowsMissingQueue(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.Exists || status.Count != 0 || status.ReviewDebtCount != 0 {
		t.Fatalf("status = %#v", status)
	}
	if status.CapacityState != "available" || status.ActiveRule != "no_open_reviews" {
		t.Fatalf("capacity = %s/%s", status.CapacityState, status.ActiveRule)
	}
}

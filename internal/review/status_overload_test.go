package review

import "testing"

func TestStatusForRootFreezesHighRiskReviewOverload(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/review/queue.jsonl", highRiskReviewJSON+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.CapacityState != "overloaded" || status.ActiveRule != "high_risk_review_overload" {
		t.Fatalf("capacity = %#v", status)
	}
	if status.ReviewDebtCount != 1 || status.HighRiskOpenCount != 1 || status.BackupAvailableCount != 1 {
		t.Fatalf("counts = %#v", status)
	}
	if status.ByQueueClass["major_ontology_change"] != 1 || status.ByRisk["high"] != 1 {
		t.Fatalf("buckets = %#v/%#v", status.ByQueueClass, status.ByRisk)
	}
}

package mediareadiness

import "testing"

func TestStatusMeasuresPlanningCases(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.CaseCount != 3 || status.AvailableCount != 3 || status.DegradedCount != 0 {
		t.Fatalf("unexpected media readiness counts: %#v", status)
	}
	if status.MaxPlanningLatencyMS > status.TargetPlanningLatencyMS {
		t.Fatalf("planning latency exceeded target: %#v", status)
	}
}

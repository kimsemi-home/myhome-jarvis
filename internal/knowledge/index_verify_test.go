package knowledge

import "testing"

func TestVerifyGeneratedRegistry(t *testing.T) {
	report, err := Verify(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	if !report.OK {
		t.Fatalf("verify failed: %#v", report)
	}
	if report.ContextCount != 9 {
		t.Fatalf("context count = %d", report.ContextCount)
	}
	if report.ConceptCount != 30 {
		t.Fatalf("concept count = %d", report.ConceptCount)
	}
	if report.EventCount != 2 {
		t.Fatalf("event count = %d", report.EventCount)
	}
	if report.HarnessCount != 3 {
		t.Fatalf("harness count = %d", report.HarnessCount)
	}
}

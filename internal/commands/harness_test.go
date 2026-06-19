package commands

import "testing"

func TestRunHomeHarness(t *testing.T) {
	report := RunHomeHarness()
	if !report.Passed {
		t.Fatalf("home harness failed: %+v", report.Results)
	}
}

func TestRunFinanceHarness(t *testing.T) {
	report := RunFinanceHarness(repoRoot(t))
	if !report.Passed {
		t.Fatalf("finance harness failed: %+v", report.Results)
	}
	if report.Name != "finance" {
		t.Fatalf("harness name = %q", report.Name)
	}
	if len(report.Results) < 8 {
		t.Fatalf("finance harness result count = %d", len(report.Results))
	}
}

func TestRunCommerceHarness(t *testing.T) {
	report := RunCommerceHarness(repoRoot(t))
	if !report.Passed {
		t.Fatalf("commerce harness failed: %+v", report.Results)
	}
	if report.Name != "commerce" {
		t.Fatalf("harness name = %q", report.Name)
	}
	if len(report.Results) < 7 {
		t.Fatalf("commerce harness result count = %d", len(report.Results))
	}
}

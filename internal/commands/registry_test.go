package commands

import "testing"

func TestBuildVolumeSetRejectsOutOfRange(t *testing.T) {
	if _, err := Build("volume-set", []byte(`{"level":101}`)); err == nil {
		t.Fatal("expected out-of-range level to fail")
	}
}

func TestBuildOpenURLRejectsUnsafeScheme(t *testing.T) {
	if _, err := Build("open-url", []byte(`{"url":"javascript:alert(1)"}`)); err == nil {
		t.Fatal("expected unsafe URL to fail")
	}
}

func TestRunHomeHarness(t *testing.T) {
	report := RunHomeHarness()
	if !report.Passed {
		t.Fatalf("home harness failed: %+v", report.Results)
	}
}

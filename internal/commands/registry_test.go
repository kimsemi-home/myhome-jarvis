package commands

import (
	"os"
	"path/filepath"
	"testing"
)

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

func repoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		next := filepath.Dir(dir)
		if next == dir {
			t.Fatal("could not find repo root")
		}
		dir = next
	}
}

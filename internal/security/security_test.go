package security

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckRejectsPythonFile(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "script.py"), []byte("print('no')\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	report, err := Check(root)
	if err != nil {
		t.Fatal(err)
	}
	if report.OK {
		t.Fatal("expected Python file to be rejected")
	}
}

func TestCheckReportDoesNotExposeLocalRoot(t *testing.T) {
	root := t.TempDir()
	report, err := Check(root)
	if err != nil {
		t.Fatal(err)
	}
	if report.Root != "." {
		t.Fatalf("root = %q, expected redacted current root", report.Root)
	}
	if strings.Contains(report.Root, root) || filepath.IsAbs(report.Root) {
		t.Fatalf("report leaked local root: %#v", report)
	}
}

func TestCheckAllowsPrivateLocalFiles(t *testing.T) {
	root := t.TempDir()
	privateDir := filepath.Join(root, "data", "private")
	if err := os.MkdirAll(privateDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(privateDir, "linear-token.txt"), []byte("local\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	report, err := Check(root)
	if err != nil {
		t.Fatal(err)
	}
	if !report.OK {
		t.Fatalf("expected private local files to be allowed, got %+v", report.Findings)
	}
}

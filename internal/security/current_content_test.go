package security

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckRejectsPrivateMarkerInCurrentContent(t *testing.T) {
	root := t.TempDir()
	privatePath := "/" + "Users" + "/" + strings.Join([]string{"al", "ice"}, "") + "/project"
	if err := os.WriteFile(filepath.Join(root, "notes.md"), []byte("old path: "+privatePath+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	report, err := Check(root)
	if err != nil {
		t.Fatal(err)
	}
	if report.OK {
		t.Fatal("expected current private marker to be rejected")
	}
	if !currentHasCode(report, "current_private_identity") {
		t.Fatalf("expected private identity finding, got %+v", report.Findings)
	}
	if strings.Contains(report.Findings[0].Message, privatePath) {
		t.Fatalf("finding leaked matched content: %+v", report.Findings[0])
	}
}

func TestCheckRejectsSecretLiteralInCurrentContent(t *testing.T) {
	root := t.TempDir()
	key := strings.Join([]string{"api", "_", "key"}, "")
	value := strings.Repeat("a", 24)
	if err := os.WriteFile(filepath.Join(root, "config.md"), []byte(key+" = "+value+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	report, err := Check(root)
	if err != nil {
		t.Fatal(err)
	}
	if report.OK {
		t.Fatal("expected current secret-looking literal to be rejected")
	}
	if !currentHasCode(report, "current_secret_literal") {
		t.Fatalf("expected secret literal finding, got %+v", report.Findings)
	}
	for _, finding := range report.Findings {
		if strings.Contains(finding.Message, value) {
			t.Fatalf("finding leaked matched secret content: %+v", finding)
		}
	}
}

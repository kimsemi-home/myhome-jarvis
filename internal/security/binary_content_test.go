package security

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckSkipsBinaryFontContent(t *testing.T) {
	root := t.TempDir()
	privatePath := "/" + "Users" + "/" + strings.Join([]string{"al", "ice"}, "") + "/font"
	if err := os.WriteFile(filepath.Join(root, "font.ttf"), []byte(privatePath), 0o600); err != nil {
		t.Fatal(err)
	}
	report, err := Check(root)
	if err != nil {
		t.Fatal(err)
	}
	if !report.OK {
		t.Fatalf("expected binary font content to be skipped, got %+v", report.Findings)
	}
}

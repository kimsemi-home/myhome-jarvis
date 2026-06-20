package contextpack

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRepoContextPackDeclarationVerifies(t *testing.T) {
	root := repoRoot(t)
	result, err := VerifyDeclarationForRoot(root, "")
	if err != nil {
		t.Fatal(err)
	}
	if !result.Valid || result.DeclarationPath != ".mhj/context-pack.json" {
		t.Fatalf("repo declaration result = %#v", result)
	}
	body, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(body), "/"+"Users"+"/") {
		t.Fatalf("verify result leaked local path: %s", body)
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
			t.Fatal("could not locate repo root")
		}
		dir = next
	}
}

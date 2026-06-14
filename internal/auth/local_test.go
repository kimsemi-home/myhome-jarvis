package auth

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateStatusReadAndRotateLocalToken(t *testing.T) {
	root := t.TempDir()
	initial := Status(root)
	if initial.Configured {
		t.Fatalf("unexpected initial status: %#v", initial)
	}

	created, err := Create(root, false)
	if err != nil {
		t.Fatal(err)
	}
	if !created.Configured || created.Token == "" || created.Path != "data/private/local-token.txt" {
		t.Fatalf("unexpected create result: %#v", created)
	}
	info, err := os.Stat(filepath.Join(root, "data", "private", "local-token.txt"))
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0o600 {
		t.Fatalf("mode = %v", info.Mode().Perm())
	}
	read, err := Read(root)
	if err != nil {
		t.Fatal(err)
	}
	if read != created.Token {
		t.Fatalf("read token mismatch")
	}

	existing, err := Create(root, false)
	if err != nil {
		t.Fatal(err)
	}
	if existing.Token != "" || existing.Message == "" {
		t.Fatalf("unexpected existing result: %#v", existing)
	}

	rotated, err := Create(root, true)
	if err != nil {
		t.Fatal(err)
	}
	if !rotated.Rotated || rotated.Token == created.Token {
		t.Fatalf("unexpected rotate result: %#v", rotated)
	}
}

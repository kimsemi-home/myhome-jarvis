package storagearchive

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func assertGzipContains(t *testing.T, root string, rel string, want string) {
	t.Helper()
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	reader, err := gzip.NewReader(file)
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()
	content, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(content), want) {
		t.Fatalf("gzip archive missing %q in %s", want, content)
	}
}

func readManifestEntry(t *testing.T, root string, rel string) manifestEntry {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		t.Fatal(err)
	}
	var entry manifestEntry
	if err := json.Unmarshal([]byte(strings.Split(strings.TrimSpace(string(data)), "\n")[0]), &entry); err != nil {
		t.Fatal(err)
	}
	return entry
}

func mustJSON(t *testing.T, value any) string {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

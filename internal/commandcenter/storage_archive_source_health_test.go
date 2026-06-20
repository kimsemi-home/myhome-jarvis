package commandcenter

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestStorageArchiveSourceHealthIsPublicSafe(t *testing.T) {
	status, err := StatusForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	storage := status.StorageArchive
	if len(storage.SourceHealth) != storage.PrivateLogSourceCount {
		t.Fatalf("source health count = %#v", storage)
	}
	if !hasStorageSource(storage.SourceHealth, "quality") ||
		!hasStorageSource(storage.SourceHealth, "authority_review") {
		t.Fatalf("source health keys = %#v", storage.SourceHealth)
	}
	body, err := json.Marshal(storage)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range [][]byte{
		[]byte("source_path"),
		[]byte("archive_path"),
		[]byte("input_sha256"),
		[]byte("data/private/quality"),
		[]byte("linear" + ".app"),
	} {
		if bytes.Contains(body, forbidden) {
			t.Fatalf("storage source health leaked %s in %s", forbidden, body)
		}
	}
}

func hasStorageSource(sources []StorageArchiveSourceHealth, key string) bool {
	for _, source := range sources {
		if source.SourceKey == key {
			return true
		}
	}
	return false
}

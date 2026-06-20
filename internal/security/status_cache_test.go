package security

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestStatusForRootCachesHistoryAggregate(t *testing.T) {
	root := cleanCommittedRepo(t)

	first, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if first.Cache.State != "miss" || !first.Cache.PublicSafe {
		t.Fatalf("expected public-safe cache miss, got %#v", first.Cache)
	}
	second, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if second.Cache.State != "hit" || second.Cache.Key != first.Cache.Key {
		t.Fatalf("expected cache hit with same key, got %#v then %#v", first.Cache, second.Cache)
	}
	if strings.Contains(second.Cache.Path, root) || second.Cache.RawDetailsPublicAllowed {
		t.Fatalf("cache evidence leaked private details: %#v", second.Cache)
	}
}

func TestStatusCacheMissesWhenHeadChanges(t *testing.T) {
	root := cleanCommittedRepo(t)
	first, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "CHANGELOG.md"), []byte("ok\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	commitAll(t, root, "second")
	second, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if second.Cache.State != "miss" || second.Cache.Key == first.Cache.Key {
		t.Fatalf("expected stale cache miss after HEAD change, got %#v then %#v", first.Cache, second.Cache)
	}
}

func TestStatusCacheKeepsCurrentScanFresh(t *testing.T) {
	root := cleanCommittedRepo(t)
	if _, err := StatusForRoot(root); err != nil {
		t.Fatal(err)
	}
	if _, err := StatusForRoot(root); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "script.py"), []byte("print('no')\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.Cache.State != "hit" || status.CurrentOK || status.OK {
		t.Fatalf("expected cached history with fresh current failure, got %#v", status)
	}
}

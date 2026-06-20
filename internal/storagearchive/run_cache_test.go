package storagearchive

import "testing"

func TestRunUsesHashCacheForUnchangedSource(t *testing.T) {
	root := t.TempDir()
	source := privateQualitySource()
	writeStoragePolicy(t, root, testStoragePolicy(source))
	writePrivateFile(t, root, source.Path,
		`{"source":"quality","kind":"run","evidence_ref":"a"}`+"\n")
	first, err := RunForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	second, err := RunForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if first.ArchivedCount != 1 || second.CachedCount != 1 {
		t.Fatalf("cache run counts = %#v / %#v", first, second)
	}
	if second.ArchivedCount != 0 || second.SkippedCount != 0 {
		t.Fatalf("unchanged input should be cache-only: %#v", second)
	}
	if second.Results[0].State != "cached" ||
		second.Results[0].ArchivePath != first.Results[0].ArchivePath {
		t.Fatalf("cache result = %#v", second.Results[0])
	}
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.ManifestEntryCount != 1 || status.ManifestArchivedCount != 1 {
		t.Fatalf("manifest should not grow on cache hit: %#v", status)
	}
}

func TestRunArchivesChangedInputAfterHashCacheHit(t *testing.T) {
	root := t.TempDir()
	source := privateQualitySource()
	writeStoragePolicy(t, root, testStoragePolicy(source))
	writePrivateFile(t, root, source.Path,
		`{"source":"quality","kind":"run","evidence_ref":"a"}`+"\n")
	first, err := RunForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	writePrivateFile(t, root, source.Path,
		`{"source":"quality","kind":"run","evidence_ref":"b"}`+"\n")
	second, err := RunForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if second.ArchivedCount != 1 || second.CachedCount != 0 {
		t.Fatalf("changed input should archive again: %#v", second)
	}
	if second.Results[0].ArchivePath == first.Results[0].ArchivePath {
		t.Fatalf("changed input reused archive path: %#v", second.Results[0])
	}
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.ManifestEntryCount != 2 || status.ManifestArchivedCount != 2 {
		t.Fatalf("changed input manifest summary = %#v", status)
	}
}

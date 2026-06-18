package main

import (
	"reflect"
	"testing"
)

func TestChangedGeneratedFilesReportsAddedModifiedAndDeleted(t *testing.T) {
	before := map[string][]byte{
		"generated/a.json": []byte("same"),
		"generated/b.json": []byte("old"),
		"generated/c.json": []byte("removed"),
	}
	after := map[string][]byte{
		"generated/a.json": []byte("same"),
		"generated/b.json": []byte("new"),
		"generated/d.json": []byte("added"),
	}

	changed := changedGeneratedFiles(before, after)
	expected := []string{
		"generated/b.json",
		"generated/c.json",
		"generated/d.json",
	}
	if !reflect.DeepEqual(changed, expected) {
		t.Fatalf("changed files = %#v, expected %#v", changed, expected)
	}
}

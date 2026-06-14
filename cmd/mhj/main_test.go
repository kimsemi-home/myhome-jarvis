package main

import (
	"encoding/json"
	"reflect"
	"strings"
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

func TestQualityReportJSONRedactsCommandAndOutput(t *testing.T) {
	report := qualityReport{
		OK: true,
		Steps: []qualityStep{{
			Name:    "flutter test",
			Status:  "pass",
			Command: []string{"/private/toolchains/flutter", "test"},
			Output:  "loading /private/checkout/apps/flutter/test/widget_test.dart",
		}},
	}

	payload, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	body := string(payload)
	for _, expected := range []string{`"ok":true`, `"name":"flutter test"`, `"status":"pass"`} {
		if !strings.Contains(body, expected) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{`command`, `output`, `/private/toolchains`, `/private/checkout`} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("quality report leaked %s in %s", forbidden, body)
		}
	}
}

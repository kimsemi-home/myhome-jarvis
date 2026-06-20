package codexsustainability

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestCaptureQualityRunRedactsPrivatePayload(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeQualityRun(t, root, true)
	capture, err := captureQualityRunAt(root, mustTime(t, "2026-06-20T01:00:00Z"))
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(capture)
	if err != nil {
		t.Fatal(err)
	}
	assertNoCaptureLeak(t, string(body))
	ledger, err := os.ReadFile(filepath.Join(root, "data/private/codex-sustainability/evidence.jsonl"))
	if err != nil {
		t.Fatal(err)
	}
	assertNoCaptureLeak(t, string(ledger))
}

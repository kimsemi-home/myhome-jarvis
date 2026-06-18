package orchestrator

import (
	"os"
	"strings"
	"testing"
)

func TestWriteCheckpointStoresAggregateSecurityStatus(t *testing.T) {
	root := t.TempDir()
	path, err := WriteCheckpoint(root, redactedCheckpointFixture())
	if err != nil {
		t.Fatal(err)
	}
	body, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	text := string(body)
	if !strings.Contains(text, `"security_status"`) {
		t.Fatalf("expected aggregate security status in %s", text)
	}
	if !strings.Contains(text, `"planner_status"`) || !strings.Contains(text, `"blocked_external_write_count": 1`) {
		t.Fatalf("expected planner status in %s", text)
	}
	if !strings.Contains(text, `"linear_next"`) || !strings.Contains(text, `"identifier": "KIM-13"`) {
		t.Fatalf("expected linear next evidence in %s", text)
	}
	if !strings.Contains(text, `"knowledge_evidence"`) || !strings.Contains(text, `"KIM-14"`) {
		t.Fatalf("expected knowledge evidence in %s", text)
	}
	for _, forbidden := range []string{`"security_report"`, `"findings"`, `"root"`, `"viewer"`, `"teams"`} {
		if strings.Contains(text, forbidden) {
			t.Fatalf("checkpoint leaked %s in %s", forbidden, text)
		}
	}
}

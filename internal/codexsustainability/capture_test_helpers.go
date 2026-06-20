package codexsustainability

import (
	"strings"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
)

func writeQualityRun(t *testing.T, root string, ok bool) {
	t.Helper()
	run := qualitylog.Run{
		At:             "2026-06-20T00:00:00Z",
		OK:             ok,
		DurationMillis: 360001,
		StepCount:      1,
		PassCount:      1,
		Steps:          []qualitylog.Step{{Name: "go test", Status: "pass"}},
	}
	if !ok {
		run.PassCount = 0
		run.FailCount = 1
		run.Steps[0].Status = "fail"
	}
	if err := qualitylog.AppendRun(root, run); err != nil {
		t.Fatal(err)
	}
}

func assertNoCaptureLeak(t *testing.T, body string) {
	t.Helper()
	for _, forbidden := range []string{
		"raw_prompt", "raw_transcript", "private_notes", "command",
		"output", "local_absolute_path", "linear_url", "token",
		"credential", "account_id", "finance_payload",
	} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("capture leaked %q in %s", forbidden, body)
		}
	}
}

package confidence

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kimsemi-home/myhome-jarvis/internal/evidence"
)

func TestStatusJSONDoesNotLeakRawEvidence(t *testing.T) {
	status := Assess(testPolicy(), clearInputs(evidence.Status{
		EdgeCount:                1,
		DanglingEvidenceRefCount: 1,
	}))
	payload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}

	body := string(payload)
	for _, forbidden := range forbiddenStatusFragments() {
		if strings.Contains(strings.ToLower(body), forbidden) {
			t.Fatalf("status leaked %q in %s", forbidden, body)
		}
	}
}

func forbiddenStatusFragments() []string {
	return []string{
		"summary",
		"next_action",
		"evidence_refs",
		"raw_prompt",
		"raw_transcript",
		"token",
		"secret",
		string(filepath.Separator) + "users" + string(filepath.Separator),
	}
}

package commandcenter

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestVisionAuditDoesNotExposePrivatePayloadFields(t *testing.T) {
	audit, err := VisionAuditForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(audit)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{
		"raw_prompt", "raw_transcript", "private_notes",
		"token", "secret", "credential", "local_absolute_path",
		"linear_url", "finance_payload", "/" + "Users" + "/",
	} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("vision audit leaked %q in %s", forbidden, body)
		}
	}
}

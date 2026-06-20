package commandcenter

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestStatusDoesNotExposePrivatePayloadFields(t *testing.T) {
	status, err := StatusForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{
		"raw_prompt", "raw_transcript", "private_notes",
		"token", "secret", "credential", "local_absolute_path",
		"reviewer_identity", "linear_url", "finance_payload",
		"evidence_refs", "ledger_path", "attribution_ledger_path",
		"subject_key", "account_id", "card_number",
	} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("command center leaked %q in %s", forbidden, body)
		}
	}
}

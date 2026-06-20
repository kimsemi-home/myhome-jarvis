package cicache

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestStatusForRootIsPublicSafe(t *testing.T) {
	root := t.TempDir()
	writeFixture(t, root, workflowFixture())
	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	body, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{
		"token", "secret", "credential", "cookie", "raw_prompt",
		"raw_transcript", "linear" + ".app", "private_evidence",
		"/Use" + "rs/",
	} {
		if strings.Contains(string(body), forbidden) {
			t.Fatalf("status leaked %q in %s", forbidden, body)
		}
	}
}

package authority

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func TestStatusForRootReturnsRedactedAuthoritySummary(t *testing.T) {
	status, err := StatusForRoot(repoRoot(t))
	if err != nil {
		t.Fatal(err)
	}
	payload, err := json.Marshal(status)
	if err != nil {
		t.Fatal(err)
	}
	body := strings.ToLower(string(payload))
	for _, expected := range []string{
		`"policy_path":"generated/authority.generated.json"`,
		`"reasoning_tier_grants_approval":false`,
		`"self_authority_allowed":false`,
	} {
		if !strings.Contains(body, expected) {
			t.Fatalf("expected %s in %s", expected, body)
		}
	}
	for _, forbidden := range []string{
		"raw_rationale",
		"raw_evidence",
		"evidence_ref",
		"evidence_refs",
		"raw_prompt",
		"raw_transcript",
		"token",
		"secret",
		"credential",
		"linear.app",
		string(filepath.Separator) + "users" + string(filepath.Separator),
	} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("authority status leaked %q in %s", forbidden, body)
		}
	}
}

func TestReadPolicyRejectsReasoningTierApproval(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	policy.ReasoningTierGrantsApproval = true
	writePolicy(t, root, policy)

	_, err := ReadPolicy(root)
	if err == nil {
		t.Fatal("expected reasoning tier approval policy to fail")
	}
}

package controlplane

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestVerifyForRootReturnsPublicSafeSummary(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, verifierPolicy())
	writeVerifier(t, root, verifierManifest())

	status, err := VerifyForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.OK || status.CheckCount != 5 || status.ManifestDebtCount != 0 {
		t.Fatalf("status = %#v", status)
	}
}

func TestVerifyForRootFailsOnManifestDebt(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, verifierPolicy())
	writeVerifier(t, root, verifierManifest())
	writeFile(t, root, "data/private/control-plane/manifests.jsonl", "{\n")

	_, err := VerifyForRoot(root)
	if err == nil || !strings.Contains(err.Error(), "manifest debt") {
		t.Fatalf("expected manifest debt failure, got %v", err)
	}
}

func verifierPolicy() Policy {
	policy := testPolicy()
	policy.VerifierGeneratedArtifact = VerificationRelativePath
	policy.VerificationCommand = "mhj control-plane verify"
	policy.VerifierChecks = []string{
		"policy-json-valid", "status-public-redacted", "lease-bounds-valid",
		"verifier-separation-required", "manifest-debt-evaluated",
	}
	return policy
}

func verifierManifest() VerificationPolicy {
	return VerificationPolicy{
		SchemaVersion:              "control-plane.verification/v1",
		Source:                     "lisp/ssot/control-plane.lisp",
		PolicyArtifact:             PolicyRelativePath,
		Command:                    "mhj control-plane verify",
		StatusCommand:              "mhj control-plane status",
		VerifierSeparationRequired: true,
		Checks: []VerificationCheck{
			{ID: "policy-json-valid", Evidence: "policy"},
			{ID: "status-public-redacted", Evidence: "status"},
			{ID: "lease-bounds-valid", Evidence: "lease"},
			{ID: "verifier-separation-required", Evidence: "roles"},
			{ID: "manifest-debt-evaluated", Evidence: "debt"},
		},
	}
}

func writeVerifier(t *testing.T, root string, policy VerificationPolicy) {
	t.Helper()
	body, err := json.Marshal(policy)
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, root, VerificationRelativePath, string(body)+"\n")
}

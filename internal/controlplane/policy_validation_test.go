package controlplane

import "testing"

func TestReadPolicyRejectsRawPublicRationale(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	policy.RawRationalePublicAllowed = true
	writePolicy(t, root, policy)

	_, err := ReadPolicy(root)
	if err == nil {
		t.Fatal("expected raw rationale policy to fail")
	}
}

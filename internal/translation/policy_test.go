package translation

import "testing"

func TestReadPolicyRejectsRawPublicLosses(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	policy.RawLossPublicAllowed = true
	writePolicy(t, root, policy)

	_, err := ReadPolicy(root)
	if err == nil {
		t.Fatal("expected raw public loss policy to fail")
	}
}

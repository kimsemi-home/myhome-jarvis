package incidents

import "testing"

func TestReadPolicyRejectsRawPublicIncidentDetails(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	policy.RawIncidentPublicAllowed = true
	writePolicy(t, root, policy)

	_, err := ReadPolicy(root)
	if err == nil {
		t.Fatal("expected raw incident public policy to fail")
	}
}

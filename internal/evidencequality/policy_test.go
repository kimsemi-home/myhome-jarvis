package evidencequality

import "testing"

func TestReadPolicyRejectsRawPublicSnapshots(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	policy.RawSnapshotPublicAllowed = true
	writePolicy(t, root, policy)

	_, err := ReadPolicy(root)
	if err == nil {
		t.Fatal("expected raw public snapshot policy to fail")
	}
}

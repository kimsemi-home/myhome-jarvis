package confidence

import "testing"

func TestReadPolicyRejectsSelfReporting(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, true)

	_, err := ReadPolicy(root)
	if err == nil {
		t.Fatal("expected self-reporting policy to fail")
	}
}

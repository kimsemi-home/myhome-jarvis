package codeshape

import "testing"

func TestReadPolicyRejectsAbsoluteLegacyPath(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy([]LegacyDebtFile{{Path: "/tmp/a.go", MaxLines: 3}}))

	_, err := ReadPolicy(root)
	if err == nil {
		t.Fatal("expected absolute legacy path to fail")
	}
}

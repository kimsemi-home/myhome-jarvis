package evidence

import "testing"

func TestReadPolicyRejectsRawPublicEvidence(t *testing.T) {
	root := t.TempDir()
	writeEvidencePolicy(t, root, false)

	_, err := StatusForRoot(root)
	if err == nil {
		t.Fatal("expected raw public evidence policy to fail")
	}
}

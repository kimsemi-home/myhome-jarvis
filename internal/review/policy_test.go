package review

import "testing"

func TestReadPolicyRejectsRawPublicReview(t *testing.T) {
	root := t.TempDir()
	policy := testPolicy()
	policy.RawReviewPublicAllowed = true
	writePolicy(t, root, policy)

	_, err := ReadPolicy(root)
	if err == nil {
		t.Fatal("expected raw public review policy to fail")
	}
}

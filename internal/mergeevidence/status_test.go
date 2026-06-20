package mergeevidence

import "testing"

func TestStatusSummarizesMergeEvidencePolicy(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.PublicSafe || !status.MergeReady {
		t.Fatalf("expected merge-ready public-safe policy: %#v", status)
	}
	if status.MissingGateCount != 0 || status.MissingRequiredEvidenceCount != 0 {
		t.Fatalf("missing policy coverage: %#v", status)
	}
	if !status.PostMergeEvidenceRequired || !status.LinearCompletionRequired ||
		!status.MainQualityRunRequired || !status.PrivateDataScanRequired {
		t.Fatalf("expected merge-first evidence proof flags: %#v", status)
	}
	if status.EligibleGateCount != 5 || status.RequiredEvidenceCount != 11 {
		t.Fatalf("unexpected merge evidence counts: %#v", status)
	}
}

func TestStatusBlocksWhenRequiredEvidenceMissing(t *testing.T) {
	policy := testPolicy()
	policy.RequiredEvidence = policy.RequiredEvidence[:2]
	status := statusFromPolicy(policy)

	if status.MergeReady || !status.MergeBlockedUntilEvidence {
		t.Fatalf("expected blocked status: %#v", status)
	}
}

func TestStatusBlocksWhenMergeFirstProofFlagsMissing(t *testing.T) {
	policy := testPolicy()
	policy.PrivateDataScanRequired = false
	status := statusFromPolicy(policy)

	if status.MergeReady || status.PublicSafe {
		t.Fatalf("expected private-data scan proof to block readiness: %#v", status)
	}
}

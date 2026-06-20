package cicache

import "testing"

func TestStatusForRootReportsHealthyUnitCacheEvidence(t *testing.T) {
	root := t.TempDir()
	writeFixture(t, root, workflowFixture())

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.OK || status.CachedUnitCount != 2 || status.UncachedUnitCount != 1 {
		t.Fatalf("status = %#v", status)
	}
	if !status.PublicSafetyNonSkippable || !status.GeneratedCoverageOK {
		t.Fatalf("contract flags = %#v", status)
	}
	if status.Units[1].GeneratedCoverageCount != 2 {
		t.Fatalf("ssot generated coverage = %#v", status.Units[1])
	}
}

func TestStatusForRootDetectsBrokenWorkflowContract(t *testing.T) {
	root := t.TempDir()
	workflow := workflowFixture()
	workflow = replaceAll(workflow, "steps.unit-cache.outputs.cache-hit != 'true'", "always()")
	writeFixture(t, root, workflow)

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.OK || status.WorkflowContractIssueCount == 0 {
		t.Fatalf("expected workflow issue, got %#v", status)
	}
}

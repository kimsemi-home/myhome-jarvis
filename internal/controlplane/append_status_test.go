package controlplane

import "testing"

func TestAppendManifestAndStatusSummarizeRedactedCounts(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	result, err := AppendManifest(root, ManifestRequest{
		DecisionKind:     "loop_once",
		PolicyVersion:    "control-plane:v1",
		OntologyVersion:  "concepts:v1",
		AuthorityProfile: "local_readonly",
		SelectedRoute:    "loop_once",
		ReviewerRole:     "go_review_gate",
		VerifierRole:     "deterministic_quality_gate",
		LeaseSeconds:     120,
		LeaseStatus:      "finished",
		EvidenceRefs:     []string{"generated/control_plane.generated.json", "generated/planner.generated.json"},
		OutputRef:        "data/private/checkpoints/20260618T000000.000000000Z.json",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.ManifestPath != "data/private/control-plane/manifests.jsonl" || result.DecisionKind != "loop_once" {
		t.Fatalf("result = %#v", result)
	}

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.Exists || status.Count != 1 || status.ManifestDebtCount != 0 {
		t.Fatalf("status = %#v", status)
	}
	if status.ByDecisionKind["loop_once"] != 1 || status.ByAuthorityProfile["local_readonly"] != 1 || status.ByLeaseStatus["finished"] != 1 {
		t.Fatalf("status maps = %#v %#v %#v", status.ByDecisionKind, status.ByAuthorityProfile, status.ByLeaseStatus)
	}
}

func TestStatusCountsMalformedManifestDebt(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	writeFile(t, root, "data/private/control-plane/manifests.jsonl", `{`+"\n")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if status.InvalidManifestCount != 1 || status.ManifestDebtCount != 1 {
		t.Fatalf("status = %#v", status)
	}
}

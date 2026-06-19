package controlplane

import "testing"

func TestAppendManifestRejectsVerifierSelfApproval(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())

	_, err := AppendManifest(root, ManifestRequest{
		DecisionKind:     "loop_once",
		PolicyVersion:    "control-plane:v1",
		OntologyVersion:  "concepts:v1",
		AuthorityProfile: "local_readonly",
		SelectedRoute:    "loop_once",
		ReviewerRole:     "same_role",
		VerifierRole:     "same_role",
		LeaseSeconds:     120,
		LeaseStatus:      "finished",
		EvidenceRefs:     []string{"generated/control_plane.generated.json"},
		OutputRef:        "data/private/checkpoints/one.json",
	})
	if err == nil {
		t.Fatal("expected verifier self-approval to fail")
	}
}

func TestAppendManifestRejectsSensitiveEvidenceRef(t *testing.T) {
	root := t.TempDir()
	writePolicy(t, root, testPolicy())
	sensitiveRef := "docs/" + "linear.app/" + "kim/private-note.md"

	_, err := AppendManifest(root, ManifestRequest{
		DecisionKind:     "loop_once",
		PolicyVersion:    "control-plane:v1",
		OntologyVersion:  "concepts:v1",
		AuthorityProfile: "local_readonly",
		SelectedRoute:    "loop_once",
		ReviewerRole:     "go_review_gate",
		VerifierRole:     "deterministic_quality_gate",
		LeaseSeconds:     120,
		LeaseStatus:      "finished",
		EvidenceRefs:     []string{sensitiveRef},
		OutputRef:        "data/private/checkpoints/one.json",
	})
	if err == nil {
		t.Fatal("expected sensitive evidence ref to fail")
	}
}

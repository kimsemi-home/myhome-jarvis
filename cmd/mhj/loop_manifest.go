package main

import "github.com/kimsemi-home/myhome-jarvis/internal/controlplane"

func appendControlPlaneManifest(root string, decisionKind string, selectedRoute string, outputRef string) (controlplane.RecordResult, error) {
	return controlplane.AppendManifest(root, controlplane.ManifestRequest{
		DecisionKind:     decisionKind,
		PolicyVersion:    "control-plane:v1",
		OntologyVersion:  "concepts:v1",
		AuthorityProfile: "local_readonly",
		SelectedRoute:    selectedRoute,
		ReviewerRole:     "go_review_gate",
		VerifierRole:     "deterministic_verifier",
		LeaseSeconds:     120,
		LeaseStatus:      "finished",
		EvidenceRefs: []string{
			"generated/control_plane.generated.json",
			"generated/planner.generated.json",
			"generated/concepts.generated.json",
		},
		OutputRef: outputRef,
	})
}

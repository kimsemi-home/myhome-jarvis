package main

import "fmt"

func verifyConformanceLinks(graph verificationGraphFile, conformance verificationConformanceFile) error {
	if conformance.GraphArtifact != "generated/verification_graph.generated.json" {
		return fmt.Errorf("conformance graph artifact mismatch")
	}
	if conformance.SchemaArtifact != "generated/verification_graph.schema.generated.json" {
		return fmt.Errorf("conformance schema artifact mismatch")
	}
	if conformance.TestsArtifact != "generated/verification_tests.generated.json" {
		return fmt.Errorf("conformance tests artifact mismatch")
	}
	if conformance.ReleaseArtifact != "generated/release_pipeline.generated.json" {
		return fmt.Errorf("conformance release artifact mismatch")
	}
	if conformance.ControlPlaneVerifierArtifact != "generated/control_plane_verification.generated.json" {
		return fmt.Errorf("conformance control-plane verifier artifact mismatch")
	}
	if conformance.VerificationEvidenceArtifact != "generated/verification_evidence.generated.json" {
		return fmt.Errorf("conformance verification evidence artifact mismatch")
	}
	if conformance.PDCAArtifact != "generated/pdca.generated.json" {
		return fmt.Errorf("conformance PDCA artifact mismatch")
	}
	if len(conformance.BackendArtifacts) != len(graph.Backends) {
		return fmt.Errorf("conformance backend count mismatch")
	}
	return nil
}

func verifyReleaseGates(graph verificationGraphFile, release verificationReleaseFile) error {
	gates := map[string]verificationGate{}
	for _, gate := range release.Gates {
		gates[gate.ID] = gate
	}
	for _, unit := range graph.Units {
		gate, ok := gates[unit.ID]
		if !ok || !gate.Required || gate.Kind != unit.Kind {
			return fmt.Errorf("release gate missing verification unit %q", unit.ID)
		}
	}
	return nil
}

package main

import "fmt"

func runVerificationVerify(root string) error {
	summary, err := validateVerificationGenerated(root)
	if err != nil {
		return err
	}
	return writeJSON(summary)
}

func validateVerificationGenerated(root string) (map[string]any, error) {
	graph, err := readVerificationJSON[verificationGraphFile](root, "generated/verification_graph.generated.json")
	if err != nil {
		return nil, err
	}
	if graph.SchemaVersion != "verification.graph/v1" {
		return nil, fmt.Errorf("verification graph schema mismatch")
	}
	if err := verifyGraphArtifacts(root, graph); err != nil {
		return nil, err
	}
	conformance, err := readVerificationJSON[verificationConformanceFile](root, "generated/verification_conformance.generated.json")
	if err != nil {
		return nil, err
	}
	release, err := readVerificationJSON[verificationReleaseFile](root, "generated/release_pipeline.generated.json")
	if err != nil {
		return nil, err
	}
	tests, err := readVerificationJSON[verificationTestsFile](root, "generated/verification_tests.generated.json")
	if err != nil {
		return nil, err
	}
	if err := verifyConformanceLinks(graph, conformance); err != nil {
		return nil, err
	}
	if err := verifyReleaseGates(graph, release); err != nil {
		return nil, err
	}
	if err := verifyTestManifest(tests); err != nil {
		return nil, err
	}
	return map[string]any{"ok": true, "artifact_count": len(graph.GeneratedArtifacts), "test_count": len(tests.Tests)}, nil
}

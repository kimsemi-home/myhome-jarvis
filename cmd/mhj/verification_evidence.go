package main

import (
	"fmt"

	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
)

func runVerificationEvidence(root string) error {
	summary, err := validateVerificationEvidence(root)
	if err != nil {
		return err
	}
	return writeJSON(summary)
}

func validateVerificationEvidence(root string) (map[string]any, error) {
	manifest, err := readVerificationJSON[verificationEvidenceFile](root, "generated/verification_evidence.generated.json")
	if err != nil {
		return nil, err
	}
	graph, err := readVerificationJSON[verificationGraphFile](root, manifest.GraphArtifact)
	if err != nil {
		return nil, err
	}
	evidence, err := readVerificationJSON[evidencePolicyFile](root, "generated/evidence.generated.json")
	if err != nil {
		return nil, err
	}
	quality, err := qualitylog.StatusForRoot(root)
	if err != nil {
		return nil, err
	}
	if err := verifyEvidenceManifest(manifest, graph, evidence, quality); err != nil {
		return nil, err
	}
	return map[string]any{"ok": true, "source_count": len(manifest.Sources),
		"quality_recorded": quality.Exists, "quality_run_count": quality.Count,
		"graph_evidence_count": len(graph.Evidence)}, nil
}

func verifyEvidenceManifest(manifest verificationEvidenceFile, graph verificationGraphFile, evidence evidencePolicyFile, quality qualitylog.Status) error {
	if manifest.SchemaVersion != "verification.evidence/v1" ||
		manifest.Command != "mhj verification evidence" {
		return fmt.Errorf("verification evidence manifest mismatch")
	}
	if manifest.QualityJournal != qualitylog.RelativePath || quality.Path != qualitylog.RelativePath {
		return fmt.Errorf("verification evidence quality journal mismatch")
	}
	if !qualitySourceLinked(evidence.PrivateSources) {
		return fmt.Errorf("verification evidence missing quality source")
	}
	return verifyEvidenceSources(manifest, graph, quality)
}

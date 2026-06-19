package main

import (
	"fmt"

	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
)

func verifyEvidenceSources(manifest verificationEvidenceFile, graph verificationGraphFile, quality qualitylog.Status) error {
	covered := map[string]bool{}
	for _, source := range manifest.Sources {
		if source.ID == "" || source.Kind == "" || source.Evidence == "" {
			return fmt.Errorf("verification evidence source requires id, kind, and evidence")
		}
		covered[source.ID] = true
	}
	for _, id := range requiredEvidenceSources() {
		if !covered[id] {
			return fmt.Errorf("verification evidence source missing %q", id)
		}
	}
	if !stringSet(graph.Evidence)["redacted local quality run ledger"] {
		return fmt.Errorf("verification graph missing quality run ledger evidence")
	}
	return verifyQualityShape(quality)
}

func qualitySourceLinked(sources []evidencePrivateSource) bool {
	for _, source := range sources {
		if source.Key == "quality" && source.Path == qualitylog.RelativePath &&
			source.NodeKind == "quality_run" && source.Format == "jsonl" {
			return true
		}
	}
	return false
}

func verifyQualityShape(status qualitylog.Status) error {
	if status.Last == nil {
		return nil
	}
	last := status.Last
	if last.StepCount != len(last.Steps) ||
		last.PassCount+last.FailCount+last.SkipCount > last.StepCount {
		return fmt.Errorf("verification evidence quality run shape mismatch")
	}
	return nil
}

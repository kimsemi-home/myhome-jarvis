package externalevidence

import (
	"net/http"
	"time"
)

func CollectForRoot(root string, maxSources int) (CollectReport, error) {
	client := &http.Client{Timeout: 12 * time.Second}
	return collectForRoot(root, maxSources, client)
}

func collectForRoot(
	root string,
	maxSources int,
	client *http.Client,
) (CollectReport, error) {
	policy, err := readPolicy(root)
	if err != nil {
		return CollectReport{}, err
	}
	if err := ValidatePolicy(policy); err != nil {
		return CollectReport{}, err
	}
	manifest, err := readManifestSummary(root, policy.ManifestPath)
	if err != nil {
		return CollectReport{}, err
	}
	now := time.Now().UTC().Format(time.RFC3339)
	report := newCollectReport(policy, now)
	for _, source := range limitedSources(policy.SourceDescriptors, maxSources) {
		result := collectSource(root, policy, source, now, client, manifest.KnownRefs)
		report.Results = append(report.Results, result)
		applyCollectResult(&report, result)
	}
	if report.CollectedCount > 0 {
		if err := appendGoldSummary(root, policy, report, now); err != nil {
			return CollectReport{}, err
		}
	}
	report.CollectionRunState = collectRunState(report)
	return report, nil
}

func limitedSources(sources []SourceDescriptor, maxSources int) []SourceDescriptor {
	if maxSources <= 0 || maxSources >= len(sources) {
		return sources
	}
	return sources[:maxSources]
}

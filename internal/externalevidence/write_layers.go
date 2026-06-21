package externalevidence

import (
	"fmt"
)

func writeSourceLayers(
	root string,
	policy Policy,
	source SourceDescriptor,
	now string,
	result CollectResult,
	body []byte,
) error {
	rawRel := fmt.Sprintf("%s/%s-%s.json", policy.RawLayerPath,
		safeKey(source.Key), result.RawSHA256[:16])
	if err := writePrivateFile(root, rawRel, body); err != nil {
		return err
	}
	manifest := manifestForSource(now, source, result, rawRel, policy)
	if err := appendJSONLine(root, policy.ManifestPath, manifest); err != nil {
		return err
	}
	normalized := normalizedForSource(now, source, result)
	if err := appendJSONLine(root, policy.BronzeLayerPath, normalized); err != nil {
		return err
	}
	return appendJSONLine(root, policy.SilverLayerPath, normalized)
}

func appendGoldSummary(root string, policy Policy, report CollectReport, now string) error {
	record := goldSummaryRecord{
		At:             now,
		Source:         "external_evidence",
		Kind:           "collection_summary",
		EvidenceRef:    "external_evidence:run:" + now,
		SourceCount:    report.SourceCount,
		CollectedCount: report.CollectedCount,
		CachedCount:    report.CachedCount,
		FailedCount:    report.FailedCount,
	}
	return appendJSONLine(root, policy.GoldLayerPath, record)
}

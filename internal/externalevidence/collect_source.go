package externalevidence

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)

func collectSource(
	root string,
	policy Policy,
	source SourceDescriptor,
	now string,
	client *http.Client,
	knownRefs map[string]bool,
) CollectResult {
	key := safeKey(source.Key)
	if key == "" {
		return failedResult(source, "invalid_source_key")
	}
	fetch, err := fetchSource(policy, source, client)
	if err != nil {
		return failedResult(source, errorCategory(err))
	}
	hash := sha256.Sum256(fetch.Body)
	rawSHA := hex.EncodeToString(hash[:])
	ref := "external_evidence:" + key + ":" + rawSHA[:16]
	result := collectedResult(source, ref, rawSHA, fetch)
	if knownRefs[ref] {
		result.State = "cached"
		return result
	}
	if err := writeSourceLayers(root, policy, source, now, result, fetch.Body); err != nil {
		return failedResult(source, errorCategory(err))
	}
	knownRefs[ref] = true
	return result
}

func collectedResult(source SourceDescriptor, ref string, rawSHA string, fetch fetchResult) CollectResult {
	return CollectResult{
		SourceKey:      source.Key,
		SourceClass:    source.Class,
		State:          "collected",
		EvidenceRef:    ref,
		RawSHA256:      rawSHA,
		PayloadBytes:   len(fetch.Body),
		HTTPStatus:     fetch.Status,
		FreshnessHours: source.FreshnessHours,
	}
}

func failedResult(source SourceDescriptor, category string) CollectResult {
	return CollectResult{
		SourceKey:      source.Key,
		SourceClass:    source.Class,
		State:          "failed",
		ErrorCategory:  category,
		FreshnessHours: source.FreshnessHours,
	}
}

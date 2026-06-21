package externalevidence

func manifestForSource(
	now string,
	source SourceDescriptor,
	result CollectResult,
	rawRel string,
	policy Policy,
) manifestRecord {
	return manifestRecord{
		At:                      now,
		Source:                  source.Key,
		Kind:                    "external_signal",
		EvidenceRef:             result.EvidenceRef,
		SourceClass:             source.Class,
		Status:                  "collected",
		HTTPStatus:              result.HTTPStatus,
		PayloadBytes:            result.PayloadBytes,
		RawSHA256:               result.RawSHA256,
		RawPrivatePath:          rawRel,
		Preprocess:              source.Preprocess,
		FreshnessHours:          source.FreshnessHours,
		RawPayloadPublicAllowed: policy.RawPayloadPublicAllowed,
	}
}

func normalizedForSource(now string, source SourceDescriptor, result CollectResult) normalizedRecord {
	return normalizedRecord{
		At:             now,
		Source:         source.Key,
		Kind:           "external_signal",
		EvidenceRef:    result.EvidenceRef,
		SourceClass:    source.Class,
		PayloadBytes:   result.PayloadBytes,
		HTTPStatus:     result.HTTPStatus,
		Preprocess:     source.Preprocess,
		FreshnessHours: source.FreshnessHours,
	}
}

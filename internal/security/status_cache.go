package security

import "time"

func cachedHistoryAggregate(root string) (historyAggregate, CacheEvidence, error) {
	key, err := buildStatusCacheKey(root)
	if err != nil {
		return historyAggregate{}, CacheEvidence{}, err
	}
	record, ok, err := readStatusCache(root, key)
	if err != nil {
		return historyAggregate{}, CacheEvidence{}, err
	}
	if ok {
		return aggregateFromRecord(record), cacheEvidence(key, "hit"), nil
	}
	report, err := CheckHistory(root)
	if err != nil {
		return historyAggregate{}, CacheEvidence{}, err
	}
	aggregate := historyAggregate{OK: report.OK, FindingCount: len(report.Findings)}
	if err := writeStatusCache(root, newStatusCacheRecord(key, aggregate)); err != nil {
		return historyAggregate{}, CacheEvidence{}, err
	}
	return aggregate, cacheEvidence(key, "miss"), nil
}

func newStatusCacheRecord(key statusCacheKey, aggregate historyAggregate) statusCacheRecord {
	return statusCacheRecord{
		Version:             statusCacheVersion,
		Key:                 key.Key,
		Head:                key.Head,
		InputHash:           key.InputHash,
		HistoryOK:           aggregate.OK,
		HistoryFindingCount: aggregate.FindingCount,
		CheckedAt:           time.Now().UTC().Format(time.RFC3339),
		ValidationCommand:   statusCacheValidation,
	}
}

func aggregateFromRecord(record statusCacheRecord) historyAggregate {
	return historyAggregate{OK: record.HistoryOK, FindingCount: record.HistoryFindingCount}
}

func cacheEvidence(key statusCacheKey, state string) CacheEvidence {
	return CacheEvidence{
		Path:                    statusCachePath,
		State:                   state,
		Key:                     key.Key,
		InputHash:               key.InputHash,
		Head:                    key.Head,
		EvidenceRef:             statusCachePath,
		ValidationCommand:       statusCacheValidation,
		PublicSafe:              true,
		RawDetailsPublicAllowed: false,
	}
}

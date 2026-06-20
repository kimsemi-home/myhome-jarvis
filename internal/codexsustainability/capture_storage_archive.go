package codexsustainability

import (
	"errors"
	"os"

	"github.com/kimsemi-home/myhome-jarvis/internal/storagearchive"
)

type storageArchiveCapture struct {
	State             string
	CacheHitCount     int64
	CacheMissCount    int64
	CacheSavingsUnits int64
}

func recordsFromStorageArchive(root string, at string) ([]Record, storageArchiveCapture, error) {
	report, err := storagearchive.RunForRoot(root)
	if errors.Is(err, os.ErrNotExist) {
		return nil, storageArchiveCapture{State: "unavailable"}, nil
	}
	if err != nil {
		return nil, storageArchiveCapture{}, err
	}
	capture := storageCaptureFromReport(report)
	return storageRecordsFromCapture(capture, at), capture, nil
}

func storageCaptureFromReport(report storagearchive.RunReport) storageArchiveCapture {
	capture := storageArchiveCapture{
		State:          "recorded",
		CacheHitCount:  int64(report.CachedCount),
		CacheMissCount: int64(report.ArchivedCount),
	}
	for _, result := range report.Results {
		if result.State == "cached" {
			capture.CacheSavingsUnits += result.InputBytes
		}
	}
	return capture
}

func storageRecordsFromCapture(capture storageArchiveCapture, at string) []Record {
	records := make([]Record, 0, 3)
	records = appendMetric(records, at, "cache_hit_count", capture.CacheHitCount)
	records = appendMetric(records, at, "cache_miss_count", capture.CacheMissCount)
	records = appendMetric(records, at, "cache_savings_units", capture.CacheSavingsUnits)
	return records
}

func appendMetric(records []Record, at string, metric string, amount int64) []Record {
	if amount <= 0 || at == "" {
		return records
	}
	return append(records, Record{
		At:           at,
		RecordKind:   "usage_sample",
		Metric:       metric,
		Amount:       amount,
		EvidenceRefs: []string{storagearchive.PolicyRelativePath, "docs/storage.md"},
	})
}

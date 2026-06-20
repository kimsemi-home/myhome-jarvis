package codexsustainability

import (
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
)

func recordsFromQualityRun(run qualitylog.Run) []Record {
	minutes := qualityCycleMinutes(run.DurationMillis)
	if run.At == "" || minutes <= 0 {
		return nil
	}
	version := trendVersion(run.At)
	return []Record{
		{
			At:                   run.At,
			RecordKind:           "trend_baseline",
			Metric:               "elapsed_cycle_minutes",
			Amount:               minutes,
			TrendBaselineVersion: version,
			TrendMeasuredAt:      run.At,
			EvidenceRefs:         []string{qualitylog.RelativePath},
		},
		{
			At:           run.At,
			RecordKind:   "cycle_sample",
			Metric:       "elapsed_cycle_minutes",
			Amount:       minutes,
			EvidenceRefs: []string{qualitylog.RelativePath},
		},
	}
}

func qualityCycleMinutes(durationMillis int64) int64 {
	if durationMillis <= 0 {
		return 1
	}
	return (durationMillis + int64(time.Minute/time.Millisecond) - 1) /
		int64(time.Minute/time.Millisecond)
}

func trendVersion(at string) string {
	parsed, err := time.Parse(time.RFC3339, at)
	if err != nil {
		return "quality-run"
	}
	return "quality-run-" + parsed.UTC().Format("20060102T150405Z")
}

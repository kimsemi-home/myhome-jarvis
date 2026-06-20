package codexsustainability

import "time"

func applyTrendBaseline(status *Status, record Record) {
	status.TrendBaselineCount++
	if laterRFC3339(status.latestTrendAt, record.TrendMeasuredAt) == record.TrendMeasuredAt {
		status.latestTrendAt = record.TrendMeasuredAt
		status.LatestTrendBaselineVersion = record.TrendBaselineVersion
		status.trendBaselineCycleMinutes = record.Amount
	}
}

func trendPosture(policy Policy, status Status, now time.Time) string {
	if status.TrendBaselineCount == 0 {
		return "missing"
	}
	if ageState(now, status.latestTrendAt, policy.TrendBaselineMaxAgeHours) != "fresh" {
		return "stale"
	}
	if status.MedianCycleMinutes > 0 &&
		status.trendBaselineCycleMinutes > 0 &&
		status.MedianCycleMinutes > status.trendBaselineCycleMinutes {
		return "slower_than_trend"
	}
	return "on_trend"
}

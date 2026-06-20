package codexsustainability

import (
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
)

func CaptureQualityRun(root string) (CaptureStatus, error) {
	return captureQualityRunAt(root, time.Now().UTC())
}

func captureQualityRunAt(root string, now time.Time) (CaptureStatus, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return CaptureStatus{}, err
	}
	quality, err := qualitylog.StatusForRoot(root)
	if err != nil {
		return CaptureStatus{}, err
	}
	status := newCaptureStatus(policy, quality, now)
	if !quality.Exists || quality.Last == nil {
		status.CaptureState = "missing_quality"
		return status, nil
	}
	status.LastQualityOK = quality.Last.OK
	if !quality.Last.OK {
		status.CaptureState = "last_run_not_successful"
		return status, nil
	}
	records := recordsFromQualityRun(*quality.Last)
	if len(records) == 0 {
		status.CaptureState = "last_run_not_successful"
		return status, nil
	}
	status.TrendBaselineVersion = records[0].TrendBaselineVersion
	status.CycleMinutes = records[0].Amount
	if capturedVersionExists(root, policy, status.TrendBaselineVersion) {
		status.CaptureState = "already_recorded"
		return status, nil
	}
	if err := appendRecords(root, policy, records); err != nil {
		return CaptureStatus{}, err
	}
	status.CaptureState = "recorded"
	status.RecordedRecordCount = len(records)
	return status, nil
}

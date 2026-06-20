package mediareadiness

import "time"

func StatusForRoot(root string) (Status, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Status{}, err
	}
	status := newStatus(policy)
	status.LocalLauncherAvailable, status.LocalLauncherProbe = localLauncherAvailable()
	for _, item := range policy.Cases {
		caseStatus := runCase(item)
		status.Cases = append(status.Cases, caseStatus)
		status.CaseCount++
		if caseStatus.Available && caseStatus.PlanningLatencyMS <= policy.TargetPlanningLatencyMS {
			status.AvailableCount++
			applyPlaybackCase(&status, caseStatus, true)
		} else {
			status.DegradedCount++
			applyPlaybackCase(&status, caseStatus, false)
		}
		if caseStatus.PlanningLatencyMS > status.MaxPlanningLatencyMS {
			status.MaxPlanningLatencyMS = caseStatus.PlanningLatencyMS
		}
	}
	return status, nil
}

func applyPlaybackCase(status *Status, item CaseStatus, available bool) {
	if item.Capability != "playback_readiness" {
		return
	}
	status.PlaybackCaseCount++
	if available {
		status.PlaybackAvailableCount++
	}
	status.PlaybackReady = status.PlaybackCaseCount > 0 &&
		status.PlaybackAvailableCount == status.PlaybackCaseCount
}

func newStatus(policy Policy) Status {
	return Status{
		Context:                 policy.Context,
		Version:                 policy.Version,
		PolicyPath:              PolicyRelativePath,
		BenchmarkKind:           policy.BenchmarkKind,
		PublicSafe:              true,
		Redaction:               "case-metadata-only",
		TargetPlanningLatencyMS: policy.TargetPlanningLatencyMS,
		CheckedAt:               time.Now().UTC().Format(time.RFC3339),
	}
}

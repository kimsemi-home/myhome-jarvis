package commandcenter

func summarizeLocalRuntime(supervisor SupervisorSummary) LocalRuntimeSummary {
	debt := localRuntimeDebt(supervisor)
	return LocalRuntimeSummary{
		PublicSafe:              true,
		EvidenceRef:             "local_runtime:supervisor",
		State:                   localRuntimeState(debt),
		Recorded:                supervisor.Recorded,
		ProcessRunning:          supervisor.ProcessRunning,
		ProbeOK:                 supervisor.ProbeOK,
		Stale:                   supervisor.Stale,
		HealthDebtCount:         debt,
		Message:                 supervisor.Message,
		NextSafeAction:          localRuntimeNextAction(debt),
		RawRuntimePublicAllowed: false,
	}
}

func localRuntimeDebt(supervisor SupervisorSummary) int {
	debt := 0
	if !supervisor.Recorded {
		debt++
	}
	if supervisor.Stale {
		debt++
	}
	if !supervisor.ProcessRunning {
		debt++
	}
	if !supervisor.ProbeOK {
		debt++
	}
	return debt
}

func localRuntimeState(debt int) string {
	if debt == 0 {
		return "healthy"
	}
	return "unhealthy"
}

func localRuntimeNextAction(debt int) string {
	if debt == 0 {
		return "none"
	}
	return "repair_local_runtime_health"
}

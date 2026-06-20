package commandcenter

import "testing"

func TestLocalRuntimeSummaryReportsSupervisorDebt(t *testing.T) {
	runtime := summarizeLocalRuntime(SupervisorSummary{
		Stale: true, Message: "daemon state is stale",
	})
	if !runtime.PublicSafe || runtime.EvidenceRef != "local_runtime:supervisor" {
		t.Fatalf("runtime evidence = %#v", runtime)
	}
	if runtime.State != "unhealthy" || runtime.HealthDebtCount != 4 {
		t.Fatalf("runtime health = %#v", runtime)
	}
	if runtime.RawRuntimePublicAllowed ||
		runtime.NextSafeAction != "repair_local_runtime_health" {
		t.Fatalf("runtime public action = %#v", runtime)
	}
}

func TestLocalRuntimeHealthySummaryHasNoDebt(t *testing.T) {
	runtime := summarizeLocalRuntime(SupervisorSummary{
		Recorded: true, ProcessRunning: true, ProbeOK: true,
	})
	if runtime.State != "healthy" || runtime.HealthDebtCount != 0 ||
		runtime.NextSafeAction != "none" {
		t.Fatalf("healthy runtime = %#v", runtime)
	}
}

package mediareadiness

import (
	"net/url"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
)

func runCase(item BenchmarkCase) CaseStatus {
	start := time.Now()
	payload, err := payloadFor(item.PayloadKind)
	if err == nil {
		var plan commands.Plan
		plan, err = commands.Build(item.Command, payload)
		return buildCaseStatus(item, plan, time.Since(start), err)
	}
	return buildCaseStatus(item, commands.Plan{}, time.Since(start), err)
}

func buildCaseStatus(item BenchmarkCase, plan commands.Plan, elapsed time.Duration, err error) CaseStatus {
	status := CaseStatus{
		ID:                item.ID,
		Capability:        item.Capability,
		Command:           item.Command,
		Available:         err == nil && len(plan.Invocations) > 0 && hostMatches(plan, item.ExpectedHost),
		PlanningLatencyMS: elapsed.Milliseconds(),
		InvocationCount:   len(plan.Invocations),
		ExpectedHost:      item.ExpectedHost,
	}
	status.Availability = availabilityLabel(status.Available)
	return status
}

func hostMatches(plan commands.Plan, expected string) bool {
	if expected == "" {
		return true
	}
	for _, invocation := range plan.Invocations {
		parsed, err := url.Parse(invocation.URL)
		if err == nil && parsed.Host == expected {
			return true
		}
	}
	return false
}

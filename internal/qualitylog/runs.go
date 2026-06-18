package qualitylog

import (
	"strings"
	"time"
)

func NewRun(started time.Time, ok bool, steps []Step) Run {
	run := Run{
		At:             started.UTC().Format(time.RFC3339),
		OK:             ok,
		DurationMillis: time.Since(started).Milliseconds(),
		StepCount:      len(steps),
		Steps:          make([]Step, 0, len(steps)),
	}
	for _, step := range steps {
		step.Name = strings.TrimSpace(step.Name)
		step.Status = normalizeStatus(step.Status)
		if step.Name == "" {
			step.Name = "unknown"
		}
		switch step.Status {
		case "pass":
			run.PassCount++
		case "fail":
			run.FailCount++
		case "skip":
			run.SkipCount++
		}
		run.Steps = append(run.Steps, step)
	}
	return run
}

func normalizeStatus(status string) string {
	status = strings.TrimSpace(strings.ToLower(status))
	switch status {
	case "pass", "fail", "skip":
		return status
	default:
		return "unknown"
	}
}

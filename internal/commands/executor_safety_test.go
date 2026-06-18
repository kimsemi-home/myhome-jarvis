package commands

import (
	"context"
	"testing"
)

func TestExecuteRejectsUnsafeExecutable(t *testing.T) {
	plan := Plan{
		Name:           "unsafe",
		DryRun:         true,
		ExecuteAllowed: true,
		Invocations: []Invocation{{
			Label: "unsafe",
			Argv:  []string{"sh", "-c", "echo no"},
		}},
	}
	if _, err := Execute(context.Background(), plan, ExecuteOptions{Platform: "darwin"}); err == nil {
		t.Fatal("expected unsafe executable to fail")
	}
}

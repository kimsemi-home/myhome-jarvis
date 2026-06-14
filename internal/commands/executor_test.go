package commands

import (
	"context"
	"testing"
)

func TestExecuteRunsAllowedArgvOnDarwin(t *testing.T) {
	plan, err := Build("volume-set", []byte(`{"level":30}`))
	if err != nil {
		t.Fatal(err)
	}
	plan = WithExecuteAllowed(plan, true)
	calls := 0
	executed, err := Execute(context.Background(), plan, ExecuteOptions{
		Platform: "darwin",
		Runner: func(context.Context, Invocation) Execution {
			calls++
			return Execution{Executed: true, ExitCode: 0}
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if executed.DryRun || calls != 1 {
		t.Fatalf("expected one real execution: %#v calls=%d", executed, calls)
	}
	if len(executed.Executions) != 1 || executed.Executions[0].Argv[0] != "osascript" {
		t.Fatalf("unexpected executions: %#v", executed.Executions)
	}
}

func TestExecuteSkipsWhenNotAllowed(t *testing.T) {
	plan, err := Build("display-sleep", []byte(`{}`))
	if err != nil {
		t.Fatal(err)
	}
	executed, err := Execute(context.Background(), plan, ExecuteOptions{
		Platform: "darwin",
		Runner: func(context.Context, Invocation) Execution {
			t.Fatal("runner must not be called")
			return Execution{}
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !executed.DryRun || len(executed.Executions) != 0 || len(executed.Warnings) == 0 {
		t.Fatalf("expected dry-run skip: %#v", executed)
	}
}

func TestExecuteSkipsOnNonDarwin(t *testing.T) {
	plan, err := Build("open-youtube", []byte(`{}`))
	if err != nil {
		t.Fatal(err)
	}
	plan = WithExecuteAllowed(plan, true)
	executed, err := Execute(context.Background(), plan, ExecuteOptions{
		Platform: "linux",
		Runner: func(context.Context, Invocation) Execution {
			t.Fatal("runner must not be called")
			return Execution{}
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !executed.DryRun || len(executed.Executions) != 1 || !executed.Executions[0].Skipped {
		t.Fatalf("expected platform skip: %#v", executed)
	}
}

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

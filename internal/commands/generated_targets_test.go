package commands

import "testing"

func TestGeneratedCommandTargetsMatchRegistry(t *testing.T) {
	catalog := loadGeneratedCommandCatalog(t)
	for _, command := range catalog.Commands {
		if command.Target == "" {
			continue
		}
		plan, err := Build(command.Name, []byte(`{}`))
		if err != nil {
			t.Fatalf("building generated target command %q: %v", command.Name, err)
		}
		if len(plan.Invocations) != 1 || plan.Invocations[0].URL != command.Target {
			t.Fatalf("target mismatch for %q: plan %#v generated target %q", command.Name, plan.Invocations, command.Target)
		}
	}
}

func TestBuildVolumeSetRejectsOutOfRange(t *testing.T) {
	if _, err := Build("volume-set", []byte(`{"level":101}`)); err == nil {
		t.Fatal("expected out-of-range level to fail")
	}
}

func TestBuildOpenURLRejectsUnsafeScheme(t *testing.T) {
	if _, err := Build("open-url", []byte(`{"url":"javascript:alert(1)"}`)); err == nil {
		t.Fatal("expected unsafe URL to fail")
	}
}

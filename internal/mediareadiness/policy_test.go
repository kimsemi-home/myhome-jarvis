package mediareadiness

import "testing"

func TestPolicyRequiresPublicSafeDryRunBenchmark(t *testing.T) {
	policy := testPolicy()
	policy.ExecuteCommands = true

	if err := ValidatePolicy(policy); err == nil {
		t.Fatal("expected executing benchmark policy to fail")
	}
}

package codeshape

import "fmt"

func validatePolicy(policy Policy) error {
	if policy.Context != "AgentCluster" {
		return fmt.Errorf("code-shape policy context = %q", policy.Context)
	}
	if policy.GeneratedArtifact != PolicyRelativePath {
		return fmt.Errorf("code-shape artifact path is invalid")
	}
	if policy.MaxFileLines != 75 {
		return fmt.Errorf("code-shape max file lines must stay 75")
	}
	if !policy.PublicStatusRedacted {
		return fmt.Errorf("code-shape public status must stay redacted")
	}
	if len(policy.SourceRoots) == 0 || len(policy.Extensions) == 0 {
		return fmt.Errorf("code-shape roots and extensions are required")
	}
	for _, root := range policy.SourceRoots {
		if err := validateRelPath(root); err != nil {
			return err
		}
	}
	for _, prefix := range policy.ExcludedPrefixes {
		if err := validateRelPath(prefix); err != nil {
			return err
		}
	}
	for _, extension := range policy.Extensions {
		if len(extension) < 2 || extension[0] != '.' {
			return fmt.Errorf("code-shape extension %q is invalid", extension)
		}
	}
	if !contains(policy.Commands, "mhj code-shape status") {
		return fmt.Errorf("code-shape status command is missing")
	}
	for _, field := range []string{"budget_regression_count", "legacy_debt_count", "ok"} {
		if !contains(policy.PublicSummaryFields, field) {
			return fmt.Errorf("code-shape summary field %q is missing", field)
		}
	}
	for _, entry := range policy.LegacyDebtFiles {
		if err := validateRelPath(entry.Path); err != nil {
			return err
		}
		if entry.MaxLines <= policy.MaxFileLines {
			return fmt.Errorf("legacy debt %q must exceed max lines", entry.Path)
		}
	}
	return nil
}

package externalevidence

import (
	"fmt"
	"net/url"
	"strings"
)

func ValidatePolicy(policy Policy) error {
	if policy.SchemaVersion != "external_evidence/v1" || !policy.PublicSafe {
		return fmt.Errorf("external evidence policy is not public-safe")
	}
	if policy.CredentialsAllowed || policy.CookiesAllowed ||
		policy.RawPayloadPublicAllowed {
		return fmt.Errorf("external evidence policy exposes forbidden collection")
	}
	for _, path := range privatePaths(policy) {
		if !strings.HasPrefix(path, "data/private/") {
			return fmt.Errorf("external evidence path must stay private")
		}
	}
	for _, source := range policy.SourceDescriptors {
		if err := validateSource(policy, source); err != nil {
			return err
		}
	}
	if !contains(policy.Commands, "mhj external-evidence repo-split-decision") {
		return fmt.Errorf("external evidence repo split decision command required")
	}
	if !contains(policy.Commands, "mhj external-evidence repo-bootstrap") {
		return fmt.Errorf("external evidence repo bootstrap command required")
	}
	return validateRepoSplit(policy.RepoSplitAssessment)
}

func privatePaths(policy Policy) []string {
	return []string{policy.PrivateRoot, policy.ManifestPath, policy.RawLayerPath,
		policy.BronzeLayerPath, policy.SilverLayerPath, policy.GoldLayerPath,
		policy.StorageArchiveSourcePath}
}

func validateSource(policy Policy, source SourceDescriptor) error {
	if source.Method != "GET" || !contains(policy.SourceClasses, source.Class) {
		return fmt.Errorf("external evidence source %q is not allowed", source.Key)
	}
	parsed, err := url.Parse(source.URL)
	if err != nil || parsed.Scheme != "https" || parsed.User != nil {
		return fmt.Errorf("external evidence source %q must use credential-free HTTPS", source.Key)
	}
	if source.FreshnessHours <= 0 || source.Preprocess == "" {
		return fmt.Errorf("external evidence source %q is incomplete", source.Key)
	}
	return nil
}

func validateRepoSplit(assessment RepoSplitAssessment) error {
	if assessment.CreationGate != "authority_review_required" {
		return fmt.Errorf("external evidence repo split must stay authority gated")
	}
	for _, required := range []string{"no_raw_payloads", "private_data_stays_private"} {
		if !contains(assessment.PublicRepoRules, required) {
			return fmt.Errorf("external evidence repo split missing %s", required)
		}
	}
	return nil
}

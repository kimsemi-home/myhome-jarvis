package controlplane

import "fmt"

func validateManifestRoles(policy Policy, manifest Manifest) error {
	if manifest.ReviewerRole == "" || manifest.VerifierRole == "" {
		return fmt.Errorf("control-plane manifest requires reviewer_role and verifier_role")
	}
	if policy.VerifierSeparationRequired && manifest.ReviewerRole == manifest.VerifierRole {
		return fmt.Errorf("control-plane reviewer and verifier roles must be separated")
	}
	if manifest.LeaseSeconds < policy.MinLeaseSeconds ||
		manifest.LeaseSeconds > policy.MaxLeaseSeconds {
		return fmt.Errorf("control-plane lease seconds out of bounds")
	}
	if !contains(normalizeList(policy.AllowedLeaseStatuses), manifest.LeaseStatus) {
		return fmt.Errorf("control-plane lease status %q is not allowed", manifest.LeaseStatus)
	}
	return nil
}

package controlplane

import "fmt"

func validateManifest(policy Policy, manifest Manifest) error {
	for _, validate := range []func(Policy, Manifest) error{
		validateManifestKinds,
		validateManifestRoles,
		validateManifestRefs,
	} {
		if err := validate(policy, manifest); err != nil {
			return err
		}
	}
	return nil
}

func validateManifestKinds(policy Policy, manifest Manifest) error {
	if !contains(normalizeList(policy.AllowedDecisionKinds), manifest.DecisionKind) {
		return fmt.Errorf("control-plane decision kind %q is not allowed", manifest.DecisionKind)
	}
	if manifest.PolicyVersion == "" || manifest.OntologyVersion == "" ||
		manifest.SelectedRoute == "" {
		return fmt.Errorf("control-plane manifest requires policy_version, ontology_version, and selected_route")
	}
	if !contains(normalizeList(policy.AllowedAuthorityProfiles), manifest.AuthorityProfile) {
		return fmt.Errorf("control-plane authority profile %q is not allowed", manifest.AuthorityProfile)
	}
	return nil
}

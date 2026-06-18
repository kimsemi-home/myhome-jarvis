package controlplane

import "fmt"

func validateManifestRefs(policy Policy, manifest Manifest) error {
	if len(manifest.EvidenceRefs) == 0 || manifest.OutputRef == "" {
		return fmt.Errorf("control-plane manifest requires evidence refs and output ref")
	}
	for _, ref := range append(append([]string{}, manifest.EvidenceRefs...), manifest.OutputRef) {
		if err := validateRef(policy, ref); err != nil {
			return err
		}
	}
	return rejectManifestSensitiveValues(policy, manifest)
}

func rejectManifestSensitiveValues(policy Policy, manifest Manifest) error {
	values := append([]string{
		manifest.PolicyVersion,
		manifest.OntologyVersion,
		manifest.SelectedRoute,
		manifest.ReviewerRole,
		manifest.VerifierRole,
		manifest.OutputRef,
	}, manifest.EvidenceRefs...)
	for _, value := range values {
		if err := rejectSensitiveText(value); err != nil {
			return err
		}
		if err := rejectForbiddenPolicyText(policy, value); err != nil {
			return err
		}
	}
	return nil
}

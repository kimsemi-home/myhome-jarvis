package controlplane

import "fmt"

func VerifyForRoot(root string) (VerificationStatus, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return VerificationStatus{}, err
	}
	verifier, err := readVerificationPolicy(root)
	if err != nil {
		return VerificationStatus{}, err
	}
	if err := validateVerificationPolicy(policy, verifier); err != nil {
		return VerificationStatus{}, err
	}
	status, err := StatusForRoot(root)
	if err != nil {
		return VerificationStatus{}, err
	}
	if err := verifyStatus(policy, status); err != nil {
		return VerificationStatus{}, err
	}
	return VerificationStatus{
		OK:                     true,
		PolicyPath:             PolicyRelativePath,
		VerifierPath:           VerificationRelativePath,
		CheckCount:             len(verifier.Checks),
		ManifestDebtCount:      status.ManifestDebtCount,
		VerifierViolationCount: status.VerifierViolationCount,
		CheckedAt:              status.CheckedAt,
	}, nil
}

func validateVerificationPolicy(policy Policy, verifier VerificationPolicy) error {
	if verifier.SchemaVersion != "control-plane.verification/v1" {
		return fmt.Errorf("control-plane verifier schema mismatch")
	}
	if policy.VerifierGeneratedArtifact != VerificationRelativePath ||
		verifier.PolicyArtifact != PolicyRelativePath {
		return fmt.Errorf("control-plane verifier artifact link mismatch")
	}
	if policy.VerificationCommand != "mhj control-plane verify" ||
		verifier.Command != policy.VerificationCommand {
		return fmt.Errorf("control-plane verifier command mismatch")
	}
	if verifier.StatusCommand != "mhj control-plane status" ||
		!policy.VerifierSeparationRequired || !verifier.VerifierSeparationRequired {
		return fmt.Errorf("control-plane verifier separation mismatch")
	}
	return requireVerifierChecks(verifier)
}

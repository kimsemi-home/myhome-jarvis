package contextpack

import "fmt"

func validatePolicy(policy Policy) error {
	if policy.Context != "CrossRepoContextPack" {
		return fmt.Errorf("context pack context = %q", policy.Context)
	}
	if policy.PackID == "" || policy.UpstreamCompatibilityVersion == "" ||
		policy.OntologyVersion == "" || policy.DeclarationPath == "" {
		return fmt.Errorf("context pack identity fields are required")
	}
	if !policy.PublicStatusRedacted || policy.RawPrivateContextPublicAllowed {
		return fmt.Errorf("context pack must be public redacted")
	}
	if policy.AuthorityContract.SelfApprovalAllowed ||
		!policy.AuthorityContract.PublicSafetyGateRequired ||
		policy.SecurityContract.PrivatePathsPublicAllowed ||
		policy.SecurityContract.LocalPathsPublicAllowed {
		return fmt.Errorf("context pack authority/security contracts are unsafe")
	}
	if err := requireAll("split criterion", splitKeys(policy), requiredSplitCriteria); err != nil {
		return err
	}
	if err := requireAll("export role", artifactRoles(policy.ExportedArtifacts), requiredExportRoles); err != nil {
		return err
	}
	if err := requireAll("declaration field", policy.RequiredDeclarationFields, requiredDeclarationFields); err != nil {
		return err
	}
	if err := requireAll("command", policy.Commands, requiredCommands); err != nil {
		return err
	}
	return nil
}

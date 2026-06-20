package contextpack

import "time"

func VerifyDeclarationForRoot(root string, declarationPath string) (VerifyResult, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return VerifyResult{}, err
	}
	declaration, resolved, err := ReadDeclaration(root, declarationPath)
	result := newVerifyResult(resolved)
	if err != nil {
		result.Valid = false
		result.Findings = append(result.Findings, "declaration_missing_or_unreadable")
		result.MissingFieldCount++
		return result, nil
	}
	verifyIdentity(policy, declaration, &result)
	verifyArtifacts(policy, declaration, &result)
	verifyForbiddenValues(declaration, &result)
	result.Valid = result.DriftCount == 0 &&
		result.MissingFieldCount == 0 &&
		result.MissingArtifactCount == 0 &&
		result.StaleVersionCount == 0 &&
		result.ForbiddenValueCount == 0
	return result, nil
}

func newVerifyResult(path string) VerifyResult {
	return VerifyResult{DeclarationPath: path,
		CheckedAt: time.Now().UTC().Format(time.RFC3339)}
}

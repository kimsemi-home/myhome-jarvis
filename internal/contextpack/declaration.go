package contextpack

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func ReadDeclaration(root string, declarationPath string) (Declaration, string, error) {
	policy, err := ReadPolicy(root)
	if err != nil {
		return Declaration{}, "", err
	}
	resolved := resolveDeclarationPath(root, policy, declarationPath)
	body, err := os.ReadFile(resolved)
	if err != nil {
		return Declaration{}, resolved, err
	}
	var declaration Declaration
	if err := json.Unmarshal(body, &declaration); err != nil {
		return Declaration{}, resolved, err
	}
	return declaration, resolved, nil
}

func resolveDeclarationPath(root string, policy Policy, declarationPath string) string {
	if declarationPath == "" {
		declarationPath = policy.DeclarationPath
	}
	if filepath.IsAbs(declarationPath) {
		return declarationPath
	}
	return filepath.Join(root, filepath.FromSlash(declarationPath))
}

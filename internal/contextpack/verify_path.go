package contextpack

import "path/filepath"

func declarationResultPath(
	root string,
	policy Policy,
	requested string,
	resolved string,
) string {
	if requested == "" {
		return policy.DeclarationPath
	}
	if !filepath.IsAbs(requested) {
		return filepath.ToSlash(requested)
	}
	rel, err := filepath.Rel(root, resolved)
	if err == nil && !startsWithParent(rel) {
		return filepath.ToSlash(rel)
	}
	return filepath.Base(resolved)
}

func startsWithParent(path string) bool {
	return path == ".." || len(path) > 3 && path[:3] == "../"
}

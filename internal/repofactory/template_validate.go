package repofactory

import "fmt"

func validateTemplateFiles(policy Policy) error {
	roles := templateRoles(policy.TemplateFiles)
	if missing := containsAll(roles, requiredTemplateRoles); len(missing) > 0 {
		return fmt.Errorf("repo factory template role missing %q", missing[0])
	}
	for _, file := range policy.TemplateFiles {
		if err := validateTemplateFile(file, policy.AllowedPublicPathPrefixes,
			policy.ForbiddenPublicFragments); err != nil {
			return err
		}
	}
	return nil
}

func validateTemplateFile(file TemplateFile, prefixes []string, fragments []string) error {
	if file.Role == "" || file.Path == "" || file.SourceArtifact == "" || file.Purpose == "" {
		return fmt.Errorf("repo factory template file %q is incomplete", file.Role)
	}
	if !allowedPublicPath(file.Path, prefixes) {
		return fmt.Errorf("repo factory template path is not public-safe: %q", file.Path)
	}
	if forbiddenTemplateValue(file, fragments) {
		return fmt.Errorf("repo factory template contains a forbidden public value")
	}
	return nil
}

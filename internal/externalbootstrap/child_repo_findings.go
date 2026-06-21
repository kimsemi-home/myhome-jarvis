package externalbootstrap

import "path/filepath"

func (status *ChildRepoStatus) addFinding(path string, code string, message string) {
	status.Findings = append(status.Findings, ChildRepoFinding{
		Path:    filepath.ToSlash(path),
		Code:    code,
		Message: message,
	})
}

func childRequiredFiles(packet Packet) []string {
	files := []string{"README.md", ".gitignore"}
	for _, file := range packet.SkeletonFiles {
		files = append(files, file.Path)
	}
	return uniqueStrings(files)
}

func uniqueStrings(values []string) []string {
	seen := map[string]bool{}
	result := []string{}
	for _, value := range values {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	return result
}

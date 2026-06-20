package repofactory

import "strings"

func allowedPublicPath(path string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(path, prefix) {
			return true
		}
	}
	return false
}

func forbiddenTemplateValue(file TemplateFile, fragments []string) bool {
	values := []string{file.Path, file.SourceArtifact, file.Purpose}
	for _, value := range values {
		if containsPrivateMarker(value) || containsForbiddenFragment(value, fragments) {
			return true
		}
	}
	return false
}

func containsForbiddenFragment(value string, fragments []string) bool {
	for _, fragment := range fragments {
		if fragment != "" && strings.Contains(value, fragment) {
			return true
		}
	}
	return false
}

func containsPrivateMarker(value string) bool {
	for _, fragment := range privateMarkerFragments() {
		if strings.Contains(value, fragment) {
			return true
		}
	}
	return false
}

func privateMarkerFragments() []string {
	oldOwner := strings.Join([]string{"kim", "jooyoon"}, "")
	oldTeam := strings.Join([]string{"kim-joo", "-yoon"}, "")
	return []string{"/" + "Users" + "/", oldOwner, oldTeam, "github.com/" + oldOwner}
}

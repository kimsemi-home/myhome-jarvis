package review

import (
	"fmt"
	"path/filepath"
	"strings"
)

func validateReviewRefs(policy Policy, refs []string) error {
	for _, ref := range refs {
		if err := validateRef(policy, ref); err != nil {
			return err
		}
		if err := rejectSensitiveText(ref); err != nil {
			return err
		}
	}
	return nil
}

func validateRef(policy Policy, ref string) error {
	ref = filepath.ToSlash(strings.TrimSpace(ref))
	if ref == "" {
		return errMissingEvidenceRef
	}
	if filepath.IsAbs(filepath.FromSlash(ref)) || strings.Contains(ref, "..") {
		return fmt.Errorf("review ref must be repo-relative")
	}
	for _, prefix := range policy.AllowedEvidencePrefixes {
		if strings.HasPrefix(ref, prefix) {
			return nil
		}
	}
	return fmt.Errorf("review ref is outside allowed prefixes")
}

func rejectSensitiveText(value string) error {
	lower := strings.ToLower(value)
	for _, marker := range sensitiveMarkers() {
		if strings.Contains(lower, marker) {
			return fmt.Errorf("review ref contains forbidden private marker")
		}
	}
	return nil
}

func sensitiveMarkers() []string {
	return []string{
		"kim" + "jooyoon", "kim-joo" + "-yoon",
		"/us" + "ers/" + "al" + "ice", "al" + "ice/" + "git" + "hub",
		"bearer ", "begin private key", "raw_prompt", "raw_transcript",
		"account_id", "card_number", "api_secret", "credential=",
		"linear." + "app/", string(filepath.Separator) + "users" + string(filepath.Separator),
		"\\" + "users" + "\\",
	}
}

package controlplane

import (
	"fmt"
	"path/filepath"
	"strings"
)

func rejectSensitiveText(value string) error {
	lower := strings.ToLower(value)
	for _, marker := range sensitiveMarkers() {
		if strings.Contains(lower, marker) {
			return fmt.Errorf("control-plane manifest contains forbidden private marker")
		}
	}
	return nil
}

func sensitiveMarkers() []string {
	return []string{
		"kim" + "jooyoon",
		"kim-joo" + "-yoon",
		"/us" + "ers/" + "al" + "ice",
		"al" + "ice/" + "git" + "hub",
		"bearer ",
		"begin private key",
		"raw_rationale",
		"selection_rationale",
		"candidate_agents",
		"raw_prompt",
		"raw_transcript",
		"private_evidence",
		"account_id",
		"card_number",
		"api_secret",
		"credential=",
		"linear.app/",
		string(filepath.Separator) + "users" + string(filepath.Separator),
		"\\" + "users" + "\\",
	}
}

package incidents

import (
	"fmt"
	"path/filepath"
	"strings"
)

func rejectSensitiveText(value string) error {
	lower := strings.ToLower(value)
	for _, marker := range forbiddenIncidentMarkers() {
		if strings.Contains(lower, marker) {
			return fmt.Errorf("incident evidence ref contains forbidden private marker")
		}
	}
	return nil
}

func forbiddenIncidentMarkers() []string {
	return []string{
		"kim" + "jooyoon",
		"kim-joo" + "-yoon",
		"/us" + "ers/" + "al" + "ice",
		"al" + "ice/" + "git" + "hub",
		"bearer ",
		"begin private key",
		"raw_prompt",
		"raw_transcript",
		"account_id",
		"card_number",
		"api_secret",
		"credential=",
		"linear." + "app/",
		string(filepath.Separator) + "users" + string(filepath.Separator),
		"\\" + "users" + "\\",
	}
}

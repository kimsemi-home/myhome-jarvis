package evidencequality

import (
	"fmt"
	"path/filepath"
	"strings"
)

func rejectSensitiveText(value string) error {
	lower := strings.ToLower(value)
	for _, marker := range []string{
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
	} {
		if strings.Contains(lower, marker) {
			return fmt.Errorf("evidence quality ref contains forbidden private marker")
		}
	}
	return nil
}

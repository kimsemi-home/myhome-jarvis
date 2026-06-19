package linear

import "strings"

func filepathJoinSlash(parts ...string) string {
	if len(parts) == 0 {
		return ""
	}
	joined := parts[0]
	for _, part := range parts[1:] {
		joined = strings.TrimRight(joined, "/") + "/" + strings.TrimLeft(part, "/")
	}
	return joined
}

package knowledge

import "strings"

func stringsContainsEither(left string, right string) bool {
	return strings.Contains(left, right) || strings.Contains(right, left)
}

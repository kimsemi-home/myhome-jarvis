package storagearchive

import "fmt"

func toText(value any) string {
	switch typed := value.(type) {
	case string:
		return typed
	case float64, bool:
		return fmt.Sprint(typed)
	default:
		return ""
	}
}

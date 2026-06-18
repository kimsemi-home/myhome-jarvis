package audit

import "path/filepath"

func commandIntentPath(root string) string {
	return filepath.Join(root, filepath.FromSlash(commandIntentRelativePath))
}

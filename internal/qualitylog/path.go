package qualitylog

import "path/filepath"

func runPath(root string) string {
	return filepath.Join(root, filepath.FromSlash(RelativePath))
}

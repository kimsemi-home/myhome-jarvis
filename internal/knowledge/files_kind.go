package knowledge

import "path/filepath"

func indexableFile(path string) bool {
	switch filepath.Ext(path) {
	case ".go", ".lisp", ".dart", ".md", ".json", ".jsonl", ".toml", ".yaml", ".yml":
		return true
	default:
		return false
	}
}

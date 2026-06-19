package knowledge

import (
	"os"
	"path/filepath"
)

func mustReadFiles(root string, concepts []Concept, harnesses []HarnessCase, hits []Hit) []string {
	seen := map[string]bool{}
	var files []string
	files = appendUniqueRel(files, seen, RegistryRelativePath)
	files = appendConceptTargets(root, files, seen, concepts)
	files = appendHarnessTargets(root, files, seen, harnesses)
	for _, hit := range hits {
		files = appendUniqueRel(files, seen, hit.Path)
		if len(files) >= 12 {
			break
		}
	}
	return files
}

func appendConceptTargets(root string, files []string, seen map[string]bool, concepts []Concept) []string {
	for _, concept := range concepts {
		for _, target := range concept.GeneratedTargets {
			if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(target))); err == nil {
				files = appendUniqueRel(files, seen, target)
			}
		}
	}
	return files
}

func appendHarnessTargets(root string, files []string, seen map[string]bool, harnesses []HarnessCase) []string {
	for _, harness := range harnesses {
		if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(harness.EvidenceTarget))); err == nil {
			files = appendUniqueRel(files, seen, harness.EvidenceTarget)
		}
	}
	return files
}

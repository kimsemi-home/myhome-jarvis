package knowledge

import (
	"os"
	"regexp"
	"sort"
)

func searchFiles(root string, files []string, terms []termMatch, limit int) ([]Hit, []string, error) {
	var hits []Hit
	issueIDs := map[string]bool{}
	issuePattern := regexp.MustCompile(`\bKIM-[0-9]+\b`)
	for _, path := range files {
		err := searchFile(root, path, terms, limit, issuePattern, &hits, issueIDs)
		if err != nil {
			return nil, nil, err
		}
	}
	issues := make([]string, 0, len(issueIDs))
	for issue := range issueIDs {
		issues = append(issues, issue)
	}
	sort.Strings(issues)
	return hits, issues, nil
}

func openSearchFile(path string) (*os.File, error) {
	return os.Open(path)
}

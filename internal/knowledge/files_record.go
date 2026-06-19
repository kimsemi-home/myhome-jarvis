package knowledge

import (
	"path/filepath"
	"regexp"
)

func recordHit(
	root string,
	path string,
	lineNumber int,
	line string,
	match termMatch,
	limit int,
	issuePattern *regexp.Regexp,
	hits *[]Hit,
	issueIDs map[string]bool,
) error {
	rel, err := filepath.Rel(root, path)
	if err != nil {
		return err
	}
	if len(*hits) < limit {
		*hits = append(*hits, Hit{
			Path:    filepath.ToSlash(rel),
			Line:    lineNumber,
			Concept: match.Concept,
			Term:    match.Term,
		})
	}
	for _, issue := range issuePattern.FindAllString(line, -1) {
		issueIDs[issue] = true
	}
	return nil
}

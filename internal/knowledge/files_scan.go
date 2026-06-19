package knowledge

import (
	"bufio"
	"regexp"
	"strings"
)

func searchFile(
	root string,
	path string,
	terms []termMatch,
	limit int,
	issuePattern *regexp.Regexp,
	hits *[]Hit,
	issueIDs map[string]bool,
) error {
	file, err := openSearchFile(path)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		if err := scanLine(root, path, lineNumber, scanner.Text(), terms, limit, issuePattern, hits, issueIDs); err != nil {
			return err
		}
	}
	return scanner.Err()
}

func scanLine(
	root string,
	path string,
	lineNumber int,
	line string,
	terms []termMatch,
	limit int,
	issuePattern *regexp.Regexp,
	hits *[]Hit,
	issueIDs map[string]bool,
) error {
	lower := strings.ToLower(line)
	for _, match := range terms {
		if !matchLine(lower, match.Term) {
			continue
		}
		return recordHit(root, path, lineNumber, line, match, limit, issuePattern, hits, issueIDs)
	}
	return nil
}

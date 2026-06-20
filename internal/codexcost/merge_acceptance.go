package codexcost

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

type mergeAcceptance struct {
	Count  int64
	Source string
	Limit  int
}

func mergeAcceptanceForRoot(root string, policy Policy) mergeAcceptance {
	limit := policy.ROIMergeLogLimit
	if limit <= 0 {
		limit = 200
	}
	count, ok := countGitHubMergeCommits(root, limit)
	if !ok {
		return mergeAcceptance{Source: "unavailable", Limit: limit}
	}
	return mergeAcceptance{Count: count, Source: "git_merge_commits", Limit: limit}
}

func countGitHubMergeCommits(root string, limit int) (int64, bool) {
	args := []string{"-C", root, "log", "--first-parent",
		"--extended-regexp", "--grep=^Merge pull request #[0-9]+",
		"--pretty=%H", "-n", strconv.Itoa(limit)}
	cmd := exec.Command("git", args...)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Run(); err != nil {
		return 0, false
	}
	return int64(nonEmptyLineCount(output.String())), true
}

func nonEmptyLineCount(output string) int {
	count := 0
	for _, line := range strings.Split(output, "\n") {
		if strings.TrimSpace(line) != "" {
			count++
		}
	}
	return count
}

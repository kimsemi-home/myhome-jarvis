package codexcost

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestMergeAcceptanceCountsGitHubMergeCommits(t *testing.T) {
	root := initMergeRepo(t)
	acceptance := mergeAcceptanceForRoot(root, Policy{ROIMergeLogLimit: 20})
	if acceptance.Count != 1 || acceptance.Source != "git_merge_commits" {
		t.Fatalf("merge acceptance = %#v", acceptance)
	}
}

func TestROISummaryUsesMergeEvidenceForAcceptedChanges(t *testing.T) {
	root := initMergeRepo(t)
	writePolicy(t, root, testPolicy())
	writeSustainabilityPolicy(t, root)
	writeSustainableLedger(t, root)
	writeStoragePolicy(t, root)
	writeFile(t, root, "data/private/codex-cost/usage.jsonl",
		`{"at":"2026-06-19T00:00:00Z","scope":"repo","unit_kind":"codex_tokens","amount":100,"status":"recorded","evidence_refs":["docs/codex-cost-governor.md"]}`+"\n")

	summary, err := ROISummaryForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if summary.AcceptedChangeCount != 1 ||
		summary.MergeAcceptedChangeCount != 1 ||
		summary.CostPerAcceptedChange != 100 {
		t.Fatalf("accepted change evidence = %#v", summary)
	}
	if summary.ValueProxyUnits != 2 {
		t.Fatalf("value proxy = %#v", summary)
	}
}

func initMergeRepo(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	runGit(t, root, "init")
	runGit(t, root, "config", "user.email", "test@example.invalid")
	runGit(t, root, "config", "user.name", "Test User")
	writeRepoFile(t, root, "README.md", "root\n")
	runGit(t, root, "add", ".")
	runGit(t, root, "commit", "-m", "initial")
	runGit(t, root, "branch", "base")
	runGit(t, root, "checkout", "-b", "feature")
	writeRepoFile(t, root, "feature.txt", "feature\n")
	runGit(t, root, "add", ".")
	runGit(t, root, "commit", "-m", "feature")
	runGit(t, root, "checkout", "base")
	runGit(t, root, "merge", "--no-ff", "feature",
		"-m", "Merge pull request #1 from kimsemi-home/feature")
	return root
}

func writeRepoFile(t *testing.T, root string, rel string, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}
}

func runGit(t *testing.T, root string, args ...string) {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git %v failed: %v\n%s", args, err, out)
	}
}

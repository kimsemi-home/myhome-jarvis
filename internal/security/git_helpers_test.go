package security

import (
	"os/exec"
	"strings"
	"testing"
)

func initGitRepo(t *testing.T) string {
	t.Helper()
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git is not installed")
	}
	root := t.TempDir()
	runGit(t, root, "init")
	runGit(t, root, "config", "user.name", "kimsemi-home")
	runGit(t, root, "config", "user.email", "293568138+kimsemi-home@users.noreply.github.com")
	return root
}

func commitAll(t *testing.T, root string, message string) {
	t.Helper()
	runGit(t, root, "add", ".")
	runGit(t, root, "commit", "-m", message)
}

func runGit(t *testing.T, root string, args ...string) {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %s failed: %v\n%s", strings.Join(args, " "), err, output)
	}
}

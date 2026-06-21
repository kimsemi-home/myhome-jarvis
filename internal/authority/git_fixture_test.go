package authority

import (
	"os/exec"
	"strings"
	"testing"
)

func commitPublicFixture(t *testing.T, root string) {
	t.Helper()
	for _, args := range [][]string{
		{"init"},
		{"config", "user.name", "kimsemi-home"},
		{"config", "user.email", "293568138+kimsemi-home@users.noreply.github.com"},
		{"add", "."},
		{"commit", "-m", "test fixture"},
	} {
		cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("git %s failed: %v\n%s", strings.Join(args, " "), err, output)
		}
	}
}

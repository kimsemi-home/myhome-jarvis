package security

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCheckRejectsPythonFile(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "script.py"), []byte("print('no')\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	report, err := Check(root)
	if err != nil {
		t.Fatal(err)
	}
	if report.OK {
		t.Fatal("expected Python file to be rejected")
	}
}

func TestCheckAllowsPrivateLocalFiles(t *testing.T) {
	root := t.TempDir()
	privateDir := filepath.Join(root, "data", "private")
	if err := os.MkdirAll(privateDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(privateDir, "linear-token.txt"), []byte("local\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	report, err := Check(root)
	if err != nil {
		t.Fatal(err)
	}
	if !report.OK {
		t.Fatalf("expected private local files to be allowed, got %+v", report.Findings)
	}
}

func TestCheckHistoryRejectsPrivateMarkerInCommittedContent(t *testing.T) {
	root := initGitRepo(t)
	privatePath := "/" + "Users" + "/" + strings.Join([]string{"al", "ice"}, "") + "/project"
	if err := os.WriteFile(filepath.Join(root, "notes.md"), []byte("old path: "+privatePath+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	commitAll(t, root, "add private marker")

	report, err := CheckHistory(root)
	if err != nil {
		t.Fatal(err)
	}
	if report.OK {
		t.Fatal("expected private marker in history to be rejected")
	}
	if !historyHasCode(report, "history_private_identity") {
		t.Fatalf("expected private identity finding, got %+v", report.Findings)
	}
}

func TestCheckHistoryRejectsPrivateDataPath(t *testing.T) {
	root := initGitRepo(t)
	privateDir := filepath.Join(root, "data", "private")
	if err := os.MkdirAll(privateDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(privateDir, "local-token.txt"), []byte("local only\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	commitAll(t, root, "add private file")

	report, err := CheckHistory(root)
	if err != nil {
		t.Fatal(err)
	}
	if report.OK {
		t.Fatal("expected private data path in history to be rejected")
	}
	if !historyHasCode(report, "history_private_data_path") {
		t.Fatalf("expected private data path finding, got %+v", report.Findings)
	}
}

func TestCheckHistoryAllowsPrivateKeepPlaceholder(t *testing.T) {
	root := initGitRepo(t)
	privateDir := filepath.Join(root, "data", "private")
	if err := os.MkdirAll(privateDir, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(privateDir, ".keep"), []byte(""), 0o600); err != nil {
		t.Fatal(err)
	}
	commitAll(t, root, "add private placeholder")

	report, err := CheckHistory(root)
	if err != nil {
		t.Fatal(err)
	}
	if !report.OK {
		t.Fatalf("expected private placeholder to be allowed, got %+v", report.Findings)
	}
}

func TestCheckHistoryAllowsRedactedTokenPlaceholder(t *testing.T) {
	root := initGitRepo(t)
	body := []byte("request := \"/health?local_token=redacted-value\"\n")
	if err := os.WriteFile(filepath.Join(root, "server_test.go"), body, 0o600); err != nil {
		t.Fatal(err)
	}
	commitAll(t, root, "add redacted placeholder")

	report, err := CheckHistory(root)
	if err != nil {
		t.Fatal(err)
	}
	if !report.OK {
		t.Fatalf("expected redacted token placeholder to be allowed, got %+v", report.Findings)
	}
}

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

func historyHasCode(report HistoryReport, code string) bool {
	for _, finding := range report.Findings {
		if finding.Code == code {
			return true
		}
	}
	return false
}

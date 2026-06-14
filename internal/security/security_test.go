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

func TestCheckReportDoesNotExposeLocalRoot(t *testing.T) {
	root := t.TempDir()
	report, err := Check(root)
	if err != nil {
		t.Fatal(err)
	}
	if report.Root != "." {
		t.Fatalf("root = %q, expected redacted current root", report.Root)
	}
	if strings.Contains(report.Root, root) || filepath.IsAbs(report.Root) {
		t.Fatalf("report leaked local root: %#v", report)
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

func TestCheckRejectsPrivateMarkerInCurrentContent(t *testing.T) {
	root := t.TempDir()
	privatePath := "/" + "Users" + "/" + strings.Join([]string{"al", "ice"}, "") + "/project"
	if err := os.WriteFile(filepath.Join(root, "notes.md"), []byte("old path: "+privatePath+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	report, err := Check(root)
	if err != nil {
		t.Fatal(err)
	}
	if report.OK {
		t.Fatal("expected current private marker to be rejected")
	}
	if !currentHasCode(report, "current_private_identity") {
		t.Fatalf("expected private identity finding, got %+v", report.Findings)
	}
	if strings.Contains(report.Findings[0].Message, privatePath) {
		t.Fatalf("finding leaked matched content: %+v", report.Findings[0])
	}
}

func TestCheckRejectsSecretLiteralInCurrentContent(t *testing.T) {
	root := t.TempDir()
	key := strings.Join([]string{"api", "_", "key"}, "")
	value := strings.Repeat("a", 24)
	if err := os.WriteFile(filepath.Join(root, "config.md"), []byte(key+" = "+value+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	report, err := Check(root)
	if err != nil {
		t.Fatal(err)
	}
	if report.OK {
		t.Fatal("expected current secret-looking literal to be rejected")
	}
	if !currentHasCode(report, "current_secret_literal") {
		t.Fatalf("expected secret literal finding, got %+v", report.Findings)
	}
	for _, finding := range report.Findings {
		if strings.Contains(finding.Message, value) {
			t.Fatalf("finding leaked matched secret content: %+v", finding)
		}
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

func TestStatusForRootAggregatesCurrentAndHistoryWithoutRoot(t *testing.T) {
	root := initGitRepo(t)
	if err := os.WriteFile(filepath.Join(root, "README.md"), []byte("ok\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	commitAll(t, root, "initial")

	status, err := StatusForRoot(root)
	if err != nil {
		t.Fatal(err)
	}
	if !status.OK || !status.CurrentOK || !status.HistoryOK {
		t.Fatalf("expected clean status, got %#v", status)
	}
	if status.CurrentFindingCount != 0 || status.HistoryFindingCount != 0 {
		t.Fatalf("unexpected finding counts: %#v", status)
	}
	if status.CheckedAt == "" {
		t.Fatal("expected checked_at to be set")
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

func currentHasCode(report Report, code string) bool {
	for _, finding := range report.Findings {
		if finding.Code == code {
			return true
		}
	}
	return false
}

package main

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
)

func runQuality(root string) error {
	started := time.Now()
	report := qualityReport{OK: true}
	goTool := envWithDefault("MHJ_GO", "go")
	gofmtTool := envWithDefault("MHJ_GOFMT", "gofmt")

	securityReport, err := security.Check(root)
	if err != nil {
		return err
	}
	if securityReport.OK {
		report.Steps = append(report.Steps, qualityStep{Name: "security check", Status: "pass"})
	} else {
		report.OK = false
		report.Steps = append(report.Steps, qualityStep{Name: "security check", Status: "fail", Output: "security findings present"})
	}
	securityHistoryReport, err := security.CheckHistory(root)
	if err != nil {
		return err
	}
	if securityHistoryReport.OK {
		report.Steps = append(report.Steps, qualityStep{Name: "security history", Status: "pass"})
	} else {
		report.OK = false
		report.Steps = append(report.Steps, qualityStep{Name: "security history", Status: "fail", Output: "security history findings present"})
	}
	report.addCheck("toolchain pins", validateToolchainPins(root))
	report.addCheck("ci workflow", validateCIWorkflowContract(root))
	report.addCommand(root, "code shape", []string{goTool, "run", "./cmd/mhj", "code-shape", "status"})

	report.addHarness("home harness", commands.RunHomeHarness())
	report.addHarness("finance harness", commands.RunFinanceHarness(root))
	report.addHarness("commerce harness", commands.RunCommerceHarness(root))

	report.addCommand(root, "go test", []string{goTool, "test", "./..."})
	report.addCommand(root, "go vet", []string{goTool, "vet", "./..."})
	report.addGofmt(root, gofmtTool)
	report.addCommand(root, "cargo test", []string{"cargo", "test", "--workspace"})
	report.addCommand(root, "benchmark smoke", benchmarkSmokeCommand())
	report.addCommand(root, "cargo fmt", []string{"cargo", "fmt", "--check"})
	report.addCommand(root, "cargo clippy", []string{"cargo", "clippy", "--workspace", "--", "-D", "warnings"})
	report.addCommand(root, "ssot validate", []string{"sbcl", "--script", "lisp/scripts/validate-ssot.lisp"})
	report.addCommand(root, "codegen verify", []string{goTool, "run", "./cmd/mhj", "codegen", "verify"})
	report.addCommand(root, "ddd verify", []string{goTool, "run", "./cmd/mhj", "ddd", "verify"})
	if _, err := os.Stat(filepath.Join(root, "apps", "flutter", "pubspec.yaml")); err == nil {
		flutterRoot := filepath.Join(root, "apps", "flutter")
		report.addCommand(flutterRoot, "flutter test", []string{"flutter", "test"})
		report.addCommand(flutterRoot, "flutter analyze", []string{"flutter", "analyze"})
	} else {
		report.Steps = append(report.Steps, qualityStep{Name: "flutter", Status: "skip", Output: "apps/flutter is not started yet"})
	}
	if err := qualitylog.AppendRun(root, qualitylog.NewRun(started, report.OK, qualityEvidenceSteps(report.Steps))); err != nil {
		return err
	}
	if err := writeJSON(report); err != nil {
		return err
	}
	if !report.OK {
		return errors.New("quality gate failed")
	}
	return nil
}

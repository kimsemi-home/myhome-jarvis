package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/agentcluster"
	"github.com/kimsemi-home/myhome-jarvis/internal/audit"
	"github.com/kimsemi-home/myhome-jarvis/internal/auth"
	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
	"github.com/kimsemi-home/myhome-jarvis/internal/connectors"
	"github.com/kimsemi-home/myhome-jarvis/internal/daemon"
	"github.com/kimsemi-home/myhome-jarvis/internal/evidence"
	"github.com/kimsemi-home/myhome-jarvis/internal/knowledge"
	"github.com/kimsemi-home/myhome-jarvis/internal/learning"
	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
	"github.com/kimsemi-home/myhome-jarvis/internal/orchestrator"
	"github.com/kimsemi-home/myhome-jarvis/internal/planner"
	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
	"github.com/kimsemi-home/myhome-jarvis/internal/repo"
	"github.com/kimsemi-home/myhome-jarvis/internal/scheduler"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
	"github.com/kimsemi-home/myhome-jarvis/internal/supervisor"
)

const version = "0.1.0-bootstrap"

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return usage()
	}
	root, err := os.Getwd()
	if err != nil {
		return err
	}

	switch args[0] {
	case "version":
		fmt.Println("myhome-jarvis " + version)
		return nil
	case "commands":
		return writeJSON(commands.Specs())
	case "auth":
		return runAuth(root, args[1:])
	case "audit":
		if len(args) == 2 && args[1] == "status" {
			return auditStatus(root)
		}
	case "ci":
		if len(args) == 2 && args[1] == "verify" {
			return runCIVerify(root)
		}
	case "security":
		if len(args) == 2 && args[1] == "check" {
			report, err := security.Check(root)
			if err != nil {
				return err
			}
			if err := writeJSON(report); err != nil {
				return err
			}
			if !report.OK {
				return errors.New("security check failed")
			}
			return nil
		}
		if len(args) == 2 && args[1] == "history" {
			report, err := security.CheckHistory(root)
			if err != nil {
				return err
			}
			if err := writeJSON(report); err != nil {
				return err
			}
			if !report.OK {
				return errors.New("security history check failed")
			}
			return nil
		}
	case "command":
		if len(args) != 3 {
			return errors.New("usage: mhj command <name> '<json-payload>'")
		}
		executeRequested := os.Getenv("MYHOME_EXECUTE") == "true"
		plan, err := commands.Build(args[1], []byte(args[2]))
		if err == nil {
			plan = commands.WithExecuteAllowed(plan, executeRequested)
		}
		if plan.ExecuteAllowed {
			plan, err = commands.Execute(context.Background(), plan, commands.ExecuteOptions{})
		}
		if auditErr := audit.AppendCommandIntent(root, audit.CommandIntentFromPlan("cli", args[1], executeRequested, plan, err)); err == nil && auditErr != nil {
			return auditErr
		}
		if err != nil {
			return err
		}
		return writeJSON(plan)
	case "connectors":
		if len(args) == 2 && args[1] == "status" {
			return connectorsStatus(root)
		}
	case "agent-cluster":
		if len(args) == 2 && args[1] == "status" {
			return agentClusterStatus(root)
		}
	case "learning":
		return runLearning(root, args[1:])
	case "evidence":
		if len(args) == 2 && args[1] == "status" {
			return evidenceStatus(root)
		}
	case "harness":
		return runHarness(root, args[1:])
	case "toolchain":
		if len(args) == 2 && args[1] == "verify" {
			return runToolchainVerify(root)
		}
	case "linear":
		if len(args) == 2 && args[1] == "status" {
			return writeJSON(linear.SummarizeStatus(linear.CurrentStatus(root)))
		}
		if len(args) == 2 && args[1] == "sync" {
			result := linear.PullIssues(context.Background(), root, http.DefaultClient)
			if !result.Synced {
				if err := linear.AppendOfflineEvent(root, "linear_sync", result.Message); err != nil {
					return err
				}
			}
			return writeJSON(linear.SummarizeOperation(result))
		}
		if len(args) == 2 && args[1] == "pull" {
			result := linear.PullIssues(context.Background(), root, http.DefaultClient)
			if !result.Synced {
				if err := linear.AppendOfflineEvent(root, "linear_pull", result.Message); err != nil {
					return err
				}
			}
			return writeJSON(linear.SummarizeOperation(result))
		}
		if len(args) == 2 && args[1] == "next" {
			result := linear.NextIssue(context.Background(), root, http.DefaultClient)
			if !result.Synced {
				if err := linear.AppendOfflineEvent(root, "linear_next", result.Message); err != nil {
					return err
				}
			}
			return writeJSON(linear.SummarizeOperation(result))
		}
		if len(args) >= 4 && args[1] == "comment" {
			result := linear.AddComment(context.Background(), root, http.DefaultClient, args[2], strings.Join(args[3:], " "))
			return writeJSON(linear.SummarizeOperation(result))
		}
		if len(args) >= 4 && args[1] == "transition" {
			result := linear.TransitionIssue(context.Background(), root, http.DefaultClient, args[2], strings.Join(args[3:], " "))
			return writeJSON(linear.SummarizeOperation(result))
		}
		if len(args) == 2 && args[1] == "create-from-backlog" {
			result := linear.CreateFromBacklog(context.Background(), root, http.DefaultClient)
			return writeJSON(linear.SummarizeOperation(result))
		}
		if len(args) == 2 && args[1] == "replay-offline" {
			return writeJSON(linear.ReplayOffline(context.Background(), root, http.DefaultClient))
		}
	case "daemon":
		return runDaemon(root, args[1:])
	case "ddd":
		if len(args) == 2 && args[1] == "verify" {
			return runDDDVerify(root)
		}
	case "knowledge":
		return runKnowledge(root, args[1:])
	case "repo":
		if len(args) == 2 && args[1] == "status" {
			return repoStatus(root)
		}
	case "planner":
		if len(args) == 2 && args[1] == "status" {
			return plannerStatus(root)
		}
	case "loop":
		if len(args) == 2 && args[1] == "once" {
			return loopOnce(root)
		}
		if len(args) == 2 && args[1] == "status" {
			return loopStatus(root)
		}
		if len(args) >= 2 && args[1] == "worker" {
			return loopWorker(root, args[2:])
		}
	case "benchmark":
		if len(args) == 2 && args[1] == "smoke" {
			return runBenchmarkSmoke(root)
		}
	case "quality":
		if len(args) == 2 && args[1] == "status" {
			return qualityStatus(root)
		}
		return runQuality(root)
	case "codegen":
		if len(args) == 2 && args[1] == "verify" {
			return runCodegenVerify(root)
		}
		return runCodegen(root)
	}
	return usage()
}

func usage() error {
	return errors.New("usage: mhj <version|commands|auth status|auth token create|auth token rotate|audit status|ci verify|security check|security history|command|connectors status|agent-cluster status|learning status|learning record|evidence status|harness home|harness finance|harness commerce|toolchain verify|linear status|linear sync|linear pull|linear next|linear comment|linear transition|linear create-from-backlog|linear replay-offline|daemon|daemon status|ddd verify|knowledge verify|knowledge search|repo status|planner status|loop once|loop status|loop worker|benchmark smoke|quality|quality status|codegen|codegen verify>")
}

func runAuth(root string, args []string) error {
	if len(args) == 1 && args[0] == "status" {
		return writeJSON(auth.Status(root))
	}
	if len(args) == 2 && args[0] == "token" && args[1] == "create" {
		result, err := auth.Create(root, false)
		if err != nil {
			return err
		}
		return writeJSON(result)
	}
	if len(args) == 2 && args[0] == "token" && args[1] == "rotate" {
		result, err := auth.Create(root, true)
		if err != nil {
			return err
		}
		return writeJSON(result)
	}
	return errors.New("usage: mhj auth <status|token create|token rotate>")
}

func runDaemon(root string, args []string) error {
	if len(args) == 1 && args[0] == "status" {
		return writeJSON(supervisor.Status(root, nil))
	}
	config := daemon.DefaultConfig(root, version)
	flags := flag.NewFlagSet("daemon", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	flags.StringVar(&config.Host, "host", config.Host, "bind host")
	flags.IntVar(&config.Port, "port", config.Port, "bind port")
	flags.BoolVar(&config.AllowLANBind, "allow-lan", false, "allow non-localhost bind")
	flags.BoolVar(&config.Execute, "execute", config.Execute, "allow explicit execute requests")
	if err := flags.Parse(args); err != nil {
		return err
	}
	server, err := daemon.New(config)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "myhome-jarvis daemon listening on %s:%d\n", config.Host, config.Port)
	return server.ListenAndServe()
}

func runHarness(root string, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: mhj harness <home|finance|commerce>")
	}
	var report commands.HarnessReport
	switch args[0] {
	case "home":
		report = commands.RunHomeHarness()
	case "finance":
		report = commands.RunFinanceHarness(root)
	case "commerce":
		report = commands.RunCommerceHarness(root)
	default:
		return errors.New("usage: mhj harness <home|finance|commerce>")
	}
	if err := writeJSON(report); err != nil {
		return err
	}
	if !report.Passed {
		return fmt.Errorf("%s harness failed", report.Name)
	}
	return nil
}

func loopOnce(root string) error {
	linearStatus := linear.CurrentStatus(root)
	linearSummary := linear.SummarizeStatus(linearStatus)
	linearNext := linear.NextIssue(context.Background(), root, http.DefaultClient)
	linearNextSummary := linear.SummarizeOperation(linearNext)
	securityStatus, err := security.StatusForRoot(root)
	if err != nil {
		return err
	}
	plannerStatus, err := planner.StatusForRoot(root)
	if err != nil {
		return err
	}
	if linearStatus.Mode == "offline" {
		if err := linear.AppendOfflineEvent(root, "loop_once", "Local loop ran without Linear sync; synced=false."); err != nil {
			return err
		}
	}
	if !linearNext.Synced {
		if err := linear.AppendOfflineEvent(root, "linear_next", linearNext.Message); err != nil {
			return err
		}
	}
	result := "checkpoint recorded"
	if !securityStatus.OK {
		result = "checkpoint recorded with public-safety findings"
	}
	path, err := orchestrator.WriteCheckpoint(root, orchestrator.Checkpoint{
		Task:           "loop once",
		LinearStatus:   linearSummary,
		LinearNext:     &linearNextSummary,
		PlannerStatus:  plannerStatus,
		SecurityStatus: securityStatus,
		Result:         result,
		Next:           "Continue local-first closed-loop hardening.",
	})
	if err != nil {
		return err
	}
	checkpointPath, err := filepath.Rel(root, path)
	if err != nil {
		return err
	}
	return writeJSON(map[string]any{
		"ok":              securityStatus.OK,
		"checkpoint":      filepath.ToSlash(checkpointPath),
		"linear":          linearSummary,
		"linear_next":     linearNextSummary,
		"planner_status":  plannerStatus,
		"security_status": securityStatus,
	})
}

func loopStatus(root string) error {
	status, err := scheduler.Status(root, scheduler.ClosedLoopPolicy())
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func repoStatus(root string) error {
	status, err := repo.Inspect(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func plannerStatus(root string) error {
	status, err := planner.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func connectorsStatus(root string) error {
	status, err := connectors.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func agentClusterStatus(root string) error {
	status, err := agentcluster.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func auditStatus(root string) error {
	status, err := audit.CommandIntentStatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func qualityStatus(root string) error {
	status, err := qualitylog.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func evidenceStatus(root string) error {
	status, err := evidence.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func loopWorker(root string, args []string) error {
	flags := flag.NewFlagSet("loop worker", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	cycles := flags.Int("cycles", 1, "bounded scheduler cycles to run")
	if err := flags.Parse(args); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	status, err := scheduler.RunCycles(ctx, root, scheduler.ClosedLoopPolicy(), *cycles, func(context.Context) (scheduler.JobResult, error) {
		linearStatus := linear.CurrentStatus(root)
		linearSummary := linear.SummarizeStatus(linearStatus)
		linearNext := linear.NextIssue(ctx, root, http.DefaultClient)
		linearNextSummary := linear.SummarizeOperation(linearNext)
		securityStatus, err := security.StatusForRoot(root)
		if err != nil {
			return scheduler.JobResult{}, err
		}
		plannerStatus, err := planner.StatusForRoot(root)
		if err != nil {
			return scheduler.JobResult{}, err
		}
		if !linearNext.Synced {
			if err := linear.AppendOfflineEvent(root, "linear_next", linearNext.Message); err != nil {
				return scheduler.JobResult{}, err
			}
		}
		result := "scheduler heartbeat checkpoint recorded"
		if !securityStatus.OK {
			result = "scheduler heartbeat checkpoint recorded with public-safety findings"
		}
		path, err := orchestrator.WriteCheckpoint(root, orchestrator.Checkpoint{
			Task:           "loop worker",
			LinearStatus:   linearSummary,
			LinearNext:     &linearNextSummary,
			PlannerStatus:  plannerStatus,
			SecurityStatus: securityStatus,
			Result:         result,
			Next:           "Continue local-first fixture and daemon surface expansion.",
		})
		if err != nil {
			return scheduler.JobResult{}, err
		}
		return scheduler.JobResult{Checkpoint: path}, nil
	})
	if err != nil {
		return err
	}
	return writeJSON(status)
}

type qualityStep struct {
	Name    string   `json:"name"`
	Status  string   `json:"status"`
	Command []string `json:"-"`
	Output  string   `json:"-"`
}

type qualityReport struct {
	OK    bool          `json:"ok"`
	Steps []qualityStep `json:"steps"`
}

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

func qualityEvidenceSteps(steps []qualityStep) []qualitylog.Step {
	evidence := make([]qualitylog.Step, 0, len(steps))
	for _, step := range steps {
		evidence = append(evidence, qualitylog.Step{Name: step.Name, Status: step.Status})
	}
	return evidence
}

func (report *qualityReport) addHarness(name string, harness commands.HarnessReport) {
	if harness.Passed {
		report.Steps = append(report.Steps, qualityStep{Name: name, Status: "pass"})
		return
	}
	report.OK = false
	report.Steps = append(report.Steps, qualityStep{Name: name, Status: "fail"})
}

func runToolchainVerify(root string) error {
	if err := validateToolchainPins(root); err != nil {
		return err
	}
	return writeJSON(map[string]any{"ok": true})
}

func runCIVerify(root string) error {
	if err := validateCIWorkflowContract(root); err != nil {
		return err
	}
	return writeJSON(map[string]any{"ok": true})
}

func runDDDVerify(root string) error {
	report, err := knowledge.Verify(root)
	if err != nil {
		return err
	}
	if err := writeJSON(report); err != nil {
		return err
	}
	if !report.OK {
		return errors.New("ddd verify failed")
	}
	return nil
}

func runKnowledge(root string, args []string) error {
	if len(args) == 1 && args[0] == "verify" {
		return runDDDVerify(root)
	}
	if len(args) >= 2 && args[0] == "search" {
		report, err := knowledge.Search(root, strings.Join(args[1:], " "))
		if err != nil {
			return err
		}
		return writeJSON(report)
	}
	return errors.New("usage: mhj knowledge <verify|search query>")
}

func runLearning(root string, args []string) error {
	if len(args) == 1 && args[0] == "status" {
		status, err := learning.StatusForRoot(root)
		if err != nil {
			return err
		}
		return writeJSON(status)
	}
	if len(args) == 2 && args[0] == "record" {
		result, err := learning.Record(root, []byte(args[1]))
		if err != nil {
			return err
		}
		return writeJSON(result)
	}
	return errors.New("usage: mhj learning <status|record json-payload>")
}

func (report *qualityReport) addCheck(name string, err error) {
	if err == nil {
		report.Steps = append(report.Steps, qualityStep{Name: name, Status: "pass"})
		return
	}
	report.OK = false
	report.Steps = append(report.Steps, qualityStep{Name: name, Status: "fail", Output: err.Error()})
}

func runBenchmarkSmoke(root string) error {
	report := qualityReport{OK: true}
	report.addCommand(root, "benchmark smoke", benchmarkSmokeCommand())
	if err := writeJSON(report); err != nil {
		return err
	}
	if !report.OK {
		return errors.New("benchmark smoke failed")
	}
	return nil
}

func benchmarkSmokeCommand() []string {
	return []string{"cargo", "test", "-p", "mhj-core", "benchmark_smoke", "--", "--nocapture"}
}

func (report *qualityReport) addCommand(root string, name string, command []string) {
	step := qualityStep{Name: name, Command: command}
	if _, err := exec.LookPath(command[0]); err != nil {
		step.Status = "fail"
		step.Output = "missing executable: " + command[0]
		report.OK = false
		report.Steps = append(report.Steps, step)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	cmd := exec.CommandContext(ctx, command[0], command[1:]...)
	cmd.Dir = root
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Run(); err != nil {
		step.Status = "fail"
		step.Output = strings.TrimSpace(output.String())
		report.OK = false
	} else {
		step.Status = "pass"
		step.Output = strings.TrimSpace(output.String())
	}
	report.Steps = append(report.Steps, step)
}

func (report *qualityReport) addGofmt(root string, gofmtTool string) {
	files, err := collectGoFiles(root)
	if err != nil {
		report.OK = false
		report.Steps = append(report.Steps, qualityStep{Name: "gofmt", Status: "fail", Output: err.Error()})
		return
	}
	if len(files) == 0 {
		report.Steps = append(report.Steps, qualityStep{Name: "gofmt", Status: "skip", Output: "no Go files"})
		return
	}
	command := append([]string{gofmtTool, "-l"}, files...)
	report.addCommand(root, "gofmt", command)
	last := &report.Steps[len(report.Steps)-1]
	if last.Status == "pass" && strings.TrimSpace(last.Output) != "" {
		last.Status = "fail"
		last.Output = "unformatted files:\n" + last.Output
		report.OK = false
	}
}

func collectGoFiles(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			switch entry.Name() {
			case ".git", "target", "build", "dist", "bin":
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		files = append(files, filepath.ToSlash(rel))
		return nil
	})
	return files, err
}

func envWithDefault(name string, fallback string) string {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return fallback
	}
	return value
}

func validateToolchainPins(root string) error {
	goVersion, err := readTrimmedFile(root, ".go-version")
	if err != nil {
		return err
	}
	goModVersion, err := parseFirstMatchFile(root, "go.mod", `(?m)^go\s+([0-9]+\.[0-9]+\.[0-9]+)\s*$`)
	if err := requireEqual("go.mod go directive", goVersion, goModVersion, err); err != nil {
		return err
	}
	generatedGo, err := generatedProjectGoVersion(root)
	if err != nil {
		return err
	}
	if err := requireEqualValue("generated project go_version", goVersion, generatedGo); err != nil {
		return err
	}
	workflowGoVersion, err := parseWorkflowEnv(root, "GO_VERSION")
	if err := requireEqual("workflow GO_VERSION", goVersion, workflowGoVersion, err); err != nil {
		return err
	}
	rustVersion, err := parseFirstMatchFile(root, "rust-toolchain.toml", `(?m)^channel\s*=\s*"([^"]+)"\s*$`)
	if err != nil {
		return err
	}
	workflowRustVersion, err := parseWorkflowEnv(root, "RUST_TOOLCHAIN")
	if err := requireEqual("workflow RUST_TOOLCHAIN", rustVersion, workflowRustVersion, err); err != nil {
		return err
	}
	return nil
}

func validateCIWorkflowContract(root string) error {
	body, err := os.ReadFile(filepath.Join(root, ".github", "workflows", "quality.yml"))
	if err != nil {
		return err
	}
	workflow := string(body)
	required := []string{
		"cancel-in-progress: true",
		"permissions:",
		"contents: read",
		"fetch-depth: 0",
		"go run ./cmd/mhj security check",
		"go run ./cmd/mhj security history",
		"go run ./cmd/mhj ci verify",
		"go run ./cmd/mhj toolchain verify",
		"'.go-version'",
		"'rust-toolchain.toml'",
		"'generated/*.json'",
		"'generated/commands.generated.json'",
		"'generated/connectors.generated.json'",
		"'generated/agent_cluster.generated.json'",
		"'generated/learning.generated.json'",
		"'generated/evidence.generated.json'",
		"github.event_name == 'push' && github.repository == 'kimsemi-home/myhome-jarvis'",
	}
	for _, token := range required {
		if !strings.Contains(workflow, token) {
			return fmt.Errorf("quality workflow missing CI contract token %q", token)
		}
	}
	forbidden := []string{
		"pull_request_target",
		"write-all",
	}
	for _, token := range forbidden {
		if strings.Contains(workflow, token) {
			return fmt.Errorf("quality workflow contains forbidden CI contract token %q", token)
		}
	}
	writePermissionPattern := regexp.MustCompile(`(?m)^\s*[A-Za-z0-9_-]+:\s*write\s*$`)
	if match := writePermissionPattern.FindString(workflow); match != "" {
		return fmt.Errorf("quality workflow contains forbidden write permission %q", strings.TrimSpace(match))
	}
	return nil
}

func readTrimmedFile(root string, rel string) (string, error) {
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		return "", err
	}
	value := strings.TrimSpace(string(body))
	if value == "" {
		return "", fmt.Errorf("%s is empty", rel)
	}
	return value, nil
}

func parseFirstMatchFile(root string, rel string, pattern string) (string, error) {
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(string(body))
	if len(matches) < 2 {
		return "", fmt.Errorf("%s does not match expected toolchain pattern", rel)
	}
	return strings.TrimSpace(matches[1]), nil
}

func parseWorkflowEnv(root string, name string) (string, error) {
	pattern := fmt.Sprintf(`(?m)^\s+%s:\s*"([^"]+)"\s*$`, regexp.QuoteMeta(name))
	return parseFirstMatchFile(root, ".github/workflows/quality.yml", pattern)
}

func generatedProjectGoVersion(root string) (string, error) {
	body, err := os.ReadFile(filepath.Join(root, "generated", "commands.generated.json"))
	if err != nil {
		return "", err
	}
	var catalog struct {
		Project struct {
			GoVersion string `json:"go_version"`
		} `json:"project"`
	}
	if err := json.Unmarshal(body, &catalog); err != nil {
		return "", err
	}
	if strings.TrimSpace(catalog.Project.GoVersion) == "" {
		return "", errors.New("generated project go_version is empty")
	}
	return strings.TrimSpace(catalog.Project.GoVersion), nil
}

func requireEqual(label string, expected string, actual string, err error) error {
	if err != nil {
		return err
	}
	return requireEqualValue(label, expected, actual)
}

func requireEqualValue(label string, expected string, actual string) error {
	if actual != expected {
		return fmt.Errorf("%s = %q, expected %q", label, actual, expected)
	}
	return nil
}

func runCodegen(root string) error {
	if _, err := exec.LookPath("sbcl"); err != nil {
		return errors.New("missing executable: sbcl")
	}
	cmd := exec.Command("sbcl", "--script", "lisp/scripts/codegen.lisp")
	cmd.Dir = root
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runCodegenVerify(root string) error {
	before, err := generatedSnapshot(root)
	if err != nil {
		return err
	}
	if err := runCodegen(root); err != nil {
		return err
	}
	after, err := generatedSnapshot(root)
	if err != nil {
		return err
	}
	changed := changedGeneratedFiles(before, after)
	if len(changed) > 0 {
		return fmt.Errorf("generated artifacts are out of date: %s", strings.Join(changed, ", "))
	}
	fmt.Fprintln(os.Stdout, "Generated artifacts verified")
	return nil
}

func generatedSnapshot(root string) (map[string][]byte, error) {
	generatedRoot := filepath.Join(root, "generated")
	files := map[string][]byte{}
	err := filepath.WalkDir(generatedRoot, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		body, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		files[filepath.ToSlash(rel)] = body
		return nil
	})
	return files, err
}

func changedGeneratedFiles(before map[string][]byte, after map[string][]byte) []string {
	seen := map[string]bool{}
	var changed []string
	for path, body := range before {
		seen[path] = true
		if next, ok := after[path]; !ok || !bytes.Equal(body, next) {
			changed = append(changed, path)
		}
	}
	for path := range after {
		if !seen[path] {
			changed = append(changed, path)
		}
	}
	sort.Strings(changed)
	return changed
}

func writeJSON(value any) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(value)
}

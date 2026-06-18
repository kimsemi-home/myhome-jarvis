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
	"strings"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/audit"
	"github.com/kimsemi-home/myhome-jarvis/internal/auth"
	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
	"github.com/kimsemi-home/myhome-jarvis/internal/daemon"
	"github.com/kimsemi-home/myhome-jarvis/internal/knowledge"
	"github.com/kimsemi-home/myhome-jarvis/internal/learning"
	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
	"github.com/kimsemi-home/myhome-jarvis/internal/qualitylog"
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
	case "code-shape":
		if len(args) == 2 && args[1] == "status" {
			return codeShapeStatus(root)
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
	case "confidence":
		if len(args) == 2 && args[1] == "status" {
			return confidenceStatus(root)
		}
	case "translation":
		if len(args) == 2 && args[1] == "status" {
			return translationStatus(root)
		}
	case "control-plane":
		if len(args) == 2 && args[1] == "status" {
			return controlPlaneStatus(root)
		}
	case "incidents":
		if len(args) == 2 && args[1] == "status" {
			return incidentsStatus(root)
		}
	case "evidence-quality":
		if len(args) == 2 && args[1] == "status" {
			return evidenceQualityStatus(root)
		}
	case "review":
		if len(args) == 2 && args[1] == "status" {
			return reviewStatus(root)
		}
	case "authority":
		if len(args) == 2 && args[1] == "status" {
			return authorityStatus(root)
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
	return errors.New("usage: mhj <version|commands|auth status|auth token create|auth token rotate|audit status|ci verify|code-shape status|security check|security history|command|connectors status|agent-cluster status|learning status|learning record|evidence status|confidence status|translation status|control-plane status|incidents status|evidence-quality status|review status|authority status|harness home|harness finance|harness commerce|toolchain verify|linear status|linear sync|linear pull|linear next|linear comment|linear transition|linear create-from-backlog|linear replay-offline|daemon|daemon status|ddd verify|knowledge verify|knowledge search|repo status|planner status|loop once|loop status|loop worker|benchmark smoke|quality|quality status|codegen|codegen verify>")
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

func writeJSON(value any) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(value)
}

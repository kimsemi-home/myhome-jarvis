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

	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
	"github.com/kimsemi-home/myhome-jarvis/internal/daemon"
	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
	"github.com/kimsemi-home/myhome-jarvis/internal/orchestrator"
	"github.com/kimsemi-home/myhome-jarvis/internal/repo"
	"github.com/kimsemi-home/myhome-jarvis/internal/scheduler"
	"github.com/kimsemi-home/myhome-jarvis/internal/security"
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
	case "command":
		if len(args) != 3 {
			return errors.New("usage: mhj command <name> '<json-payload>'")
		}
		plan, err := commands.Build(args[1], []byte(args[2]))
		if err != nil {
			return err
		}
		plan = commands.WithExecuteAllowed(plan, os.Getenv("MYHOME_EXECUTE") == "true")
		if plan.ExecuteAllowed {
			plan, err = commands.Execute(context.Background(), plan, commands.ExecuteOptions{})
			if err != nil {
				return err
			}
		}
		return writeJSON(plan)
	case "harness":
		if len(args) == 2 && args[1] == "home" {
			report := commands.RunHomeHarness()
			if err := writeJSON(report); err != nil {
				return err
			}
			if !report.Passed {
				return errors.New("home harness failed")
			}
			return nil
		}
	case "linear":
		if len(args) == 2 && args[1] == "status" {
			return writeJSON(linear.CurrentStatus(root))
		}
		if len(args) == 2 && args[1] == "sync" {
			result := linear.PullIssues(context.Background(), root, http.DefaultClient)
			if !result.Synced {
				if err := linear.AppendOfflineEvent(root, "linear_sync", result.Message); err != nil {
					return err
				}
			}
			return writeJSON(result)
		}
		if len(args) == 2 && args[1] == "pull" {
			result := linear.PullIssues(context.Background(), root, http.DefaultClient)
			if !result.Synced {
				if err := linear.AppendOfflineEvent(root, "linear_pull", result.Message); err != nil {
					return err
				}
			}
			return writeJSON(result)
		}
		if len(args) == 2 && args[1] == "next" {
			result := linear.NextIssue(context.Background(), root, http.DefaultClient)
			if !result.Synced {
				if err := linear.AppendOfflineEvent(root, "linear_next", result.Message); err != nil {
					return err
				}
			}
			return writeJSON(result)
		}
		if len(args) >= 4 && args[1] == "comment" {
			return writeJSON(linear.AddComment(context.Background(), root, http.DefaultClient, args[2], strings.Join(args[3:], " ")))
		}
		if len(args) >= 4 && args[1] == "transition" {
			return writeJSON(linear.TransitionIssue(context.Background(), root, http.DefaultClient, args[2], strings.Join(args[3:], " ")))
		}
		if len(args) == 2 && args[1] == "create-from-backlog" {
			return writeJSON(linear.CreateFromBacklog(context.Background(), root, http.DefaultClient))
		}
	case "daemon":
		return runDaemon(root, args[1:])
	case "repo":
		if len(args) == 2 && args[1] == "status" {
			return repoStatus(root)
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
	return errors.New("usage: mhj <version|commands|security check|command|harness home|linear status|linear sync|linear pull|linear next|linear comment|linear transition|linear create-from-backlog|daemon|repo status|loop once|loop status|loop worker|benchmark smoke|quality|codegen|codegen verify>")
}

func runDaemon(root string, args []string) error {
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

func loopOnce(root string) error {
	linearStatus := linear.CurrentStatus(root)
	securityReport, err := security.Check(root)
	if err != nil {
		return err
	}
	if linearStatus.Mode == "offline" {
		if err := linear.AppendOfflineEvent(root, "loop_once", "Local loop ran without Linear sync; synced=false."); err != nil {
			return err
		}
	}
	result := "checkpoint recorded"
	if !securityReport.OK {
		result = "checkpoint recorded with security findings"
	}
	path, err := orchestrator.WriteCheckpoint(root, orchestrator.Checkpoint{
		Task:           "loop once",
		LinearStatus:   linearStatus,
		SecurityReport: securityReport,
		Result:         result,
		Next:           "Implement direct Go GraphQL client and expand quality gate.",
	})
	if err != nil {
		return err
	}
	return writeJSON(map[string]any{
		"ok":         securityReport.OK,
		"checkpoint": filepath.ToSlash(path),
		"linear":     linearStatus,
		"security":   securityReport,
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
		securityReport, err := security.Check(root)
		if err != nil {
			return scheduler.JobResult{}, err
		}
		path, err := orchestrator.WriteCheckpoint(root, orchestrator.Checkpoint{
			Task:           "loop worker",
			LinearStatus:   linearStatus,
			SecurityReport: securityReport,
			Result:         "scheduler heartbeat checkpoint recorded",
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
	Command []string `json:"command,omitempty"`
	Output  string   `json:"output,omitempty"`
}

type qualityReport struct {
	OK    bool          `json:"ok"`
	Steps []qualityStep `json:"steps"`
}

func runQuality(root string) error {
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

	homeHarness := commands.RunHomeHarness()
	if homeHarness.Passed {
		report.Steps = append(report.Steps, qualityStep{Name: "home harness", Status: "pass"})
	} else {
		report.OK = false
		report.Steps = append(report.Steps, qualityStep{Name: "home harness", Status: "fail"})
	}

	report.addCommand(root, "go test", []string{goTool, "test", "./..."})
	report.addCommand(root, "go vet", []string{goTool, "vet", "./..."})
	report.addGofmt(root, gofmtTool)
	report.addCommand(root, "cargo test", []string{"cargo", "test", "--workspace"})
	report.addCommand(root, "benchmark smoke", benchmarkSmokeCommand())
	report.addCommand(root, "cargo fmt", []string{"cargo", "fmt", "--check"})
	report.addCommand(root, "cargo clippy", []string{"cargo", "clippy", "--workspace", "--", "-D", "warnings"})
	report.addCommand(root, "ssot validate", []string{"sbcl", "--script", "lisp/scripts/validate-ssot.lisp"})
	report.addCommand(root, "ssot codegen", []string{"sbcl", "--script", "lisp/scripts/codegen.lisp"})
	if _, err := os.Stat(filepath.Join(root, "apps", "flutter", "pubspec.yaml")); err == nil {
		flutterRoot := filepath.Join(root, "apps", "flutter")
		report.addCommand(flutterRoot, "flutter test", []string{"flutter", "test"})
		report.addCommand(flutterRoot, "flutter analyze", []string{"flutter", "analyze"})
	} else {
		report.Steps = append(report.Steps, qualityStep{Name: "flutter", Status: "skip", Output: "apps/flutter is not started yet"})
	}

	if err := writeJSON(report); err != nil {
		return err
	}
	if !report.OK {
		return errors.New("quality gate failed")
	}
	return nil
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
	if err := runCodegen(root); err != nil {
		return err
	}
	if _, err := exec.LookPath("git"); err != nil {
		return errors.New("missing executable: git")
	}
	cmd := exec.Command("git", "diff", "--exit-code", "--", "generated")
	cmd.Dir = root
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := cmd.Run(); err != nil {
		trimmed := strings.TrimSpace(output.String())
		if trimmed == "" {
			trimmed = "generated artifacts differ from SSOT"
		}
		return fmt.Errorf("generated artifacts are out of date:\n%s", trimmed)
	}
	fmt.Fprintln(os.Stdout, "Generated artifacts verified")
	return nil
}

func writeJSON(value any) error {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	return encoder.Encode(value)
}

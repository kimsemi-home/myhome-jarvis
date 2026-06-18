package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/kimsemi-home/myhome-jarvis/internal/audit"
	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
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

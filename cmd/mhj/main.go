package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
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
		return runSecurity(root, args[1:])
	case "command":
		return runCommand(root, args[1:])
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
		return runLinear(root, args[1:])
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

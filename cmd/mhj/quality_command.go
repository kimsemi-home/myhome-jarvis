package main

import (
	"bytes"
	"context"
	"errors"
	"os/exec"
	"strings"
	"time"
)

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

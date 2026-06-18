package main

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/scheduler"
)

func loopStatus(root string) error {
	status, err := scheduler.Status(root, scheduler.ClosedLoopPolicy())
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
		return writeLoopWorkerCycle(ctx, root)
	})
	if err != nil {
		return err
	}
	return writeJSON(status)
}

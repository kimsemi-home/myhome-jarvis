package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/kimsemi-home/myhome-jarvis/internal/linear"
)

type linearOperation func(context.Context, string, *http.Client) linear.OperationResult

func runLinear(root string, args []string) error {
	if len(args) == 1 {
		switch args[0] {
		case "status":
			return writeJSON(linear.SummarizeStatus(linear.CurrentStatus(root)))
		case "sync":
			return runLinearOperation(root, "linear_sync", linear.PullIssues)
		case "pull":
			return runLinearOperation(root, "linear_pull", linear.PullIssues)
		case "next":
			return runLinearOperation(root, "linear_next", linear.NextIssue)
		case "create-from-backlog":
			result := linear.CreateFromBacklog(context.Background(), root, http.DefaultClient)
			return writeJSON(linear.SummarizeOperation(result))
		case "replay-offline":
			return writeJSON(linear.ReplayOffline(context.Background(), root, http.DefaultClient))
		default:
			return usage()
		}
	}
	if len(args) >= 3 {
		switch args[0] {
		case "comment":
			result := linear.AddComment(context.Background(), root, http.DefaultClient, args[1], strings.Join(args[2:], " "))
			return writeJSON(linear.SummarizeOperation(result))
		case "transition":
			result := linear.TransitionIssue(context.Background(), root, http.DefaultClient, args[1], strings.Join(args[2:], " "))
			return writeJSON(linear.SummarizeOperation(result))
		default:
			return usage()
		}
	}
	return usage()
}

func runLinearOperation(root string, event string, operation linearOperation) error {
	result := operation(context.Background(), root, http.DefaultClient)
	if !result.Synced {
		if err := linear.AppendOfflineEvent(root, event, result.Message); err != nil {
			return err
		}
	}
	return writeJSON(linear.SummarizeOperation(result))
}

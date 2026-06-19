package linear

import (
	"context"
	"net/http"
	"time"
)

const Endpoint = "https://api.linear.app/graphql"

func CurrentStatus(root string) Status {
	return StatusForRequest(context.Background(), root, http.DefaultClient)
}

func StatusForRequest(parent context.Context, root string, client *http.Client) Status {
	queuePath := filepathJoinSlash(root, "data", "private", "linear-offline-queue.jsonl")
	status := Status{
		Mode:      "offline",
		Synced:    false,
		QueuePath: queuePath,
		Endpoint:  Endpoint,
	}
	token, err := loadToken(root)
	if err != nil {
		status.Message = "No Linear token found. Continuing in offline mode."
		return status
	}
	status.Mode = "configured"
	status.TokenConfigured = true
	status.TokenSource = token.Source

	ctx, cancel := context.WithTimeout(parent, 15*time.Second)
	defer cancel()
	viewer, teams, httpStatus, remaining, err := queryViewer(ctx, client, token.Value)
	status.HTTPStatus = httpStatus
	status.RateLimitRemaining = remaining
	if err != nil {
		status.Mode = "offline"
		status.Message = "Linear GraphQL status check failed; continuing with offline fallback: " + err.Error()
		return status
	}
	status.Mode = "online"
	status.Synced = true
	status.Viewer = viewer
	status.Teams = teams
	status.Message = "Linear GraphQL status check succeeded."
	return status
}

func SummarizeStatus(status Status) StatusSummary {
	return StatusSummary{
		Mode:               status.Mode,
		TokenConfigured:    status.TokenConfigured,
		Synced:             status.Synced,
		QueuePath:          privateRelativePath(status.QueuePath),
		Endpoint:           status.Endpoint,
		HTTPStatus:         status.HTTPStatus,
		RateLimitRemaining: status.RateLimitRemaining,
		ViewerConfigured:   status.Viewer != nil,
		TeamCount:          len(status.Teams),
		Message:            status.Message,
	}
}

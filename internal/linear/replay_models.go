package linear

import (
	"encoding/json"
	"net/http"
)

const (
	OfflineQueueRelativePath    = "data/private/linear-offline-queue.jsonl"
	OfflineReplayRelativePath   = "data/private/linear-offline-replay.jsonl"
	defaultReplayRateLimitFloor = 100
	offlineReplayCommentKind    = "linear_comment"
	offlineReplayTransitionKind = "linear_transition"
)

type ReplayResult struct {
	Mode                 string `json:"mode"`
	Synced               bool   `json:"synced"`
	QueuePath            string `json:"queue_path"`
	ReplayPath           string `json:"replay_path"`
	HTTPStatus           int    `json:"http_status,omitempty"`
	RateLimitRemaining   int    `json:"rate_limit_remaining,omitempty"`
	Message              string `json:"message"`
	QueuedCount          int    `json:"queued_count"`
	EligibleCount        int    `json:"eligible_count"`
	ReplayedCount        int    `json:"replayed_count"`
	AlreadyReplayedCount int    `json:"already_replayed_count,omitempty"`
	SkippedCount         int    `json:"skipped_count,omitempty"`
	FailedCount          int    `json:"failed_count,omitempty"`
	RateLimited          bool   `json:"rate_limited,omitempty"`
}

type replayRecord struct {
	At       string `json:"at"`
	EntryID  string `json:"entry_id"`
	Kind     string `json:"kind"`
	IssueKey string `json:"issue_key,omitempty"`
	Synced   bool   `json:"synced"`
}

type queuedOfflineAction struct {
	OfflineEvent
	Payload json.RawMessage `json:"payload,omitempty"`
	EntryID string          `json:"-"`
}

type commentPayload struct {
	IssueID string `json:"issue_id"`
	Body    string `json:"body"`
}

type transitionPayload struct {
	IssueID string `json:"issue_id"`
	State   string `json:"state"`
}

type replayRun struct {
	result   ReplayResult
	root     string
	client   *http.Client
	token    string
	scope    issueScope
	replayed map[string]bool
}

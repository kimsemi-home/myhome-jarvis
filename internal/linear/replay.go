package linear

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
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

func ReplayOffline(ctx context.Context, root string, client *http.Client) ReplayResult {
	result := ReplayResult{
		Mode:       "offline",
		Synced:     false,
		QueuePath:  privateRelativePath(filepathJoinSlash(root, OfflineQueueRelativePath)),
		ReplayPath: privateRelativePath(filepathJoinSlash(root, OfflineReplayRelativePath)),
	}
	entries, err := readQueuedOfflineActions(root)
	if err != nil {
		result.Message = "Offline queue could not be read; replay skipped."
		return result
	}
	result.QueuedCount = len(entries)
	if len(entries) == 0 {
		result.Synced = true
		result.Message = "No offline Linear actions are queued."
		return result
	}
	replayed, err := readReplayedEntryIDs(root)
	if err != nil {
		result.Message = "Offline replay evidence could not be read; replay skipped."
		return result
	}
	for _, entry := range entries {
		if !entry.Synced && replaySafeKind(entry.Kind) {
			result.EligibleCount++
		}
	}
	if result.EligibleCount == 0 {
		result.Synced = true
		result.Message = "No write-safe offline Linear actions are queued for replay."
		return result
	}
	token, err := loadToken(root)
	if err != nil {
		result.Message = "No Linear token found. Offline replay skipped and queued entries remain synced=false."
		return result
	}
	result.Mode = "online"
	for _, entry := range entries {
		if entry.Synced || !replaySafeKind(entry.Kind) {
			if !entry.Synced {
				result.SkippedCount++
			}
			continue
		}
		if replayed[entry.EntryID] {
			result.AlreadyReplayedCount++
			continue
		}
		operation := replayQueuedAction(ctx, root, client, token.Value, entry)
		result.HTTPStatus = operation.HTTPStatus
		result.RateLimitRemaining = operation.RateLimitRemaining
		if operation.RateLimited {
			result.RateLimited = true
			result.Message = operation.Message
			break
		}
		if !operation.Synced {
			result.FailedCount++
			result.Message = operation.Message
			break
		}
		result.ReplayedCount++
		if err := appendReplayRecord(root, entry); err != nil {
			result.FailedCount++
			result.Message = "Offline replay succeeded but private replay evidence was not recorded."
			break
		}
		if rateLimitLow(operation.RateLimitRemaining) {
			result.RateLimited = true
			result.Message = "Linear rate limit remaining is low; offline replay paused."
			break
		}
	}
	if result.Message == "" {
		result.Message = fmt.Sprintf("Replayed %d write-safe offline Linear actions.", result.ReplayedCount)
	}
	result.Synced = result.FailedCount == 0 && !result.RateLimited && result.ReplayedCount+result.AlreadyReplayedCount == result.EligibleCount
	return result
}

func replayQueuedAction(ctx context.Context, root string, client *http.Client, token string, entry queuedOfflineAction) ReplayResult {
	switch entry.Kind {
	case offlineReplayCommentKind:
		var payload commentPayload
		if err := json.Unmarshal(entry.Payload, &payload); err != nil {
			return failedReplay("Offline Linear comment payload is invalid.")
		}
		return replayComment(ctx, root, client, token, payload)
	case offlineReplayTransitionKind:
		var payload transitionPayload
		if err := json.Unmarshal(entry.Payload, &payload); err != nil {
			return failedReplay("Offline Linear transition payload is invalid.")
		}
		return replayTransition(ctx, root, client, token, payload)
	default:
		return failedReplay("Offline Linear action is not replay-safe.")
	}
}

func replayComment(ctx context.Context, root string, client *http.Client, token string, payload commentPayload) ReplayResult {
	issueID := strings.TrimSpace(payload.IssueID)
	body := strings.TrimSpace(payload.Body)
	if issueID == "" || body == "" {
		return failedReplay("Offline Linear comment payload is incomplete.")
	}
	var response struct {
		CommentCreate struct {
			Success bool     `json:"success"`
			Comment *Comment `json:"comment"`
		} `json:"commentCreate"`
	}
	query := `mutation AddComment($issueId: String!, $body: String!) { commentCreate(input: { issueId: $issueId, body: $body }) { success comment { id createdAt } } }`
	httpStatus, remaining, err := doGraphQL(ctx, client, token, query, map[string]string{"issueId": issueID, "body": body}, &response)
	if err != nil || !response.CommentCreate.Success {
		return ReplayResult{
			Mode:               "online",
			Synced:             false,
			HTTPStatus:         httpStatus,
			RateLimitRemaining: remaining,
			Message:            "Offline Linear comment replay failed; entry remains synced=false.",
		}
	}
	_ = AppendWriteEvidence(root, offlineReplayCommentKind, issueID)
	return ReplayResult{
		Mode:               "online",
		Synced:             true,
		HTTPStatus:         httpStatus,
		RateLimitRemaining: remaining,
		Message:            "Offline Linear comment replayed.",
	}
}

func replayTransition(ctx context.Context, root string, client *http.Client, token string, payload transitionPayload) ReplayResult {
	issueID := strings.TrimSpace(payload.IssueID)
	stateName := strings.TrimSpace(payload.State)
	if issueID == "" || stateName == "" {
		return failedReplay("Offline Linear transition payload is incomplete.")
	}
	stateID, httpStatus, remaining, err := findWorkflowStateID(ctx, client, token, stateName)
	if err != nil {
		return ReplayResult{
			Mode:               "online",
			Synced:             false,
			HTTPStatus:         httpStatus,
			RateLimitRemaining: remaining,
			Message:            "Offline Linear transition lookup failed; entry remains synced=false.",
		}
	}
	if rateLimitLow(remaining) {
		return ReplayResult{
			Mode:               "online",
			Synced:             false,
			HTTPStatus:         httpStatus,
			RateLimitRemaining: remaining,
			RateLimited:        true,
			Message:            "Linear rate limit remaining is low; offline replay paused before transition mutation.",
		}
	}
	var response struct {
		IssueUpdate struct {
			Success bool   `json:"success"`
			Issue   *Issue `json:"issue"`
		} `json:"issueUpdate"`
	}
	query := `mutation TransitionIssue($issueId: String!, $stateId: String!) { issueUpdate(id: $issueId, input: { stateId: $stateId }) { success issue { id identifier title state { id name type } } } }`
	httpStatus, remaining, err = doGraphQL(ctx, client, token, query, map[string]string{"issueId": issueID, "stateId": stateID}, &response)
	if err != nil || !response.IssueUpdate.Success {
		return ReplayResult{
			Mode:               "online",
			Synced:             false,
			HTTPStatus:         httpStatus,
			RateLimitRemaining: remaining,
			Message:            "Offline Linear transition replay failed; entry remains synced=false.",
		}
	}
	evidenceIssueKey := issueID
	if response.IssueUpdate.Issue != nil {
		evidenceIssueKey = response.IssueUpdate.Issue.Identifier
	}
	_ = AppendWriteEvidence(root, offlineReplayTransitionKind, evidenceIssueKey)
	return ReplayResult{
		Mode:               "online",
		Synced:             true,
		HTTPStatus:         httpStatus,
		RateLimitRemaining: remaining,
		Message:            "Offline Linear transition replayed.",
	}
}

func readQueuedOfflineActions(root string) ([]queuedOfflineAction, error) {
	path := filepath.Join(root, filepath.FromSlash(OfflineQueueRelativePath))
	file, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	entries := []queuedOfflineAction{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var entry queuedOfflineAction
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			return nil, err
		}
		entry.Kind = strings.TrimSpace(entry.Kind)
		entry.EntryID = offlineEntryID(line)
		entries = append(entries, entry)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return entries, nil
}

func readReplayedEntryIDs(root string) (map[string]bool, error) {
	path := filepath.Join(root, filepath.FromSlash(OfflineReplayRelativePath))
	file, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return map[string]bool{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	replayed := map[string]bool{}
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		var record replayRecord
		if err := json.Unmarshal(scanner.Bytes(), &record); err != nil {
			return nil, err
		}
		if record.Synced && strings.TrimSpace(record.EntryID) != "" {
			replayed[record.EntryID] = true
		}
	}
	return replayed, scanner.Err()
}

func appendReplayRecord(root string, entry queuedOfflineAction) error {
	path := filepath.Join(root, filepath.FromSlash(OfflineReplayRelativePath))
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(replayRecord{
		At:       time.Now().UTC().Format(time.RFC3339),
		EntryID:  entry.EntryID,
		Kind:     entry.Kind,
		IssueKey: replayIssueKey(entry),
		Synced:   true,
	})
}

func replayIssueKey(entry queuedOfflineAction) string {
	switch entry.Kind {
	case offlineReplayCommentKind:
		var payload commentPayload
		if json.Unmarshal(entry.Payload, &payload) == nil {
			return publicIssueKey(payload.IssueID)
		}
	case offlineReplayTransitionKind:
		var payload transitionPayload
		if json.Unmarshal(entry.Payload, &payload) == nil {
			return publicIssueKey(payload.IssueID)
		}
	}
	return ""
}

func replaySafeKind(kind string) bool {
	switch strings.TrimSpace(kind) {
	case offlineReplayCommentKind, offlineReplayTransitionKind:
		return true
	default:
		return false
	}
}

func rateLimitLow(remaining int) bool {
	return remaining > 0 && remaining <= defaultReplayRateLimitFloor
}

func offlineEntryID(line string) string {
	sum := sha256.Sum256([]byte(line))
	return hex.EncodeToString(sum[:])
}

func failedReplay(message string) ReplayResult {
	return ReplayResult{
		Mode:    "online",
		Synced:  false,
		Message: message,
	}
}

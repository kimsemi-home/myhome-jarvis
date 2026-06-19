package linear

import (
	"context"
	"encoding/json"
	"net/http"
)

func replayQueuedAction(
	ctx context.Context,
	root string,
	client *http.Client,
	token string,
	entry queuedOfflineAction,
) ReplayResult {
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

func failedReplay(message string) ReplayResult {
	return ReplayResult{
		Mode:    "online",
		Synced:  false,
		Message: message,
	}
}

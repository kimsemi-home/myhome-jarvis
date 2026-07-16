package localfinanceevidence

import (
	"encoding/json"
	"errors"
	"slices"
)

func validateShortsExecutions(value ShortsReport) error {
	if err := validateShortsReceipt(value.First, value.PlanHash, value.ContentHash); err != nil {
		return err
	}
	if err := validateShortsReceipt(value.Replay, value.PlanHash, value.ContentHash); err != nil {
		return err
	}
	if value.First.SessionRequests != 1 || value.First.ProbeRequests < 2 || value.First.ChunkRequests < 1 ||
		value.First.RecoveryAttempts != 1 || value.Replay.SessionRequests != 0 || value.Replay.ProbeRequests != 0 ||
		value.Replay.ChunkRequests != 0 || value.Replay.RecoveryAttempts != 0 ||
		value.First.AcceptedBytes != value.Replay.AcceptedBytes || value.First.VideoIdentityHash != value.Replay.VideoIdentityHash ||
		value.Emulator != (ShortsEmulator{SessionCreates: 1, VideoCreates: 1, InterruptedRequests: 1, TokenRequests: 2, ChannelRequests: 1, RedirectRequests: 1, OversizedRequests: 1}) {
		return errors.New("Shorts upload recovery or replay proof is invalid")
	}
	required := []string{"channel-identity-bound", "credentials-not-read", "external-network-disabled", "external-writes-disabled", "notify-subscribers-disabled", "oauth-official-origin-pinned", "oauth-redirect-rejected", "oauth-response-bounded", "oauth-scope-validated", "oauth-token-contract-validated", "private-upload-completed", "resumable-recovery-bounded", "session-location-pinned", "session-state-replay-idempotent"}
	if !slices.Equal(value.Checks, required) {
		return errors.New("Shorts execution proof checks are invalid")
	}
	return nil
}

func validateShortsReceipt(value ShortsReceipt, planHash, contentHash string) error {
	if value.SchemaVersion != "shorts.youtube-loopback-upload-receipt/v1" || value.ExecutionMode != "loopback_emulation" ||
		!value.LoopbackOnly || value.ExternalWrites || value.PlanHash != planHash || value.ContentHash != contentHash ||
		value.AcceptedBytes < 1 || !value.Complete || value.PrivacyStatus != "private" ||
		!hashPattern.MatchString(value.VideoIdentityHash) || !hashPattern.MatchString(value.ReceiptHash) {
		return errors.New("Shorts upload receipt boundary is invalid")
	}
	copy := value
	copy.ReceiptHash = ""
	body, err := json.Marshal(copy)
	if err != nil || value.ReceiptHash != digest(string(body)) {
		return errors.New("Shorts upload receipt hash is invalid")
	}
	return nil
}

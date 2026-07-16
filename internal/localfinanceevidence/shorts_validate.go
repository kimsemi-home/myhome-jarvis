package localfinanceevidence

import (
	"encoding/json"
	"errors"
)

func validateShortsReport(value ShortsReport, ref ProofRef) error {
	if value.SchemaVersion != shortsProofSchema || value.ExecutionMode != "loopback_emulation" || !value.LoopbackOnly ||
		value.CredentialsRead || value.ExternalNetwork || value.ExternalWrites || !value.SessionStateReplay ||
		!value.SessionLocationPinned || !value.NotificationsDisabled || !hashPattern.MatchString(value.PlanHash) ||
		!hashPattern.MatchString(value.ContentHash) || value.ReportHash != ref.ReportHash {
		return errors.New("Shorts execution proof boundary is invalid")
	}
	if value.OAuth.SchemaVersion != "shorts.youtube-oauth-token-rehearsal/v1" || !value.OAuth.AuthorizationCodeExchange ||
		!value.OAuth.RefreshTokenExchange || !value.OAuth.ScopeContractValidated || !value.OAuth.OfficialOriginPinned ||
		!value.OAuth.RedirectRejected || !value.OAuth.OversizedResponseRejected {
		return errors.New("Shorts OAuth execution proof boundary is invalid")
	}
	if value.Channel.SchemaVersion != "shorts.youtube-channel-binding-rehearsal/v1" || value.Channel.Method != "GET" ||
		value.Channel.Path != "/youtube/v3/channels" || value.Channel.Query != "mine=true&part=id" ||
		!value.Channel.ExactlyOneChannel || !value.Channel.BindingHashMatched || value.Channel.RawIdentityPersisted ||
		!value.Channel.OfficialOriginPinned {
		return errors.New("Shorts channel identity proof boundary is invalid")
	}
	if err := validateShortsExecutions(value); err != nil {
		return err
	}
	copy := value
	copy.ReportHash = ""
	body, err := json.Marshal(copy)
	if err != nil || value.ReportHash != digest(string(body)) {
		return errors.New("Shorts execution report hash is invalid")
	}
	return nil
}

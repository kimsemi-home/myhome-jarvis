package security

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

type generatedSecurityPolicy struct {
	CurrentContentScan             bool `json:"current_content_scan"`
	CurrentContentScanSkipsPrivate bool `json:"current_content_scan_skips_private_paths"`
	PrivateIdentityScan            bool `json:"private_identity_scan"`
	SecretLiteralScan              bool `json:"secret_literal_scan"`
	ReportMatchedSecretContents    bool `json:"report_matched_secret_contents"`
	StatusCache                    struct {
		Enabled                  bool   `json:"enabled"`
		Path                     string `json:"path"`
		Mode                     string `json:"mode"`
		ValidationCommand        string `json:"validation_command"`
		MissRunsFullHistory      bool   `json:"miss_runs_full_history"`
		CurrentScanAlwaysFresh   bool   `json:"current_scan_always_fresh"`
		RawFindingsPublicAllowed bool   `json:"raw_findings_public_allowed"`
		PublicSafe               bool   `json:"public_safe"`
	} `json:"status_cache"`
}

func TestGeneratedPolicyRecordsCurrentContentScan(t *testing.T) {
	data, err := os.ReadFile(filepath.Join(repoRoot(t), "generated", "security.generated.json"))
	if err != nil {
		t.Fatal(err)
	}
	var policy generatedSecurityPolicy
	if err := json.Unmarshal(data, &policy); err != nil {
		t.Fatal(err)
	}
	if !policy.CurrentContentScan {
		t.Fatal("generated security policy must enable current content scanning")
	}
	if !policy.CurrentContentScanSkipsPrivate {
		t.Fatal("generated security policy must keep current content scanning out of private paths")
	}
	if !policy.PrivateIdentityScan || !policy.SecretLiteralScan {
		t.Fatalf("generated security policy missing content scan coverage: %#v", policy)
	}
	if policy.ReportMatchedSecretContents {
		t.Fatal("generated security policy must not report matched secret contents")
	}
	if !policy.StatusCache.Enabled || policy.StatusCache.Path != statusCachePath {
		t.Fatalf("generated security status cache mismatch: %#v", policy.StatusCache)
	}
	if policy.StatusCache.Mode != "history_aggregate_only" ||
		policy.StatusCache.ValidationCommand != statusCacheValidation {
		t.Fatalf("generated security cache mode mismatch: %#v", policy.StatusCache)
	}
	if !policy.StatusCache.MissRunsFullHistory || !policy.StatusCache.CurrentScanAlwaysFresh {
		t.Fatalf("generated security cache weakens safety: %#v", policy.StatusCache)
	}
	if policy.StatusCache.RawFindingsPublicAllowed || !policy.StatusCache.PublicSafe {
		t.Fatalf("generated security cache leaks findings: %#v", policy.StatusCache)
	}
}

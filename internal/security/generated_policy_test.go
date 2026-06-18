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
}

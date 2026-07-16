package localfinanceevidence

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestRehashedUnsafeShortsActivationProofFails(t *testing.T) {
	manifestPath := filepath.Join("..", "..", "fixtures", "local_finance", "manifest.json")
	manifest, err := Read(manifestPath)
	if err != nil {
		t.Fatal(err)
	}
	var ref ProofRef
	for _, candidate := range manifest.ExecutionProofs {
		if candidate.Component == "shorts-activation" {
			ref = candidate
		}
	}
	body, err := os.ReadFile(filepath.Join(filepath.Dir(manifestPath), ref.Path))
	if err != nil {
		t.Fatal(err)
	}
	var report ShortsActivationReport
	if err := json.Unmarshal(body, &report); err != nil {
		t.Fatal(err)
	}
	report.Keychain.ActualKeychainExecuted = true
	report.ReportHash = ""
	unsigned, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	report.ReportHash = digest(string(unsigned))
	ref.ReportHash = report.ReportHash
	tampered, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	if err := validateProofBody(tampered, manifest.Month, ref); err == nil {
		t.Fatal("accepted rehashed activation proof after real Keychain execution was enabled")
	}
}

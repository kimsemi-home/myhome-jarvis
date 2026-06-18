package controlplane

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func scanLedger(root string, policy Policy, status *Status) error {
	path := filepath.Join(root, filepath.FromSlash(policy.PrivateManifestLedger))
	file, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer file.Close()

	status.Exists = true
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		recordStatusLine(policy, status, scanner.Text())
	}
	return scanner.Err()
}

func recordStatusLine(policy Policy, status *Status, line string) {
	line = strings.TrimSpace(line)
	if line == "" {
		return
	}
	manifest, ok := parseManifestLine(policy, status, line)
	if !ok {
		return
	}
	status.Count++
	status.ByDecisionKind[manifest.DecisionKind]++
	status.ByAuthorityProfile[manifest.AuthorityProfile]++
	status.ByLeaseStatus[manifest.LeaseStatus]++
	if policy.VerifierSeparationRequired && manifest.ReviewerRole == manifest.VerifierRole {
		status.VerifierViolationCount++
	}
	status.LastObservedAt = laterRFC3339(status.LastObservedAt, manifest.At)
}

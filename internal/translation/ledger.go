package translation

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func inspectLedger(root string, policy Policy, status *Status) error {
	path := filepath.Join(root, filepath.FromSlash(policy.PrivateLossLedger))
	file, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer file.Close()

	status.LedgerExists = true
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		recordLedgerLine(root, policy, status, scanner.Text())
	}
	return scanner.Err()
}

func recordLedgerLine(root string, policy Policy, status *Status, line string) {
	line = strings.TrimSpace(line)
	if line == "" {
		return
	}
	var record lossRecord
	if err := json.Unmarshal([]byte(line), &record); err != nil {
		status.InvalidLossCount++
		return
	}
	status.LossCount++
	updateLastObserved(status, record.At)
	err := recordLoss(policy, status, record.SourceContext, record.TargetContext, record.Level, record.Category, record.Status)
	if err != nil {
		status.InvalidLossCount++
	}
	if policy.ManifestRequired && validateReferencedManifest(root, policy, record.ManifestPath) != nil {
		status.MissingManifestCount++
	}
}

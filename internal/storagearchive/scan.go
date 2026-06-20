package storagearchive

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kimsemi-home/myhome-jarvis/internal/domain"
)

func scanSource(content []byte, noise domain.EvidenceNoiseBudget) (sourceScan, error) {
	seen := map[string]bool{}
	scan := sourceScan{Content: append([]byte{}, content...)}
	sum := sha256.Sum256(content)
	scan.InputSHA256 = hex.EncodeToString(sum[:])
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		scan.RecordCount++
		if recordNoise(line, noise.DedupeKeyFields, seen) {
			scan.NoiseCount++
		}
	}
	if scan.RecordCount > 0 {
		scan.NoiseRatioPercent = (scan.NoiseCount*100 + scan.RecordCount - 1) / scan.RecordCount
	}
	scan.BudgetOK = scan.NoiseRatioPercent <= noise.MaxNoiseRatioPercent &&
		scan.NoiseCount <= noise.MaxLowSignalRecordsPerWindow
	if err := scanner.Err(); err != nil {
		return sourceScan{}, fmt.Errorf("scan source jsonl: %w", err)
	}
	return scan, nil
}

func recordNoise(line string, fields []string, seen map[string]bool) bool {
	var record map[string]any
	if err := json.Unmarshal([]byte(line), &record); err != nil {
		return true
	}
	key, ok := dedupeKey(record, fields)
	if !ok {
		return false
	}
	if seen[key] {
		return true
	}
	seen[key] = true
	return false
}

func dedupeKey(record map[string]any, fields []string) (string, bool) {
	parts := make([]string, 0, len(fields))
	for _, field := range fields {
		value, ok := record[field]
		if !ok {
			return "", false
		}
		parts = append(parts, field+"="+strings.TrimSpace(toText(value)))
	}
	return strings.Join(parts, "|"), len(parts) > 0
}

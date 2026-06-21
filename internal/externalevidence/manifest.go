package externalevidence

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type manifestSummary struct {
	Present   bool
	Count     int
	LatestAt  string
	KnownRefs map[string]bool
}

func readManifestSummary(root string, rel string) (manifestSummary, error) {
	path := filepath.Join(root, filepath.FromSlash(rel))
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		return manifestSummary{KnownRefs: map[string]bool{}}, nil
	}
	if err != nil {
		return manifestSummary{}, err
	}
	defer file.Close()
	summary := manifestSummary{Present: true, KnownRefs: map[string]bool{}}
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		applyManifestLine(&summary, scanner.Text())
	}
	return summary, scanner.Err()
}

func applyManifestLine(summary *manifestSummary, line string) {
	line = strings.TrimSpace(line)
	if line == "" {
		return
	}
	var record struct {
		At          string `json:"at"`
		EvidenceRef string `json:"evidence_ref"`
	}
	if err := json.Unmarshal([]byte(line), &record); err != nil {
		return
	}
	summary.Count++
	if record.At > summary.LatestAt {
		summary.LatestAt = record.At
	}
	if record.EvidenceRef != "" {
		summary.KnownRefs[record.EvidenceRef] = true
	}
}

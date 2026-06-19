package incidents

import (
	"bufio"
	"encoding/json"
	"io"
	"strings"
	"time"
)

func scanIncidentLedger(
	reader io.Reader,
	policy Policy,
	checkedAt time.Time,
	status *Status,
) error {
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var incident Incident
		if err := json.Unmarshal([]byte(line), &incident); err != nil {
			status.InvalidIncidentCount++
			continue
		}
		normalized, err := normalizeIncident(policy, incident)
		if err != nil {
			recordIncidentError(status, err)
			continue
		}
		addIncidentStatus(policy, checkedAt, status, normalized)
	}
	return scanner.Err()
}

package financeconsent

import (
	"bufio"
	"encoding/json"
	"io"
)

func scanLedger(reader io.Reader, policy Policy, status *Status) (map[string]bool, error) {
	activeKinds := map[string]bool{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		var record Record
		if err := json.Unmarshal(scanner.Bytes(), &record); err != nil {
			status.InvalidRecordCount++
			continue
		}
		status.RecordCount++
		result := validateRecord(record, policy)
		if !result.Valid {
			status.InvalidRecordCount++
		}
		if result.MissingEvidence {
			status.MissingEvidenceCount++
		}
		if result.ReviewRequired {
			status.ReviewRequiredCount++
		}
		if result.RevokedOrExpired {
			status.RevokedOrExpiredCount++
		}
		if result.Active {
			status.ActiveConsentCount++
			activeKinds[record.ConsentKind] = true
		}
	}
	return activeKinds, scanner.Err()
}

package localfinanceevidence

import (
	"encoding/json"
	"strconv"
	"strings"
)

func creditApplyPlanHash(value CreditBatchApplyPlan) string {
	value.PlanHash = ""
	body, _ := json.Marshal(value)
	return digest(string(body))
}

func creditApprovalHash(value CreditBatchApproval) string {
	value.ApprovalHash = ""
	body, _ := json.Marshal(value)
	return digest(string(body))
}

func creditApplyReportHash(value CreditBatchApplyReport) string {
	value.ReportHash = ""
	body, _ := json.Marshal(value)
	return digest(string(body))
}

func creditApplyRehearsalHash(value CreditBatchApplyRehearsal) string {
	value.ReportHash = ""
	body, _ := json.Marshal(value)
	return digest(string(body))
}

func creditApprovalChallenge(value CreditBatchApplyPlan) string {
	parts := []string{"approve-local-batch-apply-v1", value.ManifestSHA256, value.PreviewSetHash,
		value.BatchHash, strconv.Itoa(value.StatementCount)}
	return digest(strings.Join(parts, "\x00"))
}

func creditApplySetHash(values []CreditBatchApplyStatement) string {
	parts := make([]string, 0, len(values))
	for _, value := range values {
		fields := []string{value.SourceNameSHA256, value.PreviewHash, strconv.Itoa(value.RowsRead),
			strconv.Itoa(value.RowsInserted), strconv.Itoa(value.Duplicates), strconv.Itoa(value.Suggestions)}
		parts = append(parts, strings.Join(fields, "\x00"))
	}
	return digest(strings.Join(parts, "\n"))
}

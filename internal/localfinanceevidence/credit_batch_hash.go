package localfinanceevidence

import "strings"

func creditBatchPreviewSetHash(statements []CreditBatchStatement) string {
	values := make([]string, 0, len(statements))
	for _, statement := range statements {
		values = append(values, statement.SourceNameSHA256+"\x00"+statement.Preview.PreviewHash)
	}
	return digest(strings.Join(values, "\n"))
}

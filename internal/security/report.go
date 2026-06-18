package security

import (
	"path/filepath"
	"sort"
)

func (report *Report) add(path string, code string, message string) {
	report.Findings = append(report.Findings, Finding{Path: path, Code: code, Message: message})
}

func (report *Report) addLine(path string, line int, code string, message string) {
	report.Findings = append(report.Findings, Finding{Path: path, Line: line, Code: code, Message: message})
}

func (report *HistoryReport) addHistory(commit string, path string, line int, code string, message string) {
	report.Findings = append(report.Findings, HistoryFinding{
		Commit:  commit,
		Path:    filepath.ToSlash(path),
		Line:    line,
		Code:    code,
		Message: message,
	})
}

func sortFindings(findings []Finding) {
	sort.Slice(findings, func(i, j int) bool {
		if findings[i].Path == findings[j].Path {
			if findings[i].Line != findings[j].Line {
				return findings[i].Line < findings[j].Line
			}
			return findings[i].Code < findings[j].Code
		}
		return findings[i].Path < findings[j].Path
	})
}

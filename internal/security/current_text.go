package security

import "regexp"

func checkCurrentTextLine(rel string, line int, text string, privateIdentity *regexp.Regexp, report *Report) {
	if privateIdentity.MatchString(text) {
		report.addLine(rel, line, "current_private_identity", "current files must not contain private user, path, or old-repository identity markers")
	}
	if secretHistoryPattern.MatchString(text) {
		report.addLine(rel, line, "current_secret_literal", "current files must not contain secret-looking literal values")
	}
}

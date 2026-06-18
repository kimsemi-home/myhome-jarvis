package security

import "strings"

func checkHistoryLanguagePath(commit string, rel string, ext string, report *HistoryReport) {
	switch ext {
	case ".py", ".pyi", ".ipynb":
		report.addHistory(commit, rel, 0, "history_forbidden_python_file", "Python files must never appear in git history")
	case ".ts", ".tsx", ".js", ".jsx", ".mjs", ".cjs":
		report.addHistory(commit, rel, 0, "history_forbidden_node_typescript_file", "Node.js and TypeScript files must never appear in git history")
	}
}

func checkHistoryDependencyPath(commit string, rel string, base string, report *HistoryReport) {
	switch base {
	case "pyproject.toml", "pytest.ini", "requirements.txt", "requirements-dev.txt",
		"poetry.lock", "uv.lock", "pipfile", "pipfile.lock", "setup.py", "tox.ini":
		report.addHistory(commit, rel, 0, "history_forbidden_python_dependency", "Python dependency or tooling files must never appear in git history")
	case "package.json", "package-lock.json", "pnpm-lock.yaml", "yarn.lock", "tsconfig.json":
		report.addHistory(commit, rel, 0, "history_forbidden_node_dependency", "Node.js and TypeScript dependency files must never appear in git history")
	}
}

func checkHistorySensitivePath(commit string, rel string, base string, report *HistoryReport) {
	for _, marker := range []string{"token", "secret", "credential", "cookie"} {
		if strings.Contains(base, marker) {
			report.addHistory(commit, rel, 0, "history_forbidden_sensitive_path", "sensitive-looking files must never appear in git history")
			return
		}
	}
}

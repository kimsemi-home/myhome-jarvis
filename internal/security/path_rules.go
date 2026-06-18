package security

import "strings"

func checkLanguagePath(rel string, ext string, report *Report) {
	switch ext {
	case ".py", ".pyi", ".ipynb":
		report.add(rel, "forbidden_python_file", "Python files are not allowed")
	case ".ts", ".tsx", ".js", ".jsx", ".mjs", ".cjs":
		report.add(rel, "forbidden_node_typescript_file", "Node.js and TypeScript files are not allowed")
	}
}

func checkDependencyPath(rel string, base string, report *Report) {
	switch base {
	case "pyproject.toml", "pytest.ini", "requirements.txt", "requirements-dev.txt",
		"poetry.lock", "uv.lock", "pipfile", "pipfile.lock", "setup.py", "tox.ini":
		report.add(rel, "forbidden_python_dependency", "Python dependency or tooling files are not allowed")
	case "package.json", "package-lock.json", "pnpm-lock.yaml", "yarn.lock", "tsconfig.json":
		report.add(rel, "forbidden_node_dependency", "Node.js and TypeScript dependency files are not allowed")
	}
}

func checkSensitivePath(rel string, base string, report *Report) {
	for _, marker := range []string{"token", "secret", "credential", "cookie"} {
		if strings.Contains(base, marker) {
			report.add(rel, "forbidden_sensitive_path", "sensitive-looking files must stay under data/private")
			return
		}
	}
}

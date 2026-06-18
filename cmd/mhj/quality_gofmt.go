package main

import (
	"os"
	"path/filepath"
	"strings"
)

func (report *qualityReport) addGofmt(root string, gofmtTool string) {
	files, err := collectGoFiles(root)
	if err != nil {
		report.OK = false
		report.Steps = append(report.Steps, qualityStep{Name: "gofmt", Status: "fail", Output: err.Error()})
		return
	}
	if len(files) == 0 {
		report.Steps = append(report.Steps, qualityStep{Name: "gofmt", Status: "skip", Output: "no Go files"})
		return
	}
	command := append([]string{gofmtTool, "-l"}, files...)
	report.addCommand(root, "gofmt", command)
	last := &report.Steps[len(report.Steps)-1]
	if last.Status == "pass" && strings.TrimSpace(last.Output) != "" {
		last.Status = "fail"
		last.Output = "unformatted files:\n" + last.Output
		report.OK = false
	}
}

func collectGoFiles(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			switch entry.Name() {
			case ".git", "target", "build", "dist", "bin":
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		files = append(files, filepath.ToSlash(rel))
		return nil
	})
	return files, err
}

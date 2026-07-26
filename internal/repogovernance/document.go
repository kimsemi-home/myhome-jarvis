package repogovernance

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GenerateDocuments(root string) error {
	doc, err := loadDocument(root)
	if err != nil {
		return err
	}
	output := filepath.Join(root, filepath.FromSlash(doc.Output))
	if err := os.MkdirAll(filepath.Dir(output), 0o755); err != nil {
		return err
	}
	return os.WriteFile(output, renderDocument(doc), 0o644)
}

func CheckDocuments(root string) error {
	doc, err := loadDocument(root)
	if err != nil {
		return err
	}
	body, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(doc.Output)))
	if err != nil {
		return err
	}
	if !bytes.Equal(body, renderDocument(doc)) {
		return fmt.Errorf("generated document is stale: %s", doc.Output)
	}
	return nil
}

func loadDocument(root string) (Document, error) {
	var doc Document
	if err := decode(filepath.Join(root, "docs-src", "shorts-factory.json"), &doc); err != nil {
		return doc, err
	}
	if doc.SchemaVersion != "repo.document/v1" || doc.ID == "" || doc.Title == "" || doc.Summary == "" || len(doc.Sections) == 0 {
		return doc, errors.New("document source is incomplete")
	}
	if filepath.IsAbs(doc.Output) || filepath.Clean(doc.Output) != doc.Output || strings.HasPrefix(doc.Output, "..") || !strings.HasSuffix(doc.Output, ".md") {
		return doc, errors.New("document output is unsafe")
	}
	return doc, nil
}

package financeconnector

import (
	"os"
	"path/filepath"
)

func LoadFixture(root string) ([]SourceTransaction, error) {
	file, err := os.Open(filepath.Join(root, FixtureRelativePath))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return scanFixture(file)
}

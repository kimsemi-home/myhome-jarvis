package financeconnector

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const FixtureRelativePath = "fixtures/finance_toss_mydata.jsonl"

func LoadFixture(root string) ([]SourceTransaction, error) {
	file, err := os.Open(filepath.Join(root, FixtureRelativePath))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var transactions []SourceTransaction
	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		line++
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		var transaction SourceTransaction
		if err := json.Unmarshal([]byte(text), &transaction); err != nil {
			return nil, fmt.Errorf("%s line %d: %w", FixtureRelativePath, line, err)
		}
		transactions = append(transactions, transaction)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(transactions) == 0 {
		return nil, fmt.Errorf("%s contains no transactions", FixtureRelativePath)
	}
	return transactions, nil
}

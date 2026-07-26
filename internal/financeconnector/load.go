package financeconnector

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const FixtureRelativePath = "fixtures/finance_toss_mydata.jsonl"

type LoadedData struct {
	Transactions []SourceTransaction
	Parity       ParityReport
}

func LoadConfigured(ctx context.Context, root string) (LoadedData, error) {
	config, err := RuntimeConfigFromEnv(os.Getenv)
	if err != nil {
		return LoadedData{}, err
	}
	if config.Mode == ModeMyData {
		if err := RequireActiveConsent(root); err != nil {
			return LoadedData{}, err
		}
		transactions, err := (LiveClient{Config: config}).Fetch(ctx)
		if err != nil {
			return LoadedData{}, err
		}
		report := AssessLiveParity(transactions)
		report.FixturePath = myDataTransactionsPath
		return LoadedData{Transactions: transactions, Parity: report}, nil
	}
	transactions, err := LoadFixture(root)
	if err != nil {
		return LoadedData{}, err
	}
	report := AssessParity(transactions)
	report.FixturePath = FixtureRelativePath
	return LoadedData{Transactions: transactions, Parity: report}, nil
}

func scanFixture(file *os.File) ([]SourceTransaction, error) {
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

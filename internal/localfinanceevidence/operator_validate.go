package localfinanceevidence

import (
	"encoding/json"
	"errors"
	"slices"
)

func validateOperatorReport(value OperatorReport, month string, ref ProofRef) error {
	if value.SchemaVersion != operatorProofSchema || value.Component != "finance-operator" ||
		value.ExecutionMode != "local_subprocess_emulation" || value.Month != month || !value.LoopbackOnly ||
		value.CredentialsRead || value.ExternalNetworkEnabled || value.ExternalWritesEnabled || value.FinancialActions ||
		value.ChannelWrites || value.ServiceInstalled || value.ReportHash != ref.ReportHash {
		return errors.New("Finance Operator execution proof boundary is invalid")
	}
	if !value.Day2LedgerCompleted || value.Day3Completed || value.Day3FailedStage != "portfolio" ||
		!value.Day3SkippedLedger || !value.Day5ResumeCompleted || !value.Day5SkippedLedger ||
		!value.Day6ReplayCompleted || !value.Day6AllStagesSkipped || value.RetryExitCode != 1 || value.MaxAttemptsPerRun != 3 ||
		value.CollectorCalls != (OperatorCalls{Ledger: 2, Portfolio: 4, Revenue: 1}) ||
		value.SummaryCalls != (OperatorCalls{Ledger: 1, Portfolio: 1, Revenue: 1}) || value.StageRows != 3 || value.SnapshotRows != 1 {
		return errors.New("Finance Operator checkpoint or retry proof is invalid")
	}
	if err := validateOperatorSnapshot(value.Snapshot, month); err != nil {
		return err
	}
	required := []string{"aggregate-only-snapshot", "ambient-environment-sanitized", "binary-sha-pins-enforced", "bounded-child-retry", "completed-stage-skip", "daily-catch-up-order", "failure-checkpoint-resume", "fixed-argv-no-shell", "raw-child-output-discarded", "single-snapshot-replay"}
	if !slices.Equal(value.Checks, required) {
		return errors.New("Finance Operator execution proof checks are invalid")
	}
	copy := value
	copy.ReportHash = ""
	body, err := json.Marshal(copy)
	if err != nil || value.ReportHash != digest(string(body)) {
		return errors.New("Finance Operator execution report hash is invalid")
	}
	return nil
}

func validateOperatorSnapshot(value OperatorSnapshot, month string) error {
	if value.SchemaVersion != "myhome.finance-monthly-snapshot/v1" || value.Month != month || value.Currency != "KRW" ||
		value.RevenueStatus != "estimated" || value.TransactionCount != 3 || value.CardSpendMinor != 20900 ||
		value.CardRefundMinor != 2200 || value.NetCardSpendMinor != 18700 || value.LedgerIncomeMinor != 8300 ||
		value.LedgerCostMinor != 2000 || value.LedgerNetIncomeMinor != 6300 || value.YouTubeGrossMinor != 8300 ||
		value.YouTubeCostMinor != 2000 || value.YouTubeNetMinor != 6300 || value.TrackedSurplusMinor != -12400 ||
		value.AssetAsOf != "2026-08-03T00:00:00Z" || value.AssetCashMinor != 50000 || value.AssetSecuritiesMinor != 150000 ||
		value.AssetProfitLossMinor != 10000 || value.AssetTotalMinor != 200000 || value.LiquidityRatioBPS != 2500 ||
		value.CreditSpendToAssetBPS != 935 {
		return errors.New("Finance Operator monthly snapshot metrics are invalid")
	}
	if value.CardSpendMinor-value.CardRefundMinor != value.NetCardSpendMinor ||
		value.LedgerIncomeMinor-value.LedgerCostMinor != value.LedgerNetIncomeMinor ||
		value.YouTubeGrossMinor-value.YouTubeCostMinor != value.YouTubeNetMinor ||
		value.LedgerNetIncomeMinor-value.NetCardSpendMinor != value.TrackedSurplusMinor ||
		value.AssetCashMinor+value.AssetSecuritiesMinor != value.AssetTotalMinor {
		return errors.New("Finance Operator monthly snapshot does not reconcile")
	}
	for _, hash := range []string{value.SourceHashes.Ledger, value.SourceHashes.Portfolio, value.SourceHashes.Revenue, value.SnapshotHash} {
		if !hashPattern.MatchString(hash) {
			return errors.New("Finance Operator monthly snapshot hash is invalid")
		}
	}
	required := []string{"asset-total-reconciled", "credit-refunds-netted", "month-and-asset-as-of-separated", "raw-child-output-not-persisted", "revenue-costs-netted", "source-output-hashes-bound"}
	if !slices.Equal(value.Checks, required) {
		return errors.New("Finance Operator monthly snapshot checks are invalid")
	}
	copy := value
	copy.SnapshotHash = ""
	body, err := json.Marshal(copy)
	if err != nil || value.SnapshotHash != digest(string(body)) {
		return errors.New("Finance Operator monthly snapshot hash changed")
	}
	return nil
}

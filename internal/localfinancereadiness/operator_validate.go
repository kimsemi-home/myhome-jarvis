package localfinancereadiness

import (
	"encoding/json"
	"errors"
	"slices"
)

func validateOperatorPlan(value OperatorPlan, ref Ref) error {
	if value.SchemaVersion != OperatorPlanSchema || value.Component != "finance-operator" || value.Component != ref.Component ||
		value.ExecutionMode != "plan_only" || value.CredentialsRead || value.ExternalNetworkEnabled || value.ExternalWritesEnabled ||
		value.InstallAllowed || value.PublicConfigSource != "config/operator.json" || value.KeychainService != "myhome-finance-operator" ||
		value.APITokenAccount != "local-reader-token" || value.Timezone != "Asia/Seoul" ||
		value.DailySchedule != (OperatorSchedule{Hour: 8, Minute: 10}) || len(value.Stages) != 3 {
		return errors.New("Finance Operator readiness boundary is invalid")
	}
	expected := []OperatorStage{
		{Component: "ledger", BinaryPath: "__LEDGERCTL__", WorkingDirectory: "__LEDGER_ROOT__", DueDay: 2, BinarySHA256: "__LEDGERCTL_SHA256__"},
		{Component: "portfolio", BinaryPath: "__PORTFOLIOCTL__", WorkingDirectory: "__PORTFOLIO_ROOT__", DueDay: 3, BinarySHA256: "__PORTFOLIOCTL_SHA256__"},
		{Component: "revenue", BinaryPath: "__REVENUECTL__", WorkingDirectory: "__REVENUE_ROOT__", DueDay: 5, BinarySHA256: "__REVENUECTL_SHA256__"},
	}
	if !slices.Equal(value.Stages, expected) || value.Launchd.TemplatePath != "ops/launchd/com.kimsemi.myhome-finance-operator.daily.plist.template" ||
		value.Launchd.Label != "com.kimsemi.myhome-finance-operator.daily" ||
		!slices.Equal(value.Launchd.ProgramArguments, []string{"__FINOPCTL__", "run", "due"}) ||
		!hashPattern.MatchString(value.Launchd.TemplateHash) {
		return errors.New("Finance Operator readiness execution contract is invalid")
	}
	required := []string{"ambient-environment-sanitized", "bounded-child-retry", "child-argv-fixed", "credentials-not-read", "daily-catch-up-enabled", "external-network-disabled", "external-writes-disabled", "launchd-not-installed", "private-config-ignored", "stage-order-fixed"}
	if !slices.Equal(value.Checks, required) || value.PlanHash != ref.PlanHash || value.PlanHash != operatorPlanHash(value) {
		return errors.New("Finance Operator readiness plan hash is invalid")
	}
	return nil
}

func operatorPlanHash(value OperatorPlan) string {
	value.PlanHash = ""
	body, _ := json.Marshal(value)
	return digest(body)
}

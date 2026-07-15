package localfinancereadiness

import (
	"errors"
	"slices"
)

func validateStages(stages []Stage) error {
	expected := []Stage{
		{Position: 1, Component: "ledger", Executor: "finance-operator", Day: 2, Hour: 8, Minute: 10, DependsOn: []string{}, Action: "collect-monthly"},
		{Position: 2, Component: "portfolio", Executor: "finance-operator", Day: 3, Hour: 8, Minute: 10, DependsOn: []string{"ledger"}, Action: "collect-monthly"},
		{Position: 3, Component: "revenue", Executor: "finance-operator", Day: 5, Hour: 8, Minute: 10, DependsOn: []string{"ledger", "portfolio"}, Action: "collect-monthly"},
		{Position: 4, Component: "finance-operator", Executor: "finance-operator", Day: 5, Hour: 8, Minute: 10, DependsOn: []string{"ledger", "portfolio", "revenue"}, Action: "aggregate-monthly-snapshot"},
		{Position: 5, Component: "jarvis", Executor: "jarvis", Day: 5, Hour: 10, Minute: 0, DependsOn: []string{"finance-operator"}, Action: "verify-indirect-evidence"},
	}
	for index, value := range stages {
		want := expected[index]
		if value.Position != want.Position || value.Component != want.Component || value.Executor != want.Executor || value.Day != want.Day ||
			value.Hour != want.Hour || value.Minute != want.Minute || value.Action != want.Action ||
			!slices.Equal(value.DependsOn, want.DependsOn) {
			return errors.New("monthly finance stages are not in the fail-closed order")
		}
	}
	return nil
}

package localfinancereadiness

import (
	"errors"
	"slices"
)

func validateStages(stages []Stage) error {
	expected := []Stage{
		{Position: 1, Component: "ledger", Day: 2, Hour: 7, Minute: 0, DependsOn: []string{}, Action: "collect-monthly"},
		{Position: 2, Component: "portfolio", Day: 3, Hour: 7, Minute: 20, DependsOn: []string{"ledger"}, Action: "collect-monthly"},
		{Position: 3, Component: "revenue", Day: 5, Hour: 7, Minute: 40, DependsOn: []string{"ledger", "portfolio"}, Action: "collect-monthly"},
		{Position: 4, Component: "jarvis", Day: 5, Hour: 8, Minute: 0, DependsOn: []string{"ledger", "portfolio", "revenue"}, Action: "verify-indirect-evidence"},
	}
	for index, value := range stages {
		want := expected[index]
		if value.Position != want.Position || value.Component != want.Component || value.Day != want.Day ||
			value.Hour != want.Hour || value.Minute != want.Minute || value.Action != want.Action ||
			!slices.Equal(value.DependsOn, want.DependsOn) {
			return errors.New("monthly finance stages are not in the fail-closed order")
		}
	}
	return nil
}

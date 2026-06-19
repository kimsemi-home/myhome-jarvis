package commands

func argvPlan(name string, label string, argv []string) Plan {
	return Plan{
		Name:        name,
		DryRun:      true,
		Invocations: []Invocation{{Label: label, Argv: argv}},
	}
}

func appleScriptPlan(name string, script string) Plan {
	return argvPlan(name, name, []string{"osascript", "-e", script})
}

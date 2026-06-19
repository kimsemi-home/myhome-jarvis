package commands

import "fmt"

func buildVolumeCommand(command string, payload []byte) (Plan, error) {
	switch command {
	case "volume_set":
		var body struct {
			Level *int `json:"level"`
		}
		if err := decodePayload(payload, &body); err != nil {
			return Plan{}, err
		}
		if body.Level == nil {
			return Plan{}, fmt.Errorf("%w: level is required", errInvalidPayload)
		}
		return volumeSetPlan(command, *body.Level)
	case "volume_up":
		return volumeStepPlan(command, payload, "+")
	case "volume_down":
		return volumeStepPlan(command, payload, "-")
	case "volume_mute":
		return appleScriptPlan(command, "set volume output muted true"), nil
	default:
		return Plan{}, fmt.Errorf("unknown volume command %q", command)
	}
}

func volumeSetPlan(name string, level int) (Plan, error) {
	if level < 0 || level > 100 {
		return Plan{}, fmt.Errorf("%w: level must be between 0 and 100", errInvalidPayload)
	}
	return appleScriptPlan(name, fmt.Sprintf("set volume output volume %d", level)), nil
}

func volumeStepPlan(command string, payload []byte, operator string) (Plan, error) {
	step, err := volumeStep(payload)
	if err != nil {
		return Plan{}, err
	}
	script := fmt.Sprintf("set volume output volume ((output volume of (get volume settings)) %s %d)", operator, step)
	return appleScriptPlan(command, script), nil
}

func volumeStep(payload []byte) (int, error) {
	var body struct {
		Step *int `json:"step"`
	}
	if err := decodePayload(payload, &body); err != nil {
		return 0, err
	}
	step := 10
	if body.Step != nil {
		step = *body.Step
	}
	if step < 1 || step > 100 {
		return 0, fmt.Errorf("%w: step must be between 1 and 100", errInvalidPayload)
	}
	return step, nil
}

use crate::json::json_i32_optional;
use crate::strings::apple_script_plan;
use crate::{CommandError, Plan};

pub fn volume_set(level: i32) -> Result<Plan, CommandError> {
    if !(0..=100).contains(&level) {
        return Err(CommandError::InvalidPayload(
            "level must be between 0 and 100".to_string(),
        ));
    }
    Ok(apple_script_plan(
        "volume_set",
        &format!("set volume output volume {level}"),
    ))
}

pub(crate) fn volume_step_plan(
    name: &str,
    payload: &str,
    sign: char,
) -> Result<Plan, CommandError> {
    let step = json_i32_optional(payload, "step")?.unwrap_or(10);
    if !(1..=100).contains(&step) {
        return Err(CommandError::InvalidPayload(
            "step must be between 1 and 100".to_string(),
        ));
    }
    Ok(apple_script_plan(
        name,
        &format!(
            "set volume output volume ((output volume of (get volume settings)) {sign} {step})"
        ),
    ))
}

pub(crate) fn mute_plan() -> Result<Plan, CommandError> {
    Ok(apple_script_plan(
        "volume_mute",
        "set volume output muted true",
    ))
}

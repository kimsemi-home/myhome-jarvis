use crate::CommandError;

pub(crate) fn json_raw_value<'a>(payload: &'a str, key: &str) -> Result<&'a str, CommandError> {
    json_raw_value_optional(payload, key)?
        .ok_or_else(|| CommandError::InvalidPayload(format!("{key} is required")))
}

pub(crate) fn json_raw_value_optional<'a>(
    payload: &'a str,
    key: &str,
) -> Result<Option<&'a str>, CommandError> {
    let needle = format!("\"{key}\"");
    let Some(key_index) = payload.find(&needle) else {
        return Ok(None);
    };
    let after_key = &payload[key_index + needle.len()..];
    let Some(colon_index) = after_key.find(':') else {
        return Err(CommandError::InvalidPayload(format!(
            "{key} is missing ':'"
        )));
    };
    Ok(Some(after_key[colon_index + 1..].trim_start()))
}

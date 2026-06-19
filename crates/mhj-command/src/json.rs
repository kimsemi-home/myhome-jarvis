use crate::json_raw::{json_raw_value, json_raw_value_optional};
use crate::CommandError;

pub(crate) fn json_string(payload: &str, key: &str) -> Result<String, CommandError> {
    let value = json_raw_value(payload, key)?;
    if !value.starts_with('"') {
        return Err(CommandError::InvalidPayload(format!(
            "{key} must be a string"
        )));
    }
    read_json_string(value).ok_or_else(|| CommandError::InvalidPayload(format!("{key} is invalid")))
}

pub(crate) fn json_i32(payload: &str, key: &str) -> Result<i32, CommandError> {
    json_i32_optional(payload, key)?
        .ok_or_else(|| CommandError::InvalidPayload(format!("{key} is required")))
}

pub(crate) fn json_i32_optional(payload: &str, key: &str) -> Result<Option<i32>, CommandError> {
    let Some(value) = json_raw_value_optional(payload, key)? else {
        return Ok(None);
    };
    let number: String = value
        .chars()
        .take_while(|character| character.is_ascii_digit() || *character == '-')
        .collect();
    if number.is_empty() || number == "-" {
        return Err(CommandError::InvalidPayload(format!(
            "{key} must be an integer"
        )));
    }
    number
        .parse::<i32>()
        .map(Some)
        .map_err(|_| CommandError::InvalidPayload(format!("{key} must be an integer")))
}

fn read_json_string(value: &str) -> Option<String> {
    let mut chars = value.chars();
    if chars.next()? != '"' {
        return None;
    }
    let mut out = String::new();
    let mut escaped = false;
    for character in chars {
        if escaped {
            out.push(unescape_json_char(character));
            escaped = false;
            continue;
        }
        match character {
            '\\' => escaped = true,
            '"' => return Some(out),
            other => out.push(other),
        }
    }
    None
}

fn unescape_json_char(character: char) -> char {
    match character {
        '"' => '"',
        '\\' => '\\',
        '/' => '/',
        'n' => '\n',
        'r' => '\r',
        't' => '\t',
        other => other,
    }
}

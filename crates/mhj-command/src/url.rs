use crate::CommandError;

pub fn validate_http_url(raw: &str) -> Result<String, CommandError> {
    let trimmed = raw.trim();
    let lower = trimmed.to_ascii_lowercase();
    let remainder = if lower.starts_with("https://") {
        &trimmed[8..]
    } else if lower.starts_with("http://") {
        &trimmed[7..]
    } else {
        return Err(CommandError::InvalidPayload(
            "only http and https URLs are allowed".to_string(),
        ));
    };
    let host = remainder
        .split(['/', '?', '#'])
        .next()
        .unwrap_or_default()
        .trim();
    if host.is_empty() {
        return Err(CommandError::InvalidPayload(
            "URL host is required".to_string(),
        ));
    }
    Ok(trimmed.to_string())
}

pub(crate) fn encode_query(value: &str) -> String {
    let mut out = String::new();
    for byte in value.bytes() {
        match byte {
            b'A'..=b'Z' | b'a'..=b'z' | b'0'..=b'9' | b'-' | b'_' | b'.' | b'~' => {
                out.push(byte as char)
            }
            b' ' => out.push('+'),
            other => out.push_str(&format!("%{other:02X}")),
        }
    }
    out
}

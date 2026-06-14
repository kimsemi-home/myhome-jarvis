use std::error::Error;
use std::fmt;

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct Invocation {
    pub label: String,
    pub argv: Vec<String>,
    pub url: Option<String>,
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct Plan {
    pub name: String,
    pub dry_run: bool,
    pub invocations: Vec<Invocation>,
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub enum CommandError {
    UnknownCommand(String),
    InvalidPayload(String),
}

impl fmt::Display for CommandError {
    fn fmt(&self, formatter: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            CommandError::UnknownCommand(name) => write!(formatter, "unknown command {name}"),
            CommandError::InvalidPayload(message) => {
                write!(formatter, "invalid payload: {message}")
            }
        }
    }
}

impl Error for CommandError {}

pub fn plan_for(command: &str, payload: &str) -> Result<Plan, CommandError> {
    match normalize_command(command).as_str() {
        "open_coupang_play" => ott_shortcut_plan("open_coupang_play", "coupangplay"),
        "open_disney_plus" => ott_shortcut_plan("open_disney_plus", "disney"),
        "open_netflix" => ott_shortcut_plan("open_netflix", "netflix"),
        "open_youtube" => Ok(open_url_plan("open_youtube", "https://www.youtube.com")),
        "open_youtube_search" => {
            let query = json_string(payload, "query")?;
            if query.trim().is_empty() {
                return Err(CommandError::InvalidPayload(
                    "query is required".to_string(),
                ));
            }
            Ok(open_url_plan(
                "open_youtube_search",
                &format!(
                    "https://www.youtube.com/results?search_query={}",
                    encode_query(query.trim())
                ),
            ))
        }
        "open_tving" => ott_shortcut_plan("open_tving", "tving"),
        "open_wavve" => ott_shortcut_plan("open_wavve", "wavve"),
        "open_ott" => {
            let service = json_string(payload, "service")?.to_ascii_lowercase();
            let url = ott_url(&service).ok_or_else(|| {
                CommandError::InvalidPayload(format!("unsupported ott service {service}"))
            })?;
            Ok(open_url_plan("open_ott", url))
        }
        "open_url" => {
            let url = validate_http_url(&json_string(payload, "url")?)?;
            Ok(open_url_plan("open_url", &url))
        }
        "volume_set" => volume_set(json_i32(payload, "level")?),
        "volume_up" => volume_step_plan("volume_up", payload, '+'),
        "volume_down" => volume_step_plan("volume_down", payload, '-'),
        "volume_mute" => Ok(apple_script_plan(
            "volume_mute",
            "set volume output muted true",
        )),
        "display_sleep" => Ok(argv_plan(
            "display_sleep",
            "display_sleep",
            &["pmset", "displaysleepnow"],
        )),
        "mac_sleep" => Ok(argv_plan("mac_sleep", "mac_sleep", &["pmset", "sleepnow"])),
        "movie_mode" => Ok(Plan {
            name: "movie_mode".to_string(),
            dry_run: true,
            invocations: vec![
                Invocation {
                    label: "movie_volume".to_string(),
                    argv: strings(&["osascript", "-e", "set volume output volume 35"]),
                    url: None,
                },
                Invocation {
                    label: "open_youtube".to_string(),
                    argv: strings(&["open", "https://www.youtube.com"]),
                    url: Some("https://www.youtube.com".to_string()),
                },
            ],
        }),
        "sleep_mode" => Ok(Plan {
            name: "sleep_mode".to_string(),
            dry_run: true,
            invocations: vec![
                Invocation {
                    label: "mute".to_string(),
                    argv: strings(&["osascript", "-e", "set volume output muted true"]),
                    url: None,
                },
                Invocation {
                    label: "display_sleep".to_string(),
                    argv: strings(&["pmset", "displaysleepnow"]),
                    url: None,
                },
            ],
        }),
        other => Err(CommandError::UnknownCommand(other.to_string())),
    }
}

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

pub fn plan_text(plan: &Plan) -> String {
    let mut text = plan.name.clone();
    for invocation in &plan.invocations {
        text.push(' ');
        text.push_str(&invocation.label);
        if let Some(url) = &invocation.url {
            text.push(' ');
            text.push_str(url);
        }
        for arg in &invocation.argv {
            text.push(' ');
            text.push_str(arg);
        }
    }
    text
}

fn volume_step_plan(name: &str, payload: &str, sign: char) -> Result<Plan, CommandError> {
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

fn open_url_plan(name: &str, target: &str) -> Plan {
    Plan {
        name: name.to_string(),
        dry_run: true,
        invocations: vec![Invocation {
            label: name.to_string(),
            argv: strings(&["open", target]),
            url: Some(target.to_string()),
        }],
    }
}

fn ott_shortcut_plan(name: &str, service: &str) -> Result<Plan, CommandError> {
    let target = ott_url(service).ok_or_else(|| {
        CommandError::InvalidPayload(format!("missing OTT shortcut target for {service}"))
    })?;
    Ok(open_url_plan(name, target))
}

fn apple_script_plan(name: &str, script: &str) -> Plan {
    argv_plan(name, name, &["osascript", "-e", script])
}

fn argv_plan(name: &str, label: &str, argv: &[&str]) -> Plan {
    Plan {
        name: name.to_string(),
        dry_run: true,
        invocations: vec![Invocation {
            label: label.to_string(),
            argv: strings(argv),
            url: None,
        }],
    }
}

fn ott_url(service: &str) -> Option<&'static str> {
    match service {
        "coupangplay" => Some("https://www.coupangplay.com"),
        "disney" => Some("https://www.disneyplus.com"),
        "netflix" => Some("https://www.netflix.com"),
        "tving" => Some("https://www.tving.com"),
        "wavve" => Some("https://www.wavve.com"),
        "youtube" => Some("https://www.youtube.com"),
        _ => None,
    }
}

fn normalize_command(command: &str) -> String {
    command.trim().to_ascii_lowercase().replace('-', "_")
}

fn strings(values: &[&str]) -> Vec<String> {
    values.iter().map(|value| (*value).to_string()).collect()
}

fn json_string(payload: &str, key: &str) -> Result<String, CommandError> {
    let value = json_raw_value(payload, key)?;
    if !value.starts_with('"') {
        return Err(CommandError::InvalidPayload(format!(
            "{key} must be a string"
        )));
    }
    read_json_string(value).ok_or_else(|| CommandError::InvalidPayload(format!("{key} is invalid")))
}

fn json_i32(payload: &str, key: &str) -> Result<i32, CommandError> {
    json_i32_optional(payload, key)?
        .ok_or_else(|| CommandError::InvalidPayload(format!("{key} is required")))
}

fn json_i32_optional(payload: &str, key: &str) -> Result<Option<i32>, CommandError> {
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

fn json_raw_value<'a>(payload: &'a str, key: &str) -> Result<&'a str, CommandError> {
    json_raw_value_optional(payload, key)?
        .ok_or_else(|| CommandError::InvalidPayload(format!("{key} is required")))
}

fn json_raw_value_optional<'a>(
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

fn read_json_string(value: &str) -> Option<String> {
    let mut chars = value.chars();
    if chars.next()? != '"' {
        return None;
    }
    let mut out = String::new();
    let mut escaped = false;
    for character in chars {
        if escaped {
            out.push(match character {
                '"' => '"',
                '\\' => '\\',
                '/' => '/',
                'n' => '\n',
                'r' => '\r',
                't' => '\t',
                other => other,
            });
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

fn encode_query(value: &str) -> String {
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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn open_youtube_is_deterministic() {
        let plan = plan_for("open-youtube", "{}").expect("plan");
        assert_eq!(
            plan.invocations[0].argv,
            strings(&["open", "https://www.youtube.com"])
        );
    }

    #[test]
    fn youtube_search_encodes_query() {
        let plan = plan_for("open-youtube-search", r#"{"query":"lofi music"}"#).expect("plan");
        assert!(plan_text(&plan).contains("search_query=lofi+music"));
    }

    #[test]
    fn rejects_unknown_ott() {
        assert!(plan_for("open-ott", r#"{"service":"unknown"}"#).is_err());
    }

    #[test]
    fn ott_shortcuts_use_fixed_targets() {
        let cases = [
            ("open-netflix", "https://www.netflix.com"),
            ("open-disney-plus", "https://www.disneyplus.com"),
            ("open-tving", "https://www.tving.com"),
            ("open-wavve", "https://www.wavve.com"),
            ("open-coupang-play", "https://www.coupangplay.com"),
        ];
        for (command, expected_url) in cases {
            let plan = plan_for(command, "{}").expect("shortcut plan");
            assert_eq!(plan.invocations[0].url.as_deref(), Some(expected_url));
            assert_eq!(plan.invocations[0].argv, strings(&["open", expected_url]));
        }
    }

    #[test]
    fn rejects_unsafe_url_scheme() {
        assert!(plan_for("open-url", r#"{"url":"javascript:alert(1)"}"#).is_err());
    }

    #[test]
    fn rejects_out_of_range_volume() {
        assert!(plan_for("volume-set", r#"{"level":101}"#).is_err());
        assert!(plan_for("volume-set", r#"{"level":-1}"#).is_err());
    }

    #[test]
    fn accepts_display_sleep() {
        let plan = plan_for("display-sleep", "{}").expect("plan");
        assert_eq!(
            plan.invocations[0].argv,
            strings(&["pmset", "displaysleepnow"])
        );
    }
}

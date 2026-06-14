use mhj_command::{plan_for, plan_text};

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct HarnessCase {
    pub name: &'static str,
    pub command: &'static str,
    pub payload: &'static str,
    pub should_pass: bool,
    pub contains: &'static str,
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct HarnessResult {
    pub name: &'static str,
    pub passed: bool,
    pub message: String,
}

pub fn home_cases() -> Vec<HarnessCase> {
    vec![
        HarnessCase {
            name: "open_youtube empty payload success",
            command: "open-youtube",
            payload: "{}",
            should_pass: true,
            contains: "https://www.youtube.com",
        },
        HarnessCase {
            name: "open_youtube_search lofi music success",
            command: "open-youtube-search",
            payload: r#"{"query":"lofi music"}"#,
            should_pass: true,
            contains: "search_query=lofi+music",
        },
        HarnessCase {
            name: "open_ott netflix success",
            command: "open-ott",
            payload: r#"{"service":"netflix"}"#,
            should_pass: true,
            contains: "https://www.netflix.com",
        },
        HarnessCase {
            name: "open_ott unknown fail",
            command: "open-ott",
            payload: r#"{"service":"unknown"}"#,
            should_pass: false,
            contains: "",
        },
        HarnessCase {
            name: "volume_set 30 success",
            command: "volume-set",
            payload: r#"{"level":30}"#,
            should_pass: true,
            contains: "30",
        },
        HarnessCase {
            name: "volume_set 101 fail",
            command: "volume-set",
            payload: r#"{"level":101}"#,
            should_pass: false,
            contains: "",
        },
        HarnessCase {
            name: "display_sleep success",
            command: "display-sleep",
            payload: "{}",
            should_pass: true,
            contains: "displaysleepnow",
        },
        HarnessCase {
            name: "open_url javascript fail",
            command: "open-url",
            payload: r#"{"url":"javascript:alert(1)"}"#,
            should_pass: false,
            contains: "",
        },
    ]
}

pub fn run_home_harness() -> Vec<HarnessResult> {
    home_cases()
        .into_iter()
        .map(|case| {
            let result = plan_for(case.command, case.payload);
            match (case.should_pass, result) {
                (true, Ok(plan)) => {
                    let text = plan_text(&plan);
                    if case.contains.is_empty() || text.contains(case.contains) {
                        HarnessResult {
                            name: case.name,
                            passed: true,
                            message: "ok".to_string(),
                        }
                    } else {
                        HarnessResult {
                            name: case.name,
                            passed: false,
                            message: format!("missing {}", case.contains),
                        }
                    }
                }
                (true, Err(error)) => HarnessResult {
                    name: case.name,
                    passed: false,
                    message: error.to_string(),
                },
                (false, Ok(_)) => HarnessResult {
                    name: case.name,
                    passed: false,
                    message: "expected failure".to_string(),
                },
                (false, Err(error)) => HarnessResult {
                    name: case.name,
                    passed: true,
                    message: format!("failed safely: {error}"),
                },
            }
        })
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn home_harness_passes() {
        let results = run_home_harness();
        assert!(results.iter().all(|result| result.passed), "{results:#?}");
    }
}

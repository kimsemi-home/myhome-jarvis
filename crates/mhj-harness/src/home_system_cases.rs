use crate::HarnessCase;

pub(crate) fn home_system_cases() -> Vec<HarnessCase> {
    vec![
        HarnessCase {
            name: "display_sleep success",
            command: "display-sleep",
            payload: "{}",
            should_pass: true,
            contains: "displaysleepnow",
        },
        HarnessCase {
            name: "open_url https success",
            command: "open-url",
            payload: r#"{"url":"https://example.com"}"#,
            should_pass: true,
            contains: "https://example.com",
        },
        HarnessCase {
            name: "open_url javascript fail",
            command: "open-url",
            payload: r#"{"url":"javascript:alert(1)"}"#,
            should_pass: false,
            contains: "",
        },
        HarnessCase {
            name: "movie_mode dry-run success",
            command: "movie-mode",
            payload: "{}",
            should_pass: true,
            contains: "movie_volume",
        },
        HarnessCase {
            name: "sleep_mode dry-run success",
            command: "sleep-mode",
            payload: "{}",
            should_pass: true,
            contains: "display_sleep",
        },
    ]
}

use crate::HarnessCase;

pub(crate) fn home_media_cases() -> Vec<HarnessCase> {
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
    ]
}

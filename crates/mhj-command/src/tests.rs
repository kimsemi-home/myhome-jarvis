use crate::*;

#[test]
fn open_youtube_is_deterministic() {
    let plan = plan_for("open-youtube", "{}").expect("plan");
    assert_eq!(
        plan.invocations[0].argv,
        vec!["open", "https://www.youtube.com"]
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
        assert_eq!(plan.invocations[0].argv, vec!["open", expected_url]);
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
    assert_eq!(plan.invocations[0].argv, vec!["pmset", "displaysleepnow"]);
}

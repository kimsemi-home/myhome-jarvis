use crate::HarnessCase;

pub(crate) fn home_ott_cases() -> Vec<HarnessCase> {
    vec![
        ott_case(
            "open_netflix shortcut success",
            "open-netflix",
            "https://www.netflix.com",
        ),
        ott_case(
            "open_disney_plus shortcut success",
            "open-disney-plus",
            "https://www.disneyplus.com",
        ),
        ott_case(
            "open_tving shortcut success",
            "open-tving",
            "https://www.tving.com",
        ),
        ott_case(
            "open_wavve shortcut success",
            "open-wavve",
            "https://www.wavve.com",
        ),
        ott_case(
            "open_coupang_play shortcut success",
            "open-coupang-play",
            "https://www.coupangplay.com",
        ),
    ]
}

fn ott_case(name: &'static str, command: &'static str, contains: &'static str) -> HarnessCase {
    HarnessCase {
        name,
        command,
        payload: "{}",
        should_pass: true,
        contains,
    }
}

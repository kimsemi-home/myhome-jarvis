use crate::strings::{argv_plan, strings};
use crate::{Invocation, Plan};

pub(crate) fn movie_mode_plan() -> Plan {
    Plan {
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
    }
}

pub(crate) fn sleep_mode_plan() -> Plan {
    Plan {
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
    }
}

pub(crate) fn display_sleep_plan() -> Plan {
    argv_plan(
        "display_sleep",
        "display_sleep",
        &["pmset", "displaysleepnow"],
    )
}

pub(crate) fn mac_sleep_plan() -> Plan {
    argv_plan("mac_sleep", "mac_sleep", &["pmset", "sleepnow"])
}

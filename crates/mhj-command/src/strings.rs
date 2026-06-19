use crate::{Invocation, Plan};

pub(crate) fn strings(values: &[&str]) -> Vec<String> {
    values.iter().map(|value| (*value).to_string()).collect()
}

pub(crate) fn open_url_plan(name: &str, target: &str) -> Plan {
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

pub(crate) fn apple_script_plan(name: &str, script: &str) -> Plan {
    argv_plan(name, name, &["osascript", "-e", script])
}

pub(crate) fn argv_plan(name: &str, label: &str, argv: &[&str]) -> Plan {
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

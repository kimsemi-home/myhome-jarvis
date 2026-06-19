use crate::HarnessCase;

pub(crate) fn home_control_cases() -> Vec<HarnessCase> {
    vec![
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
            name: "volume_up step 10 success",
            command: "volume-up",
            payload: r#"{"step":10}"#,
            should_pass: true,
            contains: "+ 10",
        },
        HarnessCase {
            name: "volume_down step 10 success",
            command: "volume-down",
            payload: r#"{"step":10}"#,
            should_pass: true,
            contains: "- 10",
        },
    ]
}

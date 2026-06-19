use crate::HarnessResult;

pub(crate) fn check(name: &'static str, passed: bool, message: impl Into<String>) -> HarnessResult {
    HarnessResult {
        name,
        passed,
        message: if passed {
            "ok".to_string()
        } else {
            message.into()
        },
    }
}

use crate::home_control_cases::home_control_cases;
use crate::home_media_cases::home_media_cases;
use crate::home_ott_cases::home_ott_cases;
use crate::home_system_cases::home_system_cases;
use crate::HarnessCase;

pub fn home_cases() -> Vec<HarnessCase> {
    let mut cases = Vec::new();
    cases.extend(home_media_cases());
    cases.extend(home_ott_cases());
    cases.extend(home_control_cases());
    cases.extend(home_system_cases());
    cases
}

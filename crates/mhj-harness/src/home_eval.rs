use crate::{HarnessCase, HarnessResult};
use mhj_command::{plan_for, plan_text};

pub(crate) fn run_home_case(case: HarnessCase) -> HarnessResult {
    let result = plan_for(case.command, case.payload);
    match (case.should_pass, result) {
        (true, Ok(plan)) => success_case_result(case, plan_text(&plan)),
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
}

fn success_case_result(case: HarnessCase, text: String) -> HarnessResult {
    if case.contains.is_empty() || text.contains(case.contains) {
        return HarnessResult {
            name: case.name,
            passed: true,
            message: "ok".to_string(),
        };
    }
    HarnessResult {
        name: case.name,
        passed: false,
        message: format!("missing {}", case.contains),
    }
}

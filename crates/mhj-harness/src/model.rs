#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub struct HarnessCase {
    pub name: &'static str,
    pub command: &'static str,
    pub payload: &'static str,
    pub should_pass: bool,
    pub contains: &'static str,
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct HarnessResult {
    pub name: &'static str,
    pub passed: bool,
    pub message: String,
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct HarnessReport {
    pub name: &'static str,
    pub passed: bool,
    pub results: Vec<HarnessResult>,
}

impl HarnessReport {
    pub fn new(name: &'static str, results: Vec<HarnessResult>) -> Self {
        let passed = results.iter().all(|result| result.passed);
        Self {
            name,
            passed,
            results,
        }
    }
}

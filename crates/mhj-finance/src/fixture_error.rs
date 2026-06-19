use std::error::Error;
use std::fmt;

#[derive(Debug, Clone, PartialEq, Eq)]
pub enum FixtureError {
    Json {
        line: usize,
        message: String,
    },
    Validation {
        line: usize,
        field: &'static str,
        message: String,
    },
}

impl fmt::Display for FixtureError {
    fn fmt(&self, formatter: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            FixtureError::Json { line, message } => {
                write!(formatter, "line {line}: invalid json: {message}")
            }
            FixtureError::Validation {
                line,
                field,
                message,
            } => {
                write!(formatter, "line {line}: {field}: {message}")
            }
        }
    }
}

impl Error for FixtureError {}

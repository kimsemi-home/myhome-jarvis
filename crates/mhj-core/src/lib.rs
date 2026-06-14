pub mod benchmark;
pub mod commerce;
pub mod finance;
pub mod storage;

use serde::de::DeserializeOwned;
use serde::{Deserialize, Serialize};
use std::error::Error;
use std::fmt;

pub const PROJECT_NAME: &str = "myhome-jarvis";
pub const VERSION: &str = "0.1.0-bootstrap";
pub const DEFAULT_BIND_HOST: &str = "127.0.0.1";
pub const DRY_RUN_DEFAULT: bool = true;

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum Owner {
    User,
    Spouse,
    Household,
    Unknown,
}

impl Owner {
    pub fn is_known(self) -> bool {
        !matches!(self, Owner::Unknown)
    }
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct MoneyAmount {
    pub minor_units: i64,
    pub currency: String,
}

impl MoneyAmount {
    pub fn validate_positive(&self, field: &'static str) -> Result<(), ValidationError> {
        ensure_currency(field, &self.currency)?;
        if self.minor_units <= 0 {
            return Err(ValidationError::new(
                field,
                "amount must be greater than zero",
            ));
        }
        Ok(())
    }

    pub fn validate_non_negative(&self, field: &'static str) -> Result<(), ValidationError> {
        ensure_currency(field, &self.currency)?;
        if self.minor_units < 0 {
            return Err(ValidationError::new(field, "amount must not be negative"));
        }
        Ok(())
    }
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct ValidationError {
    pub field: &'static str,
    pub message: String,
}

impl ValidationError {
    pub fn new(field: &'static str, message: impl Into<String>) -> Self {
        Self {
            field,
            message: message.into(),
        }
    }
}

impl fmt::Display for ValidationError {
    fn fmt(&self, formatter: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(formatter, "{}: {}", self.field, self.message)
    }
}

impl Error for ValidationError {}

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
            } => write!(formatter, "line {line}: {field}: {message}"),
        }
    }
}

impl Error for FixtureError {}

pub fn parse_jsonl<T, F>(input: &str, mut validate: F) -> Result<Vec<T>, FixtureError>
where
    T: DeserializeOwned,
    F: FnMut(&T) -> Result<(), ValidationError>,
{
    let mut records = Vec::new();
    for (index, line) in input.lines().enumerate() {
        let trimmed = line.trim();
        if trimmed.is_empty() {
            continue;
        }
        let record: T = serde_json::from_str(trimmed).map_err(|error| FixtureError::Json {
            line: index + 1,
            message: error.to_string(),
        })?;
        validate(&record).map_err(|error| FixtureError::Validation {
            line: index + 1,
            field: error.field,
            message: error.message,
        })?;
        records.push(record);
    }
    Ok(records)
}

pub(crate) fn ensure_non_empty(field: &'static str, value: &str) -> Result<(), ValidationError> {
    if value.trim().is_empty() {
        return Err(ValidationError::new(field, "must not be empty"));
    }
    Ok(())
}

pub(crate) fn ensure_currency(field: &'static str, value: &str) -> Result<(), ValidationError> {
    if value.len() != 3
        || !value
            .chars()
            .all(|character| character.is_ascii_uppercase())
    {
        return Err(ValidationError::new(
            field,
            "currency must be a three-letter ISO code",
        ));
    }
    Ok(())
}

pub(crate) fn ensure_iso_datetime(field: &'static str, value: &str) -> Result<(), ValidationError> {
    ensure_non_empty(field, value)?;
    let bytes = value.as_bytes();
    let valid_date_prefix = bytes.len() >= 10
        && bytes[0..4].iter().all(u8::is_ascii_digit)
        && bytes[4] == b'-'
        && bytes[5..7].iter().all(u8::is_ascii_digit)
        && bytes[7] == b'-'
        && bytes[8..10].iter().all(u8::is_ascii_digit);
    if !valid_date_prefix {
        return Err(ValidationError::new(
            field,
            "must start with an ISO-8601 date",
        ));
    }
    if bytes.len() > 10 && !value.contains('T') {
        return Err(ValidationError::new(
            field,
            "timestamp values must include 'T'",
        ));
    }
    Ok(())
}

pub(crate) fn ensure_known_owner(field: &'static str, owner: Owner) -> Result<(), ValidationError> {
    if !owner.is_known() {
        return Err(ValidationError::new(field, "owner must be known"));
    }
    Ok(())
}

pub(crate) fn ensure_tags(field: &'static str, tags: &[String]) -> Result<(), ValidationError> {
    for tag in tags {
        let trimmed = tag.trim();
        if trimmed.is_empty() {
            return Err(ValidationError::new(field, "tags must not be empty"));
        }
        if trimmed != tag {
            return Err(ValidationError::new(
                field,
                "tags must not contain outer whitespace",
            ));
        }
    }
    Ok(())
}

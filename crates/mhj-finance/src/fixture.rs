use crate::{FixtureError, TransactionIr, ValidationError};
use serde::de::DeserializeOwned;

pub const FINANCE_FIXTURE_JSONL: &str =
    include_str!("../../../fixtures/finance_transactions.jsonl");

pub fn parse_transactions_jsonl(input: &str) -> Result<Vec<TransactionIr>, FixtureError> {
    parse_jsonl(input, TransactionIr::validate)
}

pub fn fixture_transactions() -> Result<Vec<TransactionIr>, FixtureError> {
    parse_transactions_jsonl(FINANCE_FIXTURE_JSONL)
}

fn parse_jsonl<T, F>(input: &str, mut validate: F) -> Result<Vec<T>, FixtureError>
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

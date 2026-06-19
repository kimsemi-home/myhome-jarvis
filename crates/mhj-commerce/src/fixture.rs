use crate::{FixtureError, PurchaseIr, ValidationError};
use serde::de::DeserializeOwned;

pub const COMMERCE_FIXTURE_JSONL: &str = include_str!("../../../fixtures/commerce_purchases.jsonl");

pub fn parse_purchases_jsonl(input: &str) -> Result<Vec<PurchaseIr>, FixtureError> {
    parse_jsonl(input, PurchaseIr::validate)
}

pub fn fixture_purchases() -> Result<Vec<PurchaseIr>, FixtureError> {
    parse_purchases_jsonl(COMMERCE_FIXTURE_JSONL)
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

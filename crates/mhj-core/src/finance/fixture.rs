use crate::{parse_jsonl, FixtureError};

use super::TransactionIr;

pub const FINANCE_FIXTURE_JSONL: &str =
    include_str!("../../../../fixtures/finance_transactions.jsonl");

pub fn parse_transactions_jsonl(input: &str) -> Result<Vec<TransactionIr>, FixtureError> {
    parse_jsonl(input, TransactionIr::validate)
}

pub fn fixture_transactions() -> Result<Vec<TransactionIr>, FixtureError> {
    parse_transactions_jsonl(FINANCE_FIXTURE_JSONL)
}

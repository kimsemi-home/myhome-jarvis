mod model;
mod recurring;
mod validate;

#[cfg(test)]
mod tests;

use crate::{parse_jsonl, FixtureError};

pub use model::PurchaseIr;
pub use recurring::{recurring_candidates, RecurringPurchaseCandidate};

pub const COMMERCE_FIXTURE_JSONL: &str = include_str!("../../../fixtures/commerce_purchases.jsonl");

pub fn parse_purchases_jsonl(input: &str) -> Result<Vec<PurchaseIr>, FixtureError> {
    parse_jsonl(input, PurchaseIr::validate)
}

pub fn fixture_purchases() -> Result<Vec<PurchaseIr>, FixtureError> {
    parse_purchases_jsonl(COMMERCE_FIXTURE_JSONL)
}

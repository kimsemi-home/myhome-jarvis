mod candidate_model;
mod candidates;
mod cashflow;
mod fixture;
mod model;
mod validation;

pub use candidate_model::CardUsageCandidate;
pub use candidates::card_usage_candidates;
pub use cashflow::{summarize_cashflow, CashflowSummary};
pub use fixture::{fixture_transactions, parse_transactions_jsonl, FINANCE_FIXTURE_JSONL};
pub use model::{TransactionDirection, TransactionIr};

#[cfg(test)]
mod tests;

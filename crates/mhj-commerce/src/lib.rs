mod amount;
mod commerce;
mod datetime;
mod fixture;
mod fixture_error;
mod merchant;
mod owner;
mod owner_summary;
mod purchase;
mod purchase_validate;
mod recurring;
mod summary;
mod validation;
mod validation_error;

pub use amount::MoneyAmount;
pub use commerce::summarize_commerce;
pub use fixture::{fixture_purchases, parse_purchases_jsonl, COMMERCE_FIXTURE_JSONL};
pub use fixture_error::FixtureError;
pub use merchant::summarize_by_merchant;
pub use owner::Owner;
pub use owner_summary::summarize_by_owner;
pub use purchase::PurchaseIr;
pub use recurring::recurring_candidates;
pub use summary::{
    CommerceSummary, MerchantSpendSummary, OwnerSpendSummary, RecurringPurchaseCandidate,
};
pub use validation_error::ValidationError;

pub(crate) use datetime::ensure_iso_datetime;
pub(crate) use validation::{ensure_currency, ensure_known_owner, ensure_non_empty, ensure_tags};

#[cfg(test)]
mod tests;

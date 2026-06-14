use serde::de::DeserializeOwned;
use serde::{Deserialize, Serialize};
use std::error::Error;
use std::fmt;

pub const FINANCE_FIXTURE_JSONL: &str =
    include_str!("../../../fixtures/finance_transactions.jsonl");

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

    pub fn as_str(self) -> &'static str {
        match self {
            Owner::User => "user",
            Owner::Spouse => "spouse",
            Owner::Household => "household",
            Owner::Unknown => "unknown",
        }
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
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum TransactionDirection {
    Debit,
    Credit,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct TransactionIr {
    pub transaction_id: String,
    pub source: String,
    pub owner: Owner,
    pub occurred_at: String,
    pub posted_at: Option<String>,
    pub amount: MoneyAmount,
    pub direction: TransactionDirection,
    pub merchant_name: Option<String>,
    pub category: Option<String>,
    pub account_id: Option<String>,
    pub card_id: Option<String>,
    pub raw_ref: String,
    pub tags: Vec<String>,
}

impl TransactionIr {
    pub fn validate(&self) -> Result<(), ValidationError> {
        ensure_non_empty("transaction_id", &self.transaction_id)?;
        ensure_non_empty("source", &self.source)?;
        ensure_known_owner("owner", self.owner)?;
        ensure_iso_datetime("occurred_at", &self.occurred_at)?;
        if let Some(posted_at) = &self.posted_at {
            ensure_iso_datetime("posted_at", posted_at)?;
        }
        self.amount.validate_positive("amount")?;
        if self
            .account_id
            .as_deref()
            .unwrap_or_default()
            .trim()
            .is_empty()
            && self
                .card_id
                .as_deref()
                .unwrap_or_default()
                .trim()
                .is_empty()
        {
            return Err(ValidationError::new(
                "account_id",
                "account_id or card_id is required",
            ));
        }
        if matches!(self.direction, TransactionDirection::Debit)
            && self
                .merchant_name
                .as_deref()
                .unwrap_or_default()
                .trim()
                .is_empty()
        {
            return Err(ValidationError::new(
                "merchant_name",
                "merchant_name is required for debit transactions",
            ));
        }
        ensure_non_empty("raw_ref", &self.raw_ref)?;
        ensure_tags("tags", &self.tags)?;
        Ok(())
    }

    pub fn signed_minor_units(&self) -> i64 {
        match self.direction {
            TransactionDirection::Debit => -self.amount.minor_units,
            TransactionDirection::Credit => self.amount.minor_units,
        }
    }
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct CashflowSummary {
    pub records: usize,
    pub currency: String,
    pub credit_minor_units: i64,
    pub debit_minor_units: i64,
    pub net_minor_units: i64,
    pub subscription_minor_units: i64,
    pub subscription_count: usize,
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct OwnerCashflowSummary {
    pub owner: Owner,
    pub records: usize,
    pub currency: String,
    pub credit_minor_units: i64,
    pub debit_minor_units: i64,
    pub net_minor_units: i64,
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct SubscriptionCandidate {
    pub transaction_id: String,
    pub owner: Owner,
    pub merchant_name: String,
    pub currency: String,
    pub monthly_minor_units: i64,
    pub evidence_tag_count: usize,
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

pub fn parse_transactions_jsonl(input: &str) -> Result<Vec<TransactionIr>, FixtureError> {
    parse_jsonl(input, TransactionIr::validate)
}

pub fn fixture_transactions() -> Result<Vec<TransactionIr>, FixtureError> {
    parse_transactions_jsonl(FINANCE_FIXTURE_JSONL)
}

pub fn summarize_cashflow(
    transactions: &[TransactionIr],
    currency: &str,
) -> Result<CashflowSummary, ValidationError> {
    ensure_currency("currency", currency)?;
    let mut summary = CashflowSummary {
        records: 0,
        currency: currency.to_string(),
        credit_minor_units: 0,
        debit_minor_units: 0,
        net_minor_units: 0,
        subscription_minor_units: 0,
        subscription_count: 0,
    };
    for transaction in transactions {
        transaction.validate()?;
        if transaction.amount.currency != currency {
            continue;
        }
        summary.records += 1;
        match transaction.direction {
            TransactionDirection::Debit => {
                summary.debit_minor_units += transaction.amount.minor_units;
                if is_subscription(transaction) {
                    summary.subscription_minor_units += transaction.amount.minor_units;
                    summary.subscription_count += 1;
                }
            }
            TransactionDirection::Credit => {
                summary.credit_minor_units += transaction.amount.minor_units;
            }
        }
    }
    summary.net_minor_units = summary.credit_minor_units - summary.debit_minor_units;
    Ok(summary)
}

pub fn summarize_by_owner(
    transactions: &[TransactionIr],
    currency: &str,
) -> Result<Vec<OwnerCashflowSummary>, ValidationError> {
    ensure_currency("currency", currency)?;
    let mut summaries = [
        owner_summary(Owner::User, currency),
        owner_summary(Owner::Spouse, currency),
        owner_summary(Owner::Household, currency),
    ];
    for transaction in transactions {
        transaction.validate()?;
        if transaction.amount.currency != currency {
            continue;
        }
        let Some(summary) = summaries
            .iter_mut()
            .find(|summary| summary.owner == transaction.owner)
        else {
            continue;
        };
        summary.records += 1;
        match transaction.direction {
            TransactionDirection::Debit => {
                summary.debit_minor_units += transaction.amount.minor_units
            }
            TransactionDirection::Credit => {
                summary.credit_minor_units += transaction.amount.minor_units
            }
        }
        summary.net_minor_units = summary.credit_minor_units - summary.debit_minor_units;
    }
    Ok(summaries.into_iter().collect())
}

pub fn subscription_candidates(
    transactions: &[TransactionIr],
) -> Result<Vec<SubscriptionCandidate>, ValidationError> {
    let mut candidates = Vec::new();
    for transaction in transactions {
        transaction.validate()?;
        if !is_subscription(transaction) {
            continue;
        }
        candidates.push(SubscriptionCandidate {
            transaction_id: transaction.transaction_id.clone(),
            owner: transaction.owner,
            merchant_name: transaction
                .merchant_name
                .clone()
                .unwrap_or_else(|| "unknown".to_string()),
            currency: transaction.amount.currency.clone(),
            monthly_minor_units: transaction.amount.minor_units,
            evidence_tag_count: transaction.tags.len(),
        });
    }
    candidates.sort_by(|left, right| {
        right
            .monthly_minor_units
            .cmp(&left.monthly_minor_units)
            .then_with(|| left.merchant_name.cmp(&right.merchant_name))
    });
    Ok(candidates)
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

fn owner_summary(owner: Owner, currency: &str) -> OwnerCashflowSummary {
    OwnerCashflowSummary {
        owner,
        records: 0,
        currency: currency.to_string(),
        credit_minor_units: 0,
        debit_minor_units: 0,
        net_minor_units: 0,
    }
}

fn is_subscription(transaction: &TransactionIr) -> bool {
    matches!(transaction.direction, TransactionDirection::Debit)
        && transaction
            .category
            .as_deref()
            .unwrap_or_default()
            .eq_ignore_ascii_case("subscription")
}

fn ensure_non_empty(field: &'static str, value: &str) -> Result<(), ValidationError> {
    if value.trim().is_empty() {
        return Err(ValidationError::new(field, "must not be empty"));
    }
    Ok(())
}

fn ensure_currency(field: &'static str, value: &str) -> Result<(), ValidationError> {
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

fn ensure_iso_datetime(field: &'static str, value: &str) -> Result<(), ValidationError> {
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

fn ensure_known_owner(field: &'static str, owner: Owner) -> Result<(), ValidationError> {
    if !owner.is_known() {
        return Err(ValidationError::new(field, "owner must be known"));
    }
    Ok(())
}

fn ensure_tags(field: &'static str, tags: &[String]) -> Result<(), ValidationError> {
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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn finance_fixture_ir_is_valid_jsonl() {
        let transactions = fixture_transactions().expect("finance fixture parses");
        assert_eq!(transactions.len(), 3);
        assert!(transactions
            .iter()
            .all(|transaction| transaction.validate().is_ok()));
    }

    #[test]
    fn cashflow_summary_uses_direction_not_signed_amounts() {
        let transactions = fixture_transactions().expect("finance fixture parses");
        let summary = summarize_cashflow(&transactions, "KRW").expect("cashflow summarizes");
        assert_eq!(summary.records, 3);
        assert_eq!(summary.credit_minor_units, 4_500_000);
        assert_eq!(summary.debit_minor_units, 153_200);
        assert_eq!(summary.net_minor_units, 4_346_800);
        assert_eq!(summary.subscription_minor_units, 65_900);
        assert_eq!(summary.subscription_count, 1);
    }

    #[test]
    fn owner_cashflow_keeps_household_and_personal_views_separate() {
        let transactions = fixture_transactions().expect("finance fixture parses");
        let summaries = summarize_by_owner(&transactions, "KRW").expect("owners summarize");
        let user = summaries
            .iter()
            .find(|summary| summary.owner == Owner::User)
            .expect("user summary exists");
        let spouse = summaries
            .iter()
            .find(|summary| summary.owner == Owner::Spouse)
            .expect("spouse summary exists");
        let household = summaries
            .iter()
            .find(|summary| summary.owner == Owner::Household)
            .expect("household summary exists");

        assert_eq!(user.records, 1);
        assert_eq!(user.net_minor_units, -87_300);
        assert_eq!(spouse.records, 0);
        assert_eq!(household.records, 2);
        assert_eq!(household.net_minor_units, 4_434_100);
    }

    #[test]
    fn subscription_candidates_are_review_only() {
        let transactions = fixture_transactions().expect("finance fixture parses");
        let candidates = subscription_candidates(&transactions).expect("subscriptions summarize");
        assert_eq!(candidates.len(), 1);
        assert_eq!(candidates[0].merchant_name, "Streaming Bundle");
        assert_eq!(candidates[0].monthly_minor_units, 65_900);
        assert_eq!(candidates[0].owner, Owner::Household);
    }

    #[test]
    fn debit_transactions_require_merchants() {
        let mut transaction = fixture_transactions()
            .expect("finance fixture parses")
            .into_iter()
            .find(|transaction| matches!(transaction.direction, TransactionDirection::Debit))
            .expect("fixture includes debit");
        transaction.merchant_name = None;
        let error = transaction.validate().expect_err("missing merchant fails");
        assert_eq!(error.field, "merchant_name");
    }

    #[test]
    fn unknown_owners_are_rejected() {
        let mut transaction = fixture_transactions()
            .expect("finance fixture parses")
            .into_iter()
            .next()
            .expect("fixture includes transactions");
        transaction.owner = Owner::Unknown;
        let error = transaction.validate().expect_err("unknown owner fails");
        assert_eq!(error.field, "owner");
    }
}

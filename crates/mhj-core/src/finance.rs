use crate::{
    ensure_iso_datetime, ensure_known_owner, ensure_non_empty, ensure_tags, parse_jsonl,
    FixtureError, MoneyAmount, Owner, ValidationError,
};
use serde::{Deserialize, Serialize};

pub const FINANCE_FIXTURE_JSONL: &str =
    include_str!("../../../fixtures/finance_transactions.jsonl");

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
    pub currency: String,
    pub credit_minor_units: i64,
    pub debit_minor_units: i64,
    pub net_minor_units: i64,
}

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
    let mut credit_minor_units = 0;
    let mut debit_minor_units = 0;
    for transaction in transactions {
        transaction.validate()?;
        if transaction.amount.currency != currency {
            continue;
        }
        match transaction.direction {
            TransactionDirection::Debit => debit_minor_units += transaction.amount.minor_units,
            TransactionDirection::Credit => credit_minor_units += transaction.amount.minor_units,
        }
    }
    Ok(CashflowSummary {
        currency: currency.to_string(),
        credit_minor_units,
        debit_minor_units,
        net_minor_units: credit_minor_units - debit_minor_units,
    })
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
        assert_eq!(summary.credit_minor_units, 4_500_000);
        assert_eq!(summary.debit_minor_units, 153_200);
        assert_eq!(summary.net_minor_units, 4_346_800);
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
}

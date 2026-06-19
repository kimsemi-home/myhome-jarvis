use crate::{MoneyAmount, Owner};
use serde::{Deserialize, Serialize};

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

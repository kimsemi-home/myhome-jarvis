use super::validate::validate_purchase;
use crate::{MoneyAmount, Owner, ValidationError};
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct PurchaseIr {
    pub purchase_id: String,
    pub source: String,
    pub owner: Owner,
    pub purchased_at: String,
    pub merchant_name: String,
    pub order_id: Option<String>,
    pub item_name: String,
    pub brand: Option<String>,
    pub quantity: u32,
    pub unit_price: MoneyAmount,
    pub total_price: MoneyAmount,
    pub category: Option<String>,
    pub recurring_candidate: bool,
    pub raw_ref: String,
    pub tags: Vec<String>,
}

impl PurchaseIr {
    pub fn validate(&self) -> Result<(), ValidationError> {
        validate_purchase(self)
    }
}

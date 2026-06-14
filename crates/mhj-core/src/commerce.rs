use crate::{
    ensure_iso_datetime, ensure_known_owner, ensure_non_empty, ensure_tags, parse_jsonl,
    FixtureError, MoneyAmount, Owner, ValidationError,
};
use serde::{Deserialize, Serialize};
use std::collections::BTreeMap;

pub const COMMERCE_FIXTURE_JSONL: &str = include_str!("../../../fixtures/commerce_purchases.jsonl");

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
        ensure_non_empty("purchase_id", &self.purchase_id)?;
        ensure_non_empty("source", &self.source)?;
        ensure_known_owner("owner", self.owner)?;
        ensure_iso_datetime("purchased_at", &self.purchased_at)?;
        ensure_non_empty("merchant_name", &self.merchant_name)?;
        ensure_non_empty("item_name", &self.item_name)?;
        if self.quantity == 0 {
            return Err(ValidationError::new(
                "quantity",
                "quantity must be greater than zero",
            ));
        }
        self.unit_price.validate_non_negative("unit_price")?;
        self.total_price.validate_non_negative("total_price")?;
        if self.unit_price.currency != self.total_price.currency {
            return Err(ValidationError::new(
                "total_price",
                "total price currency must match unit price currency",
            ));
        }
        let expected_total = self.unit_price.minor_units * i64::from(self.quantity);
        if expected_total != self.total_price.minor_units {
            return Err(ValidationError::new(
                "total_price",
                "total price must equal unit price times quantity",
            ));
        }
        ensure_non_empty("raw_ref", &self.raw_ref)?;
        ensure_tags("tags", &self.tags)?;
        Ok(())
    }
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct RecurringPurchaseCandidate {
    pub merchant_name: String,
    pub item_name: String,
    pub currency: String,
    pub purchase_count: usize,
    pub latest_total_minor_units: i64,
}

pub fn parse_purchases_jsonl(input: &str) -> Result<Vec<PurchaseIr>, FixtureError> {
    parse_jsonl(input, PurchaseIr::validate)
}

pub fn fixture_purchases() -> Result<Vec<PurchaseIr>, FixtureError> {
    parse_purchases_jsonl(COMMERCE_FIXTURE_JSONL)
}

pub fn recurring_candidates(purchases: &[PurchaseIr]) -> Vec<RecurringPurchaseCandidate> {
    let mut grouped: BTreeMap<(String, String, String), Vec<&PurchaseIr>> = BTreeMap::new();
    for purchase in purchases {
        if purchase.validate().is_err() || !purchase.recurring_candidate {
            continue;
        }
        grouped
            .entry((
                purchase.merchant_name.clone(),
                purchase.item_name.clone(),
                purchase.total_price.currency.clone(),
            ))
            .or_default()
            .push(purchase);
    }

    grouped
        .into_iter()
        .filter_map(|((merchant_name, item_name, currency), records)| {
            let latest = records
                .iter()
                .max_by_key(|purchase| purchase.purchased_at.as_str())?;
            if records.len() < 2 {
                return None;
            }
            Some(RecurringPurchaseCandidate {
                merchant_name,
                item_name,
                currency,
                purchase_count: records.len(),
                latest_total_minor_units: latest.total_price.minor_units,
            })
        })
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn commerce_fixture_ir_is_valid_jsonl() {
        let purchases = fixture_purchases().expect("commerce fixture parses");
        assert_eq!(purchases.len(), 3);
        assert!(purchases.iter().all(|purchase| purchase.validate().is_ok()));
    }

    #[test]
    fn recurring_candidates_require_repeated_flagged_purchases() {
        let purchases = fixture_purchases().expect("commerce fixture parses");
        let candidates = recurring_candidates(&purchases);
        assert_eq!(candidates.len(), 1);
        assert_eq!(candidates[0].merchant_name, "Coupang");
        assert_eq!(candidates[0].item_name, "Bottled water 2L x 6");
        assert_eq!(candidates[0].purchase_count, 2);
    }

    #[test]
    fn purchase_totals_must_match_quantity() {
        let mut purchase = fixture_purchases()
            .expect("commerce fixture parses")
            .into_iter()
            .next()
            .expect("fixture has purchase");
        purchase.total_price.minor_units += 1;
        let error = purchase.validate().expect_err("bad total fails");
        assert_eq!(error.field, "total_price");
    }
}

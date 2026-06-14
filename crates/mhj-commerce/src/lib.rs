use serde::de::DeserializeOwned;
use serde::{Deserialize, Serialize};
use std::collections::BTreeMap;
use std::error::Error;
use std::fmt;

pub const COMMERCE_FIXTURE_JSONL: &str = include_str!("../../../fixtures/commerce_purchases.jsonl");

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Serialize, Deserialize)]
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
    pub fn validate_non_negative(&self, field: &'static str) -> Result<(), ValidationError> {
        ensure_currency(field, &self.currency)?;
        if self.minor_units < 0 {
            return Err(ValidationError::new(field, "amount must not be negative"));
        }
        Ok(())
    }
}

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
pub struct CommerceSummary {
    pub records: usize,
    pub currency: String,
    pub total_spend_minor_units: i64,
    pub recurring_candidate_count: usize,
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct OwnerSpendSummary {
    pub owner: Owner,
    pub records: usize,
    pub currency: String,
    pub total_spend_minor_units: i64,
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct MerchantSpendSummary {
    pub merchant_name: String,
    pub records: usize,
    pub currency: String,
    pub total_spend_minor_units: i64,
}

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct RecurringPurchaseCandidate {
    pub merchant_name: String,
    pub item_name: String,
    pub currency: String,
    pub purchase_count: usize,
    pub latest_total_minor_units: i64,
    pub latest_purchased_at: String,
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

pub fn parse_purchases_jsonl(input: &str) -> Result<Vec<PurchaseIr>, FixtureError> {
    parse_jsonl(input, PurchaseIr::validate)
}

pub fn fixture_purchases() -> Result<Vec<PurchaseIr>, FixtureError> {
    parse_purchases_jsonl(COMMERCE_FIXTURE_JSONL)
}

pub fn summarize_commerce(
    purchases: &[PurchaseIr],
    currency: &str,
) -> Result<CommerceSummary, ValidationError> {
    ensure_currency("currency", currency)?;
    let mut summary = CommerceSummary {
        records: 0,
        currency: currency.to_string(),
        total_spend_minor_units: 0,
        recurring_candidate_count: 0,
    };
    for purchase in purchases {
        purchase.validate()?;
        if purchase.total_price.currency != currency {
            continue;
        }
        summary.records += 1;
        summary.total_spend_minor_units += purchase.total_price.minor_units;
        if purchase.recurring_candidate {
            summary.recurring_candidate_count += 1;
        }
    }
    Ok(summary)
}

pub fn summarize_by_owner(
    purchases: &[PurchaseIr],
    currency: &str,
) -> Result<Vec<OwnerSpendSummary>, ValidationError> {
    ensure_currency("currency", currency)?;
    let mut summaries = [
        owner_summary(Owner::User, currency),
        owner_summary(Owner::Spouse, currency),
        owner_summary(Owner::Household, currency),
    ];
    for purchase in purchases {
        purchase.validate()?;
        if purchase.total_price.currency != currency {
            continue;
        }
        let Some(summary) = summaries
            .iter_mut()
            .find(|summary| summary.owner == purchase.owner)
        else {
            continue;
        };
        summary.records += 1;
        summary.total_spend_minor_units += purchase.total_price.minor_units;
    }
    Ok(summaries.into_iter().collect())
}

pub fn summarize_by_merchant(
    purchases: &[PurchaseIr],
    currency: &str,
) -> Result<Vec<MerchantSpendSummary>, ValidationError> {
    ensure_currency("currency", currency)?;
    let mut merchants: BTreeMap<String, MerchantSpendSummary> = BTreeMap::new();
    for purchase in purchases {
        purchase.validate()?;
        if purchase.total_price.currency != currency {
            continue;
        }
        let summary = merchants
            .entry(purchase.merchant_name.clone())
            .or_insert_with(|| MerchantSpendSummary {
                merchant_name: purchase.merchant_name.clone(),
                records: 0,
                currency: currency.to_string(),
                total_spend_minor_units: 0,
            });
        summary.records += 1;
        summary.total_spend_minor_units += purchase.total_price.minor_units;
    }
    let mut summaries: Vec<MerchantSpendSummary> = merchants.into_values().collect();
    summaries.sort_by(|left, right| {
        right
            .total_spend_minor_units
            .cmp(&left.total_spend_minor_units)
            .then_with(|| left.merchant_name.cmp(&right.merchant_name))
    });
    Ok(summaries)
}

pub fn recurring_candidates(
    purchases: &[PurchaseIr],
) -> Result<Vec<RecurringPurchaseCandidate>, ValidationError> {
    let mut grouped: BTreeMap<(String, String, String), Vec<&PurchaseIr>> = BTreeMap::new();
    for purchase in purchases {
        purchase.validate()?;
        if !purchase.recurring_candidate {
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

    let mut candidates: Vec<RecurringPurchaseCandidate> = grouped
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
                latest_purchased_at: latest.purchased_at.clone(),
            })
        })
        .collect();
    candidates.sort_by(|left, right| {
        right
            .purchase_count
            .cmp(&left.purchase_count)
            .then_with(|| left.merchant_name.cmp(&right.merchant_name))
            .then_with(|| left.item_name.cmp(&right.item_name))
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

fn owner_summary(owner: Owner, currency: &str) -> OwnerSpendSummary {
    OwnerSpendSummary {
        owner,
        records: 0,
        currency: currency.to_string(),
        total_spend_minor_units: 0,
    }
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
    fn commerce_fixture_ir_is_valid_jsonl() {
        let purchases = fixture_purchases().expect("commerce fixture parses");
        assert_eq!(purchases.len(), 3);
        assert!(purchases.iter().all(|purchase| purchase.validate().is_ok()));
    }

    #[test]
    fn commerce_summary_counts_fixture_spend_and_recurring_flags() {
        let purchases = fixture_purchases().expect("commerce fixture parses");
        let summary = summarize_commerce(&purchases, "KRW").expect("commerce summarizes");

        assert_eq!(summary.records, 3);
        assert_eq!(summary.total_spend_minor_units, 26_800);
        assert_eq!(summary.recurring_candidate_count, 2);
    }

    #[test]
    fn owner_spend_keeps_household_and_personal_purchases_separate() {
        let purchases = fixture_purchases().expect("commerce fixture parses");
        let summaries = summarize_by_owner(&purchases, "KRW").expect("owners summarize");
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
        assert_eq!(user.total_spend_minor_units, 3_200);
        assert_eq!(spouse.records, 0);
        assert_eq!(household.records, 2);
        assert_eq!(household.total_spend_minor_units, 23_600);
    }

    #[test]
    fn merchant_spend_is_sorted_by_total_spend() {
        let purchases = fixture_purchases().expect("commerce fixture parses");
        let merchants = summarize_by_merchant(&purchases, "KRW").expect("merchants summarize");

        assert_eq!(merchants.len(), 2);
        assert_eq!(merchants[0].merchant_name, "Coupang");
        assert_eq!(merchants[0].total_spend_minor_units, 23_600);
        assert_eq!(merchants[1].merchant_name, "Local Mart");
    }

    #[test]
    fn recurring_candidates_require_repeated_flagged_purchases() {
        let purchases = fixture_purchases().expect("commerce fixture parses");
        let candidates = recurring_candidates(&purchases).expect("recurring summarizes");

        assert_eq!(candidates.len(), 1);
        assert_eq!(candidates[0].merchant_name, "Coupang");
        assert_eq!(candidates[0].item_name, "Bottled water 2L x 6");
        assert_eq!(candidates[0].purchase_count, 2);
        assert_eq!(candidates[0].latest_total_minor_units, 11_800);
        assert_eq!(
            candidates[0].latest_purchased_at,
            "2026-06-11T11:15:00+09:00"
        );
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

    #[test]
    fn unknown_owners_are_rejected() {
        let mut purchase = fixture_purchases()
            .expect("commerce fixture parses")
            .into_iter()
            .next()
            .expect("fixture has purchase");
        purchase.owner = Owner::Unknown;
        let error = purchase.validate().expect_err("unknown owner fails");
        assert_eq!(error.field, "owner");
    }
}

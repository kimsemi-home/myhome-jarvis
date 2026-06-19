use crate::Owner;

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

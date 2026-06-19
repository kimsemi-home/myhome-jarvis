use crate::Owner;

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
pub struct CardUsageCandidate {
    pub currency: String,
    pub debit_minor_units: i64,
    pub transaction_count: usize,
    pub subscription_minor_units: i64,
    pub subscription_count: usize,
}

pub(crate) fn owner_summary(owner: Owner, currency: &str) -> OwnerCashflowSummary {
    OwnerCashflowSummary {
        owner,
        records: 0,
        currency: currency.to_string(),
        credit_minor_units: 0,
        debit_minor_units: 0,
        net_minor_units: 0,
    }
}

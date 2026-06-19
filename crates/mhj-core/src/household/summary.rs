use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct HouseholdScopeSummary {
    pub scope: String,
    pub label: String,
    pub currency: String,
    pub finance_records: usize,
    pub finance_credit_minor_units: i64,
    pub finance_debit_minor_units: i64,
    pub finance_net_minor_units: i64,
    pub purchase_records: usize,
    pub purchase_spend_minor_units: i64,
}

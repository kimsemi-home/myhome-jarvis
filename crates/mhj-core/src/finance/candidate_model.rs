#[derive(Debug, Clone, PartialEq, Eq)]
pub struct CardUsageCandidate {
    pub currency: String,
    pub debit_minor_units: i64,
    pub transaction_count: usize,
    pub subscription_minor_units: i64,
    pub subscription_count: usize,
}

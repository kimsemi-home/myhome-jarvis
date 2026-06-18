use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum RecommendationKind {
    CardUsageReview,
    CashBuffer,
    RecurringPurchaseReview,
    SubscriptionReview,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct Recommendation {
    pub kind: RecommendationKind,
    pub title: String,
    pub rationale: String,
    pub score: u8,
    pub currency: String,
    pub estimated_monthly_minor_units: i64,
    pub evidence_count: usize,
}

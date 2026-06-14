use crate::commerce::{recurring_candidates, PurchaseIr};
use crate::finance::{summarize_cashflow, TransactionDirection, TransactionIr};
use crate::{FixtureError, ValidationError};
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum RecommendationKind {
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

pub fn fixture_recommendations() -> Result<Vec<Recommendation>, FixtureError> {
    let transactions = crate::finance::fixture_transactions()?;
    let purchases = crate::commerce::fixture_purchases()?;
    score_recommendations(&transactions, &purchases).map_err(|error| FixtureError::Validation {
        line: 0,
        field: error.field,
        message: error.message,
    })
}

pub fn score_recommendations(
    transactions: &[TransactionIr],
    purchases: &[PurchaseIr],
) -> Result<Vec<Recommendation>, ValidationError> {
    for transaction in transactions {
        transaction.validate()?;
    }
    for purchase in purchases {
        purchase.validate()?;
    }

    let mut recommendations = Vec::new();
    let cashflow = summarize_cashflow(transactions, "KRW")?;
    if cashflow.net_minor_units > 0 {
        recommendations.push(Recommendation {
            kind: RecommendationKind::CashBuffer,
            title: "Keep household cash buffer".to_string(),
            rationale: "Fixture cashflow is positive; reserve surplus before recommendations become executable.".to_string(),
            score: clamp_score(45 + cashflow.net_minor_units / 1_000_000),
            currency: cashflow.currency,
            estimated_monthly_minor_units: cashflow.net_minor_units,
            evidence_count: transactions.len(),
        });
    }

    let subscription_total: i64 = transactions
        .iter()
        .filter(|transaction| {
            matches!(transaction.direction, TransactionDirection::Debit)
                && transaction.category.as_deref() == Some("subscription")
        })
        .map(|transaction| transaction.amount.minor_units)
        .sum();
    let subscription_count = transactions
        .iter()
        .filter(|transaction| {
            matches!(transaction.direction, TransactionDirection::Debit)
                && transaction.category.as_deref() == Some("subscription")
        })
        .count();
    if subscription_total > 0 {
        recommendations.push(Recommendation {
            kind: RecommendationKind::SubscriptionReview,
            title: "Review household subscriptions".to_string(),
            rationale:
                "Subscription-like debit fixtures exist; keep this as a review-only recommendation."
                    .to_string(),
            score: clamp_score(55 + subscription_total / 10_000),
            currency: "KRW".to_string(),
            estimated_monthly_minor_units: subscription_total,
            evidence_count: subscription_count,
        });
    }

    for candidate in recurring_candidates(purchases) {
        recommendations.push(Recommendation {
            kind: RecommendationKind::RecurringPurchaseReview,
            title: format!("Compare recurring purchase: {}", candidate.item_name),
            rationale: format!(
                "{} appears repeatedly in local purchase fixtures.",
                candidate.merchant_name
            ),
            score: clamp_score(
                50 + i64::try_from(candidate.purchase_count).unwrap_or(0) * 10
                    + candidate.latest_total_minor_units / 1_000,
            ),
            currency: candidate.currency,
            estimated_monthly_minor_units: candidate.latest_total_minor_units,
            evidence_count: candidate.purchase_count,
        });
    }

    recommendations.sort_by(|left, right| {
        right
            .score
            .cmp(&left.score)
            .then_with(|| left.title.cmp(&right.title))
    });
    Ok(recommendations)
}

fn clamp_score(value: i64) -> u8 {
    value.clamp(0, 100) as u8
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn fixture_recommendations_are_ranked_and_bounded() {
        let recommendations = fixture_recommendations().expect("recommendations score");
        assert_eq!(recommendations.len(), 3);
        assert!(recommendations
            .windows(2)
            .all(|window| window[0].score >= window[1].score));
        assert!(recommendations
            .iter()
            .all(|recommendation| recommendation.score <= 100));
        assert!(recommendations
            .iter()
            .any(|recommendation| recommendation.kind == RecommendationKind::SubscriptionReview));
        assert!(recommendations.iter().any(|recommendation| {
            recommendation.kind == RecommendationKind::RecurringPurchaseReview
                && recommendation.evidence_count == 2
        }));
    }
}

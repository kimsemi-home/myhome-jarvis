use crate::finance::{TransactionDirection, TransactionIr};

use super::score::clamp_score;
use super::{Recommendation, RecommendationKind};

pub(crate) fn push_subscription_review(
    recommendations: &mut Vec<Recommendation>,
    transactions: &[TransactionIr],
) {
    let (total, count) = subscription_debits(transactions);
    if total <= 0 {
        return;
    }
    recommendations.push(Recommendation {
        kind: RecommendationKind::SubscriptionReview,
        title: "Review household subscriptions".to_string(),
        rationale:
            "Subscription-like debit fixtures exist; keep this as a review-only recommendation."
                .to_string(),
        score: clamp_score(55 + total / 10_000),
        currency: "KRW".to_string(),
        estimated_monthly_minor_units: total,
        evidence_count: count,
    });
}

fn subscription_debits(transactions: &[TransactionIr]) -> (i64, usize) {
    transactions
        .iter()
        .filter(|transaction| {
            matches!(transaction.direction, TransactionDirection::Debit)
                && transaction.category.as_deref() == Some("subscription")
        })
        .fold((0, 0), |(total, count), transaction| {
            (total + transaction.amount.minor_units, count + 1)
        })
}

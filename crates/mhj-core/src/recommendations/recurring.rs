use crate::commerce::{recurring_candidates, PurchaseIr};

use super::score::clamp_score;
use super::{Recommendation, RecommendationKind};

pub(crate) fn push_recurring_purchase_reviews(
    recommendations: &mut Vec<Recommendation>,
    purchases: &[PurchaseIr],
) {
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
}

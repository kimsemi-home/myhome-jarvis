use crate::finance::{card_usage_candidates, TransactionIr};
use crate::ValidationError;

use super::score::clamp_score;
use super::{Recommendation, RecommendationKind};

pub(crate) fn push_card_usage_reviews(
    recommendations: &mut Vec<Recommendation>,
    transactions: &[TransactionIr],
) -> Result<(), ValidationError> {
    for candidate in card_usage_candidates(transactions, "KRW")? {
        recommendations.push(Recommendation {
            kind: RecommendationKind::CardUsageReview,
            title: "Review card-linked household spend".to_string(),
            rationale: "Card-linked debit fixtures exist; keep this as a review-only recommendation, not a card action.".to_string(),
            score: clamp_score(52 + candidate.debit_minor_units / 10_000),
            currency: candidate.currency,
            estimated_monthly_minor_units: candidate.debit_minor_units,
            evidence_count: candidate.transaction_count,
        });
    }
    Ok(())
}

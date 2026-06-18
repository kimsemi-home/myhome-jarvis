use crate::commerce::PurchaseIr;
use crate::finance::TransactionIr;
use crate::ValidationError;

use super::card::push_card_usage_reviews;
use super::cashflow::push_cash_buffer;
use super::recurring::push_recurring_purchase_reviews;
use super::subscription::push_subscription_review;
use super::Recommendation;

pub fn score_recommendations(
    transactions: &[TransactionIr],
    purchases: &[PurchaseIr],
) -> Result<Vec<Recommendation>, ValidationError> {
    validate_inputs(transactions, purchases)?;

    let mut recommendations = Vec::new();
    push_cash_buffer(&mut recommendations, transactions)?;
    push_subscription_review(&mut recommendations, transactions);
    push_card_usage_reviews(&mut recommendations, transactions)?;
    push_recurring_purchase_reviews(&mut recommendations, purchases);
    sort_recommendations(&mut recommendations);
    Ok(recommendations)
}

pub(crate) fn clamp_score(value: i64) -> u8 {
    value.clamp(0, 100) as u8
}

fn validate_inputs(
    transactions: &[TransactionIr],
    purchases: &[PurchaseIr],
) -> Result<(), ValidationError> {
    for transaction in transactions {
        transaction.validate()?;
    }
    for purchase in purchases {
        purchase.validate()?;
    }
    Ok(())
}

fn sort_recommendations(recommendations: &mut [Recommendation]) {
    recommendations.sort_by(|left, right| {
        right
            .score
            .cmp(&left.score)
            .then_with(|| left.title.cmp(&right.title))
    });
}

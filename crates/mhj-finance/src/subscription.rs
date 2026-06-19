use crate::classify::is_subscription;
use crate::{SubscriptionCandidate, TransactionIr, ValidationError};

pub fn subscription_candidates(
    transactions: &[TransactionIr],
) -> Result<Vec<SubscriptionCandidate>, ValidationError> {
    let mut candidates = Vec::new();
    for transaction in transactions {
        transaction.validate()?;
        if !is_subscription(transaction) {
            continue;
        }
        candidates.push(subscription_candidate(transaction));
    }
    candidates.sort_by(|left, right| {
        right
            .monthly_minor_units
            .cmp(&left.monthly_minor_units)
            .then_with(|| left.merchant_name.cmp(&right.merchant_name))
    });
    Ok(candidates)
}

fn subscription_candidate(transaction: &TransactionIr) -> SubscriptionCandidate {
    SubscriptionCandidate {
        transaction_id: transaction.transaction_id.clone(),
        owner: transaction.owner,
        merchant_name: transaction
            .merchant_name
            .clone()
            .unwrap_or_else(|| "unknown".to_string()),
        currency: transaction.amount.currency.clone(),
        monthly_minor_units: transaction.amount.minor_units,
        evidence_tag_count: transaction.tags.len(),
    }
}

use crate::ValidationError;
use std::collections::BTreeMap;

use super::{CardUsageCandidate, TransactionDirection, TransactionIr};

pub fn card_usage_candidates(
    transactions: &[TransactionIr],
    currency: &str,
) -> Result<Vec<CardUsageCandidate>, ValidationError> {
    let mut cards = BTreeMap::<String, CardUsageCandidate>::new();
    for transaction in transactions {
        transaction.validate()?;
        if !is_card_debit(transaction, currency) {
            continue;
        }
        apply_card_transaction(&mut cards, currency, transaction);
    }
    let mut candidates = cards.into_values().collect::<Vec<_>>();
    candidates.sort_by(|left, right| {
        right
            .debit_minor_units
            .cmp(&left.debit_minor_units)
            .then_with(|| right.transaction_count.cmp(&left.transaction_count))
    });
    Ok(candidates)
}

fn is_card_debit(transaction: &TransactionIr, currency: &str) -> bool {
    transaction.amount.currency == currency
        && matches!(transaction.direction, TransactionDirection::Debit)
        && !transaction
            .card_id
            .as_deref()
            .unwrap_or_default()
            .trim()
            .is_empty()
}

fn apply_card_transaction(
    cards: &mut BTreeMap<String, CardUsageCandidate>,
    currency: &str,
    transaction: &TransactionIr,
) {
    let candidate = cards
        .entry(
            transaction
                .card_id
                .as_deref()
                .unwrap_or_default()
                .to_string(),
        )
        .or_insert_with(|| CardUsageCandidate {
            currency: currency.to_string(),
            debit_minor_units: 0,
            transaction_count: 0,
            subscription_minor_units: 0,
            subscription_count: 0,
        });
    candidate.debit_minor_units += transaction.amount.minor_units;
    candidate.transaction_count += 1;
    if is_subscription(transaction) {
        candidate.subscription_minor_units += transaction.amount.minor_units;
        candidate.subscription_count += 1;
    }
}

fn is_subscription(transaction: &TransactionIr) -> bool {
    transaction
        .category
        .as_deref()
        .unwrap_or_default()
        .eq_ignore_ascii_case("subscription")
}

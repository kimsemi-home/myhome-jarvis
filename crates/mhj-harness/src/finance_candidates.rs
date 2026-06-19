use crate::check::check;
use crate::finance_snapshot::FinanceSnapshot;
use crate::HarnessResult;

pub(crate) fn finance_candidate_checks(snapshot: &FinanceSnapshot) -> Vec<HarnessResult> {
    let summary = &snapshot.summary;
    vec![
        check(
            "subscription review candidates",
            snapshot.subscriptions.len() == 1
                && summary.subscription_count == 1
                && summary.subscription_minor_units == 65_900
                && snapshot.subscriptions[0].merchant_name == "Streaming Bundle",
            format!(
                "subscriptions={} total={}",
                summary.subscription_count, summary.subscription_minor_units
            ),
        ),
        check(
            "card-linked debit review candidates",
            card_candidate_ok(snapshot),
            card_candidate_message(snapshot),
        ),
    ]
}

fn card_candidate_ok(snapshot: &FinanceSnapshot) -> bool {
    snapshot.cards.len() == 1
        && snapshot.cards[0].transaction_count == 2
        && snapshot.cards[0].debit_minor_units == 153_200
        && snapshot.cards[0].subscription_count == 1
}

fn card_candidate_message(snapshot: &FinanceSnapshot) -> String {
    snapshot
        .cards
        .first()
        .map(|card| {
            format!(
                "cards={} transactions={} debit={}",
                snapshot.cards.len(),
                card.transaction_count,
                card.debit_minor_units
            )
        })
        .unwrap_or_else(|| "cards=0".to_string())
}

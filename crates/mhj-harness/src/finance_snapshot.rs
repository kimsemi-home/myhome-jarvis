use crate::check::check;
use crate::HarnessResult;
use mhj_finance as finance;

pub(crate) struct FinanceSnapshot {
    pub transaction_count: usize,
    pub summary: finance::CashflowSummary,
    pub owners: Vec<finance::OwnerCashflowSummary>,
    pub subscriptions: Vec<finance::SubscriptionCandidate>,
    pub cards: Vec<finance::CardUsageCandidate>,
}

pub(crate) fn finance_snapshot() -> Result<FinanceSnapshot, HarnessResult> {
    let transactions = finance::fixture_transactions().map_err(|error| {
        check(
            "finance fixture parses",
            false,
            format!("fixture parse failed: {error}"),
        )
    })?;
    let summary = finance::summarize_cashflow(&transactions, "KRW").map_err(|error| {
        check(
            "finance cashflow summarizes",
            false,
            format!("cashflow failed: {error}"),
        )
    })?;
    let owners = finance::summarize_by_owner(&transactions, "KRW").map_err(|error| {
        check(
            "finance owners summarize",
            false,
            format!("owner summary failed: {error}"),
        )
    })?;
    let subscriptions = finance::subscription_candidates(&transactions).map_err(|error| {
        check(
            "subscription candidates summarize",
            false,
            format!("subscription summary failed: {error}"),
        )
    })?;
    let cards = finance::card_usage_candidates(&transactions, "KRW").map_err(|error| {
        check(
            "card usage candidates summarize",
            false,
            format!("card summary failed: {error}"),
        )
    })?;
    Ok(FinanceSnapshot {
        transaction_count: transactions.len(),
        summary,
        owners,
        subscriptions,
        cards,
    })
}

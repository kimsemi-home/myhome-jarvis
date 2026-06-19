use crate::check::check;
use crate::HarnessResult;
use mhj_commerce as commerce;

pub(crate) struct CommerceSnapshot {
    pub purchase_count: usize,
    pub summary: commerce::CommerceSummary,
    pub owners: Vec<commerce::OwnerSpendSummary>,
    pub merchants: Vec<commerce::MerchantSpendSummary>,
    pub recurring: Vec<commerce::RecurringPurchaseCandidate>,
}

pub(crate) fn commerce_snapshot() -> Result<CommerceSnapshot, HarnessResult> {
    let purchases = commerce::fixture_purchases().map_err(|error| {
        check(
            "commerce fixture parses",
            false,
            format!("fixture parse failed: {error}"),
        )
    })?;
    let summary = commerce::summarize_commerce(&purchases, "KRW")
        .map_err(|error| check("commerce summary builds", false, format!("{error}")))?;
    let owners = commerce::summarize_by_owner(&purchases, "KRW").map_err(|error| {
        check(
            "commerce owners summarize",
            false,
            format!("owner summary failed: {error}"),
        )
    })?;
    let merchants = commerce::summarize_by_merchant(&purchases, "KRW").map_err(|error| {
        check(
            "commerce merchants summarize",
            false,
            format!("merchant summary failed: {error}"),
        )
    })?;
    let recurring = commerce::recurring_candidates(&purchases).map_err(|error| {
        check(
            "recurring purchase candidates summarize",
            false,
            format!("recurring summary failed: {error}"),
        )
    })?;
    Ok(CommerceSnapshot {
        purchase_count: purchases.len(),
        summary,
        owners,
        merchants,
        recurring,
    })
}

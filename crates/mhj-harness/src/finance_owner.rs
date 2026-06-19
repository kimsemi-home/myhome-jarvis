use crate::check::check;
use crate::finance_snapshot::FinanceSnapshot;
use crate::HarnessResult;
use mhj_finance as finance;

pub(crate) fn finance_owner_checks(snapshot: &FinanceSnapshot) -> Vec<HarnessResult> {
    let user = finance_owner(&snapshot.owners, finance::Owner::User);
    let household = finance_owner(&snapshot.owners, finance::Owner::Household);
    vec![
        check(
            "user finance scope",
            user.is_some_and(|owner| owner.records == 1 && owner.net_minor_units == -87_300),
            finance_owner_message(user),
        ),
        check(
            "household finance scope",
            household.is_some_and(|owner| owner.records == 2 && owner.net_minor_units == 4_434_100),
            finance_owner_message(household),
        ),
    ]
}

fn finance_owner(
    owners: &[finance::OwnerCashflowSummary],
    owner: finance::Owner,
) -> Option<&finance::OwnerCashflowSummary> {
    owners.iter().find(|summary| summary.owner == owner)
}

fn finance_owner_message(owner: Option<&finance::OwnerCashflowSummary>) -> String {
    owner
        .map(|owner| {
            format!(
                "{} records={} net={}",
                owner.owner.as_str(),
                owner.records,
                owner.net_minor_units
            )
        })
        .unwrap_or_else(|| "missing owner".to_string())
}

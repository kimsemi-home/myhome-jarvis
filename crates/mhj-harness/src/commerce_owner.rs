use crate::check::check;
use crate::commerce_snapshot::CommerceSnapshot;
use crate::HarnessResult;
use mhj_commerce as commerce;

pub(crate) fn commerce_owner_checks(snapshot: &CommerceSnapshot) -> Vec<HarnessResult> {
    let user = commerce_owner(&snapshot.owners, commerce::Owner::User);
    let household = commerce_owner(&snapshot.owners, commerce::Owner::Household);
    vec![
        check(
            "user commerce scope",
            user.is_some_and(|owner| owner.records == 1 && owner.total_spend_minor_units == 3_200),
            commerce_owner_message(user),
        ),
        check(
            "household commerce scope",
            household
                .is_some_and(|owner| owner.records == 2 && owner.total_spend_minor_units == 23_600),
            commerce_owner_message(household),
        ),
    ]
}

fn commerce_owner(
    owners: &[commerce::OwnerSpendSummary],
    owner: commerce::Owner,
) -> Option<&commerce::OwnerSpendSummary> {
    owners.iter().find(|summary| summary.owner == owner)
}

fn commerce_owner_message(owner: Option<&commerce::OwnerSpendSummary>) -> String {
    owner
        .map(|owner| {
            format!(
                "{} records={} spend={}",
                owner.owner.as_str(),
                owner.records,
                owner.total_spend_minor_units
            )
        })
        .unwrap_or_else(|| "missing owner".to_string())
}

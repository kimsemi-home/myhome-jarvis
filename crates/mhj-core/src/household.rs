use crate::commerce::PurchaseIr;
use crate::finance::{TransactionDirection, TransactionIr};
use crate::{FixtureError, Owner, ValidationError};
use serde::{Deserialize, Serialize};
use std::collections::BTreeSet;

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum HouseholdScope {
    User,
    Spouse,
    Household,
}

impl HouseholdScope {
    fn label(self) -> &'static str {
        match self {
            HouseholdScope::User => "User",
            HouseholdScope::Spouse => "Spouse",
            HouseholdScope::Household => "Household",
        }
    }

    fn key(self) -> &'static str {
        match self {
            HouseholdScope::User => "user",
            HouseholdScope::Spouse => "spouse",
            HouseholdScope::Household => "household",
        }
    }

    fn includes(self, owner: Owner) -> bool {
        match self {
            HouseholdScope::User => owner == Owner::User,
            HouseholdScope::Spouse => owner == Owner::Spouse,
            HouseholdScope::Household => owner.is_known(),
        }
    }
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct HouseholdScopeSummary {
    pub scope: String,
    pub label: String,
    pub currency: String,
    pub finance_records: usize,
    pub finance_credit_minor_units: i64,
    pub finance_debit_minor_units: i64,
    pub finance_net_minor_units: i64,
    pub purchase_records: usize,
    pub purchase_spend_minor_units: i64,
}

pub fn fixture_household_scopes() -> Result<Vec<HouseholdScopeSummary>, FixtureError> {
    let transactions = crate::finance::fixture_transactions()?;
    let purchases = crate::commerce::fixture_purchases()?;
    build_household_scopes(&transactions, &purchases).map_err(|error| FixtureError::Validation {
        line: 0,
        field: error.field,
        message: error.message,
    })
}

pub fn build_household_scopes(
    transactions: &[TransactionIr],
    purchases: &[PurchaseIr],
) -> Result<Vec<HouseholdScopeSummary>, ValidationError> {
    let scopes = [
        HouseholdScope::User,
        HouseholdScope::Spouse,
        HouseholdScope::Household,
    ];
    scopes
        .into_iter()
        .map(|scope| build_scope(scope, transactions, purchases))
        .collect()
}

fn build_scope(
    scope: HouseholdScope,
    transactions: &[TransactionIr],
    purchases: &[PurchaseIr],
) -> Result<HouseholdScopeSummary, ValidationError> {
    let mut currencies = BTreeSet::new();
    let mut finance_records = 0;
    let mut finance_credit_minor_units = 0;
    let mut finance_debit_minor_units = 0;
    let mut purchase_records = 0;
    let mut purchase_spend_minor_units = 0;

    for transaction in transactions {
        transaction.validate()?;
        if !scope.includes(transaction.owner) {
            continue;
        }
        finance_records += 1;
        currencies.insert(transaction.amount.currency.as_str());
        match transaction.direction {
            TransactionDirection::Credit => {
                finance_credit_minor_units += transaction.amount.minor_units
            }
            TransactionDirection::Debit => {
                finance_debit_minor_units += transaction.amount.minor_units
            }
        }
    }

    for purchase in purchases {
        purchase.validate()?;
        if !scope.includes(purchase.owner) {
            continue;
        }
        purchase_records += 1;
        currencies.insert(purchase.total_price.currency.as_str());
        purchase_spend_minor_units += purchase.total_price.minor_units;
    }

    Ok(HouseholdScopeSummary {
        scope: scope.key().to_string(),
        label: scope.label().to_string(),
        currency: scope_currency(&currencies),
        finance_records,
        finance_credit_minor_units,
        finance_debit_minor_units,
        finance_net_minor_units: finance_credit_minor_units - finance_debit_minor_units,
        purchase_records,
        purchase_spend_minor_units,
    })
}

fn scope_currency(currencies: &BTreeSet<&str>) -> String {
    match currencies.len() {
        0 => "KRW".to_string(),
        1 => currencies
            .iter()
            .next()
            .copied()
            .unwrap_or("KRW")
            .to_string(),
        _ => "mixed".to_string(),
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn fixture_household_scopes_include_user_spouse_and_household() {
        let scopes = fixture_household_scopes().expect("household scopes build");
        assert_eq!(scopes.len(), 3);

        let user = scopes
            .iter()
            .find(|scope| scope.scope == "user")
            .expect("user scope exists");
        assert_eq!(user.finance_records, 1);
        assert_eq!(user.finance_net_minor_units, -87_300);
        assert_eq!(user.purchase_records, 1);
        assert_eq!(user.purchase_spend_minor_units, 3_200);

        let spouse = scopes
            .iter()
            .find(|scope| scope.scope == "spouse")
            .expect("spouse scope exists");
        assert_eq!(spouse.finance_records, 0);
        assert_eq!(spouse.purchase_records, 0);

        let household = scopes
            .iter()
            .find(|scope| scope.scope == "household")
            .expect("household scope exists");
        assert_eq!(household.finance_records, 3);
        assert_eq!(household.finance_net_minor_units, 4_346_800);
        assert_eq!(household.purchase_records, 3);
        assert_eq!(household.purchase_spend_minor_units, 26_800);
    }
}

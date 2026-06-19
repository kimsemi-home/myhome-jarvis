use crate::commerce::PurchaseIr;
use crate::finance::{TransactionDirection, TransactionIr};
use crate::ValidationError;
use std::collections::BTreeSet;

use super::currency::scope_currency;
use super::{HouseholdScope, HouseholdScopeSummary};

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

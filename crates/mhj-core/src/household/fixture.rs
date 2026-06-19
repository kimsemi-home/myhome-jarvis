use crate::{FixtureError, ValidationError};

use super::{build_household_scopes, HouseholdScopeSummary};

pub fn fixture_household_scopes() -> Result<Vec<HouseholdScopeSummary>, FixtureError> {
    let transactions = crate::finance::fixture_transactions()?;
    let purchases = crate::commerce::fixture_purchases()?;
    build_household_scopes(&transactions, &purchases).map_err(fixture_error)
}

fn fixture_error(error: ValidationError) -> FixtureError {
    FixtureError::Validation {
        line: 0,
        field: error.field,
        message: error.message,
    }
}

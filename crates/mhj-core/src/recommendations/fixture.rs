use crate::FixtureError;

use super::{score_recommendations, Recommendation};

pub fn fixture_recommendations() -> Result<Vec<Recommendation>, FixtureError> {
    let transactions = crate::finance::fixture_transactions()?;
    let purchases = crate::commerce::fixture_purchases()?;
    score_recommendations(&transactions, &purchases).map_err(|error| FixtureError::Validation {
        line: 0,
        field: error.field,
        message: error.message,
    })
}

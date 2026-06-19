use crate::Owner;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum HouseholdScope {
    User,
    Spouse,
    Household,
}

impl HouseholdScope {
    pub(super) fn label(self) -> &'static str {
        match self {
            HouseholdScope::User => "User",
            HouseholdScope::Spouse => "Spouse",
            HouseholdScope::Household => "Household",
        }
    }

    pub(super) fn key(self) -> &'static str {
        match self {
            HouseholdScope::User => "user",
            HouseholdScope::Spouse => "spouse",
            HouseholdScope::Household => "household",
        }
    }

    pub(super) fn includes(self, owner: Owner) -> bool {
        match self {
            HouseholdScope::User => owner == Owner::User,
            HouseholdScope::Spouse => owner == Owner::Spouse,
            HouseholdScope::Household => owner.is_known(),
        }
    }
}

use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum Owner {
    User,
    Spouse,
    Household,
    Unknown,
}

impl Owner {
    pub fn is_known(self) -> bool {
        !matches!(self, Owner::Unknown)
    }

    pub fn as_str(self) -> &'static str {
        match self {
            Owner::User => "user",
            Owner::Spouse => "spouse",
            Owner::Household => "household",
            Owner::Unknown => "unknown",
        }
    }
}

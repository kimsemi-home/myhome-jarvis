use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
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
}

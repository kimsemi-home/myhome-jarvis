use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum LakeLayer {
    Raw,
    Bronze,
    Silver,
    Gold,
}

impl LakeLayer {
    pub fn as_str(self) -> &'static str {
        match self {
            LakeLayer::Raw => "raw",
            LakeLayer::Bronze => "bronze",
            LakeLayer::Silver => "silver",
            LakeLayer::Gold => "gold",
        }
    }

    pub fn all() -> [LakeLayer; 4] {
        [
            LakeLayer::Raw,
            LakeLayer::Bronze,
            LakeLayer::Silver,
            LakeLayer::Gold,
        ]
    }
}

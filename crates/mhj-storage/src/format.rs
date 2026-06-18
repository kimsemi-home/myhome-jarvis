use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum StorageFormat {
    Jsonl,
    Parquet,
}

impl StorageFormat {
    pub fn extension(self) -> &'static str {
        match self {
            StorageFormat::Jsonl => "jsonl",
            StorageFormat::Parquet => "parquet",
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum Compression {
    None,
    Zstd,
}

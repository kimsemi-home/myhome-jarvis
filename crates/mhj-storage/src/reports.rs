use crate::{Compression, DatasetKind, LakeLayer, StorageFormat};
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct DatasetFile {
    pub layer: LakeLayer,
    pub dataset: DatasetKind,
    pub format: StorageFormat,
    pub compression: Compression,
    pub schema_version: String,
    pub relative_path: String,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct LakeManifest {
    pub lake_root: String,
    pub schema_version: String,
    pub files: Vec<DatasetFile>,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct CuratedWriteReport {
    pub relative_path: String,
    pub layer: LakeLayer,
    pub dataset: DatasetKind,
    pub format: StorageFormat,
    pub compression: Compression,
    pub schema_version: String,
    pub row_count: usize,
}

#[derive(Debug, Clone, PartialEq, Eq, Serialize, Deserialize)]
pub struct CuratedReadReport {
    pub relative_path: String,
    pub layer: LakeLayer,
    pub dataset: DatasetKind,
    pub format: StorageFormat,
    pub compression: Compression,
    pub schema_version: String,
    pub row_count: usize,
    pub row_group_count: usize,
    pub column_count: usize,
}

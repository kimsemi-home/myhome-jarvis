use crate::ValidationError;
use serde::{Deserialize, Serialize};
use std::path::{Path, PathBuf};

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
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub enum DatasetKind {
    FinanceTransactions,
    CommercePurchases,
}

impl DatasetKind {
    pub fn slug(self) -> &'static str {
        match self {
            DatasetKind::FinanceTransactions => "finance_transactions",
            DatasetKind::CommercePurchases => "commerce_purchases",
        }
    }
}

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

#[derive(Debug, Clone, PartialEq, Eq)]
pub struct DatasetPlan {
    pub layer: LakeLayer,
    pub dataset: DatasetKind,
    pub format: StorageFormat,
    pub compression: Compression,
    pub path: PathBuf,
}

pub fn plan_dataset(root: &Path, layer: LakeLayer, dataset: DatasetKind) -> DatasetPlan {
    let format = match layer {
        LakeLayer::Raw => StorageFormat::Jsonl,
        LakeLayer::Bronze | LakeLayer::Silver | LakeLayer::Gold => StorageFormat::Parquet,
    };
    let compression = match format {
        StorageFormat::Jsonl => Compression::None,
        StorageFormat::Parquet => Compression::Zstd,
    };
    let filename = format!("{}.{}", dataset.slug(), format.extension());
    DatasetPlan {
        layer,
        dataset,
        format,
        compression,
        path: root
            .join(layer.as_str())
            .join(dataset.slug())
            .join(filename),
    }
}

pub fn partitioned_dataset_path(
    root: &Path,
    layer: LakeLayer,
    dataset: DatasetKind,
    partition: &str,
) -> Result<PathBuf, ValidationError> {
    validate_partition(partition)?;
    let base = plan_dataset(root, layer, dataset);
    let filename = base
        .path
        .file_name()
        .expect("dataset plan always has file name");
    Ok(base.path.with_file_name(partition).join(filename))
}

pub fn validate_partition(partition: &str) -> Result<(), ValidationError> {
    if partition.trim().is_empty() {
        return Err(ValidationError::new("partition", "must not be empty"));
    }
    if partition.starts_with('/') || partition.contains('\\') || partition.contains("..") {
        return Err(ValidationError::new(
            "partition",
            "must be a relative path segment without traversal",
        ));
    }
    if partition
        .split('/')
        .any(|segment| segment.trim().is_empty())
    {
        return Err(ValidationError::new(
            "partition",
            "must not contain empty path segments",
        ));
    }
    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn raw_layer_uses_jsonl_without_compression() {
        let plan = plan_dataset(
            Path::new("data/lake"),
            LakeLayer::Raw,
            DatasetKind::FinanceTransactions,
        );
        assert_eq!(plan.format, StorageFormat::Jsonl);
        assert_eq!(plan.compression, Compression::None);
        assert_eq!(
            plan.path,
            Path::new("data/lake/raw/finance_transactions/finance_transactions.jsonl")
        );
    }

    #[test]
    fn curated_layers_are_parquet_zstd_ready() {
        let plan = plan_dataset(
            Path::new("data/lake"),
            LakeLayer::Silver,
            DatasetKind::CommercePurchases,
        );
        assert_eq!(plan.format, StorageFormat::Parquet);
        assert_eq!(plan.compression, Compression::Zstd);
        assert_eq!(
            plan.path,
            Path::new("data/lake/silver/commerce_purchases/commerce_purchases.parquet")
        );
    }

    #[test]
    fn partitions_reject_traversal() {
        let error = partitioned_dataset_path(
            Path::new("data/lake"),
            LakeLayer::Gold,
            DatasetKind::FinanceTransactions,
            "../private",
        )
        .expect_err("traversal partition fails");
        assert_eq!(error.field, "partition");
    }
}

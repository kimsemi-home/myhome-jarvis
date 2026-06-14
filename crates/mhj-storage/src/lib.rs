use serde::{Deserialize, Serialize};
use std::error::Error;
use std::fmt;
use std::fs;
use std::path::{Component, Path, PathBuf};

pub const DEFAULT_LAKE_ROOT: &str = "data/lake";
pub const SCHEMA_VERSION: &str = "1.0";

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

    pub fn all() -> [DatasetKind; 2] {
        [
            DatasetKind::FinanceTransactions,
            DatasetKind::CommercePurchases,
        ]
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

#[derive(Debug, Clone, PartialEq, Eq)]
pub enum StorageError {
    InvalidPath {
        field: &'static str,
        message: String,
    },
    Io {
        path: PathBuf,
        message: String,
    },
}

impl fmt::Display for StorageError {
    fn fmt(&self, formatter: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            StorageError::InvalidPath { field, message } => {
                write!(formatter, "{field}: {message}")
            }
            StorageError::Io { path, message } => {
                write!(formatter, "{}: {}", path.display(), message)
            }
        }
    }
}

impl Error for StorageError {}

pub fn default_manifest() -> Result<LakeManifest, StorageError> {
    manifest_for(Path::new(DEFAULT_LAKE_ROOT), &DatasetKind::all())
}

pub fn manifest_for(
    lake_root: &Path,
    datasets: &[DatasetKind],
) -> Result<LakeManifest, StorageError> {
    validate_repo_relative_path("lake_root", lake_root)?;
    let mut files = Vec::new();
    for layer in LakeLayer::all() {
        for dataset in datasets {
            files.push(plan_dataset(lake_root, layer, *dataset)?);
        }
    }
    Ok(LakeManifest {
        lake_root: path_to_repo_string(lake_root),
        schema_version: SCHEMA_VERSION.to_string(),
        files,
    })
}

pub fn plan_dataset(
    lake_root: &Path,
    layer: LakeLayer,
    dataset: DatasetKind,
) -> Result<DatasetFile, StorageError> {
    validate_repo_relative_path("lake_root", lake_root)?;
    let format = format_for_layer(layer);
    let compression = compression_for_format(format);
    let filename = format!("{}.{}", dataset.slug(), format.extension());
    let relative_path = lake_root
        .join(layer.as_str())
        .join(dataset.slug())
        .join(filename);
    validate_repo_relative_path("relative_path", &relative_path)?;
    Ok(DatasetFile {
        layer,
        dataset,
        format,
        compression,
        schema_version: SCHEMA_VERSION.to_string(),
        relative_path: path_to_repo_string(&relative_path),
    })
}

pub fn partitioned_dataset_path(
    lake_root: &Path,
    layer: LakeLayer,
    dataset: DatasetKind,
    partition: &str,
) -> Result<PathBuf, StorageError> {
    validate_partition(partition)?;
    let file = plan_dataset(lake_root, layer, dataset)?;
    let planned_path = PathBuf::from(&file.relative_path);
    let filename = planned_path
        .file_name()
        .ok_or_else(|| StorageError::InvalidPath {
            field: "relative_path",
            message: "dataset path must include a file name".to_string(),
        })?;
    let partitioned = planned_path.with_file_name(partition).join(filename);
    validate_repo_relative_path("partitioned_path", &partitioned)?;
    Ok(partitioned)
}

pub fn write_raw_jsonl(
    repository_root: &Path,
    lake_root: &Path,
    dataset: DatasetKind,
    jsonl: &str,
) -> Result<PathBuf, StorageError> {
    let file = plan_dataset(lake_root, LakeLayer::Raw, dataset)?;
    let relative_path = PathBuf::from(file.relative_path);
    let destination = repository_root.join(&relative_path);
    let parent = destination
        .parent()
        .ok_or_else(|| StorageError::InvalidPath {
            field: "relative_path",
            message: "raw dataset path must have a parent directory".to_string(),
        })?;
    fs::create_dir_all(parent).map_err(|error| StorageError::Io {
        path: parent.to_path_buf(),
        message: error.to_string(),
    })?;
    let mut data = jsonl.as_bytes().to_vec();
    if !data.ends_with(b"\n") {
        data.push(b'\n');
    }
    fs::write(&destination, data).map_err(|error| StorageError::Io {
        path: destination.clone(),
        message: error.to_string(),
    })?;
    Ok(relative_path)
}

pub fn validate_partition(partition: &str) -> Result<(), StorageError> {
    let path = Path::new(partition);
    validate_repo_relative_path("partition", path)?;
    if partition.contains('/') {
        return Err(StorageError::InvalidPath {
            field: "partition",
            message: "must be one relative partition segment".to_string(),
        });
    }
    Ok(())
}

fn format_for_layer(layer: LakeLayer) -> StorageFormat {
    match layer {
        LakeLayer::Raw => StorageFormat::Jsonl,
        LakeLayer::Bronze | LakeLayer::Silver | LakeLayer::Gold => StorageFormat::Parquet,
    }
}

fn compression_for_format(format: StorageFormat) -> Compression {
    match format {
        StorageFormat::Jsonl => Compression::None,
        StorageFormat::Parquet => Compression::Zstd,
    }
}

fn validate_repo_relative_path(field: &'static str, path: &Path) -> Result<(), StorageError> {
    let rendered = path.to_string_lossy();
    if rendered.trim().is_empty() {
        return Err(StorageError::InvalidPath {
            field,
            message: "must not be empty".to_string(),
        });
    }
    if rendered.contains('\\') {
        return Err(StorageError::InvalidPath {
            field,
            message: "must use forward slashes".to_string(),
        });
    }
    if path.is_absolute() {
        return Err(StorageError::InvalidPath {
            field,
            message: "must be repo-relative".to_string(),
        });
    }
    for component in path.components() {
        match component {
            Component::Normal(segment) if !segment.to_string_lossy().trim().is_empty() => {}
            _ => {
                return Err(StorageError::InvalidPath {
                    field,
                    message: "must not contain traversal or empty segments".to_string(),
                });
            }
        }
    }
    Ok(())
}

fn path_to_repo_string(path: &Path) -> String {
    path.to_string_lossy().replace('\\', "/")
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::process;
    use std::time::{SystemTime, UNIX_EPOCH};

    #[test]
    fn default_manifest_covers_each_dataset_and_layer() {
        let manifest = default_manifest().expect("default manifest builds");

        assert_eq!(manifest.lake_root, DEFAULT_LAKE_ROOT);
        assert_eq!(manifest.schema_version, SCHEMA_VERSION);
        assert_eq!(manifest.files.len(), 8);
        assert!(manifest.files.iter().all(|file| {
            file.relative_path.starts_with("data/lake/")
                && !file.relative_path.contains("..")
                && !Path::new(&file.relative_path).is_absolute()
        }));
    }

    #[test]
    fn raw_files_are_jsonl_without_compression() {
        let file = plan_dataset(
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Raw,
            DatasetKind::FinanceTransactions,
        )
        .expect("raw plan builds");

        assert_eq!(file.format, StorageFormat::Jsonl);
        assert_eq!(file.compression, Compression::None);
        assert_eq!(
            file.relative_path,
            "data/lake/raw/finance_transactions/finance_transactions.jsonl"
        );
    }

    #[test]
    fn curated_files_are_parquet_zstd_plans() {
        let file = plan_dataset(
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Gold,
            DatasetKind::CommercePurchases,
        )
        .expect("gold plan builds");

        assert_eq!(file.format, StorageFormat::Parquet);
        assert_eq!(file.compression, Compression::Zstd);
        assert_eq!(
            file.relative_path,
            "data/lake/gold/commerce_purchases/commerce_purchases.parquet"
        );
    }

    #[test]
    fn lake_root_rejects_absolute_and_traversal_paths() {
        assert!(default_manifest().is_ok());
        assert!(manifest_for(Path::new("/tmp/lake"), &DatasetKind::all()).is_err());
        assert!(manifest_for(Path::new("../lake"), &DatasetKind::all()).is_err());
    }

    #[test]
    fn partitioned_paths_reject_nested_or_unsafe_segments() {
        assert!(partitioned_dataset_path(
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Silver,
            DatasetKind::FinanceTransactions,
            "month=2026-06",
        )
        .is_ok());
        assert!(partitioned_dataset_path(
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Silver,
            DatasetKind::FinanceTransactions,
            "../private",
        )
        .is_err());
        assert!(partitioned_dataset_path(
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Silver,
            DatasetKind::FinanceTransactions,
            "month=2026-06/day=14",
        )
        .is_err());
        assert!(partitioned_dataset_path(
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Silver,
            DatasetKind::FinanceTransactions,
            "month\\2026-06",
        )
        .is_err());
    }

    #[test]
    fn raw_jsonl_writer_materializes_only_raw_fixture_files() {
        let root = temp_root();
        let relative = write_raw_jsonl(
            &root,
            Path::new(DEFAULT_LAKE_ROOT),
            DatasetKind::CommercePurchases,
            "{\"purchase_id\":\"fixture-1\"}",
        )
        .expect("raw writer succeeds");

        assert_eq!(
            relative,
            Path::new("data/lake/raw/commerce_purchases/commerce_purchases.jsonl")
        );
        let written = fs::read_to_string(root.join(&relative)).expect("raw file readable");
        assert_eq!(written, "{\"purchase_id\":\"fixture-1\"}\n");
        let _ = fs::remove_dir_all(root);
    }

    fn temp_root() -> PathBuf {
        let nanos = SystemTime::now()
            .duration_since(UNIX_EPOCH)
            .expect("clock after epoch")
            .as_nanos();
        std::env::temp_dir().join(format!("mhj-storage-{}-{nanos}", process::id()))
    }
}

use mhj_storage::{
    default_manifest, manifest_for, plan_dataset, Compression, DatasetKind, LakeLayer,
    StorageFormat, DEFAULT_LAKE_ROOT, SCHEMA_VERSION,
};
use std::path::Path;

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

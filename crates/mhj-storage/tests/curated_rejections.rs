mod common;

use common::temp_root;
use mhj_storage::{
    inspect_curated_parquet, write_curated_parquet_from_jsonl, DatasetKind, LakeLayer,
    StorageError, DEFAULT_LAKE_ROOT,
};
use std::{fs, path::Path};

#[test]
fn curated_writer_rejects_raw_layer() {
    let root = temp_root();
    let result = write_curated_parquet_from_jsonl(
        &root,
        Path::new(DEFAULT_LAKE_ROOT),
        LakeLayer::Raw,
        DatasetKind::FinanceTransactions,
        "{}",
    );

    assert!(matches!(
        result,
        Err(StorageError::InvalidPath { field: "layer", .. })
    ));
    let _ = fs::remove_dir_all(root);
}

#[test]
fn curated_reader_rejects_raw_layer() {
    let root = temp_root();
    let result = inspect_curated_parquet(
        &root,
        Path::new(DEFAULT_LAKE_ROOT),
        LakeLayer::Raw,
        DatasetKind::FinanceTransactions,
    );

    assert!(matches!(
        result,
        Err(StorageError::InvalidPath { field: "layer", .. })
    ));
    let _ = fs::remove_dir_all(root);
}

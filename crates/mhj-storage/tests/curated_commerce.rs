mod common;

use common::{assert_parquet_file, temp_root};
use mhj_storage::{
    inspect_curated_parquet, write_curated_parquet_from_jsonl, Compression, DatasetKind, LakeLayer,
    StorageFormat, DEFAULT_LAKE_ROOT,
};
use std::{fs, path::Path};

#[test]
fn curated_writer_materializes_commerce_parquet_zstd_fixture() {
    let root = temp_root();
    let report = write_curated_parquet_from_jsonl(
        &root,
        Path::new(DEFAULT_LAKE_ROOT),
        LakeLayer::Gold,
        DatasetKind::CommercePurchases,
        include_str!("../../../fixtures/commerce_purchases.jsonl"),
    )
    .expect("curated commerce writer succeeds");

    assert_eq!(
        report.relative_path,
        "data/lake/gold/commerce_purchases/commerce_purchases.parquet"
    );
    assert_eq!(report.format, StorageFormat::Parquet);
    assert_eq!(report.compression, Compression::Zstd);
    assert_eq!(report.row_count, 3);
    assert_parquet_file(&root.join(&report.relative_path), 3);

    let read_report = inspect_curated_parquet(
        &root,
        Path::new(DEFAULT_LAKE_ROOT),
        LakeLayer::Gold,
        DatasetKind::CommercePurchases,
    )
    .expect("curated commerce reader succeeds");
    assert_eq!(read_report.row_count, 3);
    assert_eq!(read_report.row_group_count, 1);
    assert_eq!(read_report.column_count, 17);
    assert_eq!(read_report.compression, Compression::Zstd);
    let _ = fs::remove_dir_all(root);
}

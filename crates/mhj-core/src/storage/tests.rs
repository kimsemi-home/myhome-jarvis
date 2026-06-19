use super::*;
use std::path::Path;

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

mod common;

use common::temp_root;
use mhj_storage::{write_raw_jsonl, DatasetKind, DEFAULT_LAKE_ROOT};
use std::fs;
use std::path::Path;

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

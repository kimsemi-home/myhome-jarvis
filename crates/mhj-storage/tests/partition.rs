use mhj_storage::{partitioned_dataset_path, DatasetKind, LakeLayer, DEFAULT_LAKE_ROOT};
use std::path::Path;

#[test]
fn partitioned_paths_accept_single_safe_segments() {
    assert!(partitioned_dataset_path(
        Path::new(DEFAULT_LAKE_ROOT),
        LakeLayer::Silver,
        DatasetKind::FinanceTransactions,
        "month=2026-06",
    )
    .is_ok());
}

#[test]
fn partitioned_paths_reject_nested_or_unsafe_segments() {
    for partition in ["../private", "month=2026-06/day=14", "month\\2026-06"] {
        assert!(partitioned_dataset_path(
            Path::new(DEFAULT_LAKE_ROOT),
            LakeLayer::Silver,
            DatasetKind::FinanceTransactions,
            partition,
        )
        .is_err());
    }
}

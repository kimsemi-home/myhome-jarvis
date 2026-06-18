use parquet::basic::Compression as ParquetCompression;
use parquet::file::reader::{FileReader, SerializedFileReader};
use std::fs::{self, File};
use std::path::{Path, PathBuf};
use std::process;
use std::sync::atomic::{AtomicU64, Ordering};
use std::time::{SystemTime, UNIX_EPOCH};

static TEMP_ROOT_COUNTER: AtomicU64 = AtomicU64::new(0);

pub fn temp_root() -> PathBuf {
    let nanos = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .expect("clock after epoch")
        .as_nanos();
    let counter = TEMP_ROOT_COUNTER.fetch_add(1, Ordering::Relaxed);
    std::env::temp_dir().join(format!("mhj-storage-{}-{nanos}-{counter}", process::id()))
}

#[allow(dead_code)]
pub fn assert_parquet_file(path: &Path, expected_rows: i64) {
    let bytes = fs::read(path).expect("parquet file readable");
    assert!(bytes.len() > 8);
    assert_eq!(&bytes[..4], b"PAR1");
    assert_eq!(&bytes[bytes.len() - 4..], b"PAR1");

    let file = File::open(path).expect("parquet file open");
    let reader = SerializedFileReader::new(file).expect("parquet metadata readable");
    let metadata = reader.metadata();
    assert_eq!(metadata.file_metadata().num_rows(), expected_rows);
    assert_eq!(metadata.num_row_groups(), 1);
    assert!(matches!(
        metadata.row_group(0).column(0).compression(),
        ParquetCompression::ZSTD(_)
    ));
}

# Storage

Initial fixtures are JSONL.

Long-term storage is organized as:

- `raw`
- `bronze`
- `silver`
- `gold`

Lake data lives under `data/lake` and is ignored. Parquet+Zstd support belongs
to the Rust storage phase.

Rust skeleton:

- `mhj-core::storage::LakeLayer` models `raw`, `bronze`, `silver`, and `gold`.
- `mhj-core::storage::DatasetKind` currently covers finance transactions and
  commerce purchases.
- Raw dataset plans use JSONL without compression.
- Bronze, silver, and gold dataset plans use Parquet with Zstd compression.
- Partition paths reject empty segments, absolute paths, backslashes, and path
  traversal.

Go daemon read surface:

- `GET /domain/summary` includes the generated storage policy.
- The endpoint is local read-only and does not read from `data/lake`.

# Storage

Initial fixtures are JSONL.

Long-term storage is organized as:

- `raw`
- `bronze`
- `silver`
- `gold`

Lake data lives under `data/lake` and is ignored. Parquet+Zstd support belongs
to the Rust storage phase.

Rust storage boundary:

- `mhj-storage::LakeLayer` models `raw`, `bronze`, `silver`, and `gold`.
- `mhj-storage::DatasetKind` currently covers finance transactions and commerce
  purchases.
- `mhj-storage::default_manifest` emits deterministic dataset file plans for
  each layer and dataset.
- Raw dataset plans use JSONL without compression and have a raw writer smoke
  path for fixture-only local files.
- Bronze, silver, and gold dataset plans use Parquet with Zstd compression in
  the manifest, but Parquet writing remains deferred until the Parquet phase is
  explicitly enabled.
- Lake roots and partition paths reject empty segments, absolute paths,
  backslashes, and path traversal.

Schema evolution policy:

- Schema versions start at `1.0` and are recorded in the storage manifest.
- Additive fields that preserve existing readers increment the minor version.
- Removed fields, renamed fields, changed meaning, or changed physical format
  increment the major version.
- Raw JSONL fixture writes may land before Parquet support, but curated
  bronze/silver/gold writes must not claim success until Parquet+Zstd output is
  actually produced by Rust.
- Go and Flutter may read summaries from generated policy and fixtures, but
  storage writes remain behind the Rust storage boundary.

Go daemon read surface:

- `GET /domain/summary` includes the generated storage policy.
- The endpoint is local read-only and does not read from `data/lake`.

Validation:

```sh
cargo test -p mhj-storage
```

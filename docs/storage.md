# Storage

Initial fixtures are JSONL.

Long-term storage is organized as:

- `raw`
- `bronze`
- `silver`
- `gold`

Lake data lives under `data/lake` and is ignored. Parquet+Zstd support belongs
to the Rust storage phase.

Private operational logs use a lighter archive lane before they ever become
long-term lake data. JSONL ledgers under `data/private` are collected as
private log sources, summarized without payloads, compressed with gzip, and
recorded under `data/private/archive` with a private JSONL manifest. The
compression and archive configuration itself is public-safe evidence because it
proves which local data can be compacted and where the private archive boundary
is.

Evidence noise is also configured as evidence. The storage SSOT records an
enabled noise budget, a maximum noise ratio percent, a low-signal record window,
dedupe keys, and the rule that a noise-budget breach blocks archive promotion.
Public surfaces may report this configuration and counts, but never raw log
payloads.

Rust storage boundary:

- `mhj-storage::LakeLayer` models `raw`, `bronze`, `silver`, and `gold`.
- `mhj-storage::DatasetKind` currently covers finance transactions and commerce
  purchases.
- `mhj-storage::default_manifest` emits deterministic dataset file plans for
  each layer and dataset.
- Raw dataset plans use JSONL without compression and have a raw writer smoke
  path for fixture-only local files.
- Bronze, silver, and gold dataset plans use Parquet with Zstd compression in
  the manifest.
- `mhj-storage::write_curated_parquet_from_jsonl` materializes fixture-only
  finance and commerce JSONL into curated Parquet files with Zstd compression.
  It rejects the raw layer, validates normalized fixture fields, writes only
  repo-relative lake paths, and reports row counts without reading credentials
  or external services.
- `mhj-storage::inspect_curated_parquet` reads curated Parquet metadata for
  repo-relative fixture lake paths and verifies row count, row groups, column
  count, and Zstd compression without exposing row contents.
- Lake roots and partition paths reject empty segments, absolute paths,
  backslashes, and path traversal.

Schema evolution policy:

- Schema versions start at `1.0` and are recorded in the storage manifest.
- Additive fields that preserve existing readers increment the minor version.
- Removed fields, renamed fields, changed meaning, or changed physical format
  increment the major version.
- Raw JSONL fixture writes, curated bronze/silver/gold Parquet+Zstd fixture
  writes, and curated metadata reads are produced by Rust.
- Go and Flutter may read summaries from generated policy and fixtures, but
  storage writes remain behind the Rust storage boundary.

Go daemon read surface:

- `GET /domain/summary` includes the generated storage policy.
- The endpoint is local read-only and does not read from `data/lake`.
- `mhj storage-archive status` exposes a redacted summary of the compression
  lane, archive root, manifest path, source count, and noise budget.
- `mhj codex-cost roi` reuses this redacted archive/noise status so cost
  decisions can verify that local evidence logs are compacted and governed.
- `mhj storage-archive run` executes the local private archive lane. Missing or
  empty sources are skipped, present JSONL sources are scanned for invalid or
  duplicate low-signal records, and sources that pass the noise budget are
  written as `.jsonl.gz` files under `data/private/archive`.
- Each run appends private manifest rows to
  `data/private/archive/manifest.jsonl` with source key/path, archive path,
  input/output bytes, input hash, record/noise counts, and budget verdict.
  Public command output reports the same aggregate metadata without raw log
  payloads or local absolute paths.

Validation:

```sh
cargo test -p mhj-storage
cargo test --workspace
```

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
source list, compression, archive configuration, and noise budget are
public-safe evidence because they prove which local data can be compacted and
where the private archive boundary is.

External evidence uses the same private archive lane. `mhj external-evidence
collect` writes a source manifest under
`data/private/external-evidence/manifest.jsonl`, and `mhj storage-archive run`
archives that manifest through the `external_evidence` source key. Raw fetched
payloads remain in the private external evidence lake and are not emitted by
public status surfaces.

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
  lane, archive root, manifest path, source count, noise budget, config
  evidence hash, and private manifest health. Manifest status is summary-only:
  entry counts, archived/skipped/breach/invalid counts, compression ratio, and
  latest archive timestamps are allowed; raw JSONL rows and gzip payloads are
  not.
- `mhj codex-cost roi` reuses this redacted archive/noise status so cost
  decisions can verify that local evidence logs are compacted and governed.
  ROI also reports manifest counts and compression ratio so Codex usage,
  cache savings, accepted changes, and local archive health are reviewed
  together.
  The archive source list includes Codex cost attribution records, allowing
  scope-level ROI evidence to be compressed without publishing raw subjects.
  It also includes monetization experiment records, finance consent records, and
  authority review request records so revenue hypotheses, read-only household
  finance consent evidence, and human-review queue evidence can be compressed
  without exposing raw revenue, finance details, or private review context. It
  also includes the external evidence manifest so news, economic, trend, GitHub,
  and community intake evidence can be compressed without publishing fetched
  payloads.
- `mhj storage-archive run` executes the local private archive lane. Missing or
  empty sources are skipped, present JSONL sources are scanned for invalid or
  duplicate low-signal records, and sources that pass the noise budget are
  written as `.jsonl.gz` files under `data/private/archive`.
- The run report includes a `policy_evidence` summary with the
  compress-then-archive mode, gzip compression, enabled noise budget,
  threshold/window/dedupe settings, breach-blocking rule, and config evidence
  hash. This makes the noise-budget configuration part of every archive run
  proof, while still keeping raw log rows private.
- If a source input hash and config evidence hash already have an archived
  blob, the run reports a cache hit instead of recompressing the same payload
  or appending duplicate manifest noise.
- Archive, skip, and breach decisions append private manifest rows to
  `data/private/archive/manifest.jsonl` with source key/path, archive path,
  input/output bytes, compression ratio percent, input hash, config evidence
  hash, record/noise counts, and budget verdict.
  Public command output reports the same aggregate metadata without raw log
  payloads or local absolute paths.
- The config evidence hash is derived from the SSOT-declared
  `config_hash_inputs`: `private_log_sources`, `log_archive`, and
  `evidence_noise_budget`. That makes the source list, noise threshold, and
  compression/archive settings part of the evidence for each local archive
  decision without publishing the private log body.
- `mhj assistant status` includes a `storage_archive` summary. It can gate the
  closed loop when the archive policy is not public-safe, the noise budget is
  not ready, or the private manifest records invalid rows or budget breaches.
  A missing private manifest alone does not block a fresh clone or CI run.
  The same summary also exposes per-source `source_health` rows with only source
  keys, latest state/verdict, record/noise counts, compression ratio, archive
  evidence presence, hash-cache-key presence, and health debt booleans. It does
  not expose source paths, archive paths, hashes, raw JSONL rows, or gzip
  payloads.

Validation:

```sh
cargo test -p mhj-storage
cargo test --workspace
```

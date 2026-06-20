# Codex Sustainability Evidence Loop

`CodexSustainabilityLoop` measures whether heavier Codex use is still worth it,
fast enough, and supported by current trend evidence.

The raw ledger is private:

`data/private/codex-sustainability/evidence.jsonl`

The public status is intentionally aggregate-only. It may expose counts, posture,
cycle-time summaries, cache counts, review-gate counts, and freshness state. It
must not expose prompts, transcripts, credentials, local absolute paths, private
finance data, unpublished revenue details, or raw evidence references.

## Record Kinds

`usage_sample` records aggregate cost and operating metrics such as Codex tokens,
Codex coin, GitHub Actions minutes, cache hits/misses, validation failures,
human-review debt, accepted changes, and cache savings.

`cycle_sample` records elapsed cycle minutes for accepted work loops.

`trend_baseline` records a versioned, timestamped trend baseline. The first
baseline metric is `elapsed_cycle_minutes`, so the assistant can detect when its
own median cycle time is slower than the current trend instead of guessing.

`feature_proposal` records the evidence behind optimization or feature proposals,
including cost per accepted change, median cycle time, cache savings, defect or
rework rate, and monetization linkage.

## Review Gates

The assistant raises a gate when evidence is stale or missing, trend baselines are
missing or stale, current cycle time is slower than the trend baseline, usage cost
outpaces accepted value, feature proposals lack evidence, or validation/rework
review debt is open.

CLI:

```sh
mhj codex-sustainability status
mhj codex-sustainability record-quality
```

Daemon:

```http
GET /codex-sustainability/status
```

## Quality Run Capture

`mhj codex-sustainability record-quality` reads the private redacted
`data/private/quality/runs.jsonl` journal and appends aggregate trend evidence
to the private Codex sustainability ledger. It records a trend baseline and
cycle sample derived from the latest successful quality run duration.

The capture output is public-safe: it reports state, counts, repo-relative
ledger paths, the trend baseline version, and non-approval flags. It does not
expose command output, argv, prompts, transcripts, local absolute paths,
credentials, tokens, private Linear URLs, finance payloads, or private notes.

# KnowledgeIndex

KnowledgeIndex is a local lexical index used before planning or semantic code
changes. It reads only repo-local files and optional private offline Linear
records. It does not use Cloud RAG, SaaS retrieval, or external vector
databases.

Indexed areas come from the Common Lisp SSOT:

- `lisp/ssot`
- `generated`
- `cmd`
- `internal`
- `apps/flutter`
- `docs`
- `fixtures`
- `harness/golden`
- `data/private/linear-offline-queue.jsonl` when present

Default search output is redacted. It reports concept names, DDD kinds, bounded
contexts, owners, domain event summaries, harness case contracts, repo-relative
paths, line numbers, matched terms, related Linear issue keys, duplicate
suspicions, and must-read files. It does not print source line snippets or raw
private queue contents.

Useful commands:

- `mhj knowledge search KnowledgeIndex`
- `mhj knowledge search planner`
- `mhj knowledge verify`
- `mhj ddd verify`

Planner status performs the SSOT-configured default KnowledgeIndex lookup
before returning a plan summary. Closed-loop checkpoints include the planner
status, so each loop checkpoint carries the KnowledgeIndex evidence summary.

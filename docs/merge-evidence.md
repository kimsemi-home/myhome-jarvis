# Merge Evidence Policy

The SSOT source is `lisp/ssot/merge-evidence.lisp`; the public generated
artifact is `generated/merge_evidence.generated.json`.

Default behavior: open a PR, validate it, merge it when eligible, and record
completion evidence after merge. The merge preference is to merge after checks
pass, then prove the main-branch result. A feature is not considered fully
closed just because a branch or PR exists.

Merge eligibility requires:

- a clean branch state;
- successful required GitHub Actions checks;
- public-safety scans with no findings;
- no unresolved review gate;
- generated artifact drift checks passing.

Required completion evidence:

- PR URL;
- feature commit;
- merge commit;
- push quality run;
- PR quality run;
- required PR checks;
- main quality run after merge;
- Linear completion comment;
- public-safety scan result;
- private-data scan result;
- merge decision comment.

`mhj merge-evidence status` and daemon `GET /merge-evidence/status` expose only
policy-level counts, readiness, proof flags, and public evidence key names. They
do not return private Linear workspace URLs, local absolute paths, raw review
notes, tokens, credentials, raw prompts, raw transcripts, or private evidence
payloads.

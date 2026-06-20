(in-package #:myhome-jarvis.ssot)

(defun validate-codex-sustainability-policy (policy)
  (require-string-equal (getf policy :context)
                        "CodexSustainabilityEvidenceLoop"
                        "Codex sustainability policy context mismatch")
  (require-private-jsonl (getf policy :private_evidence_ledger)
                         "Codex sustainability ledger must stay private JSONL")
  (require-true (and (getf policy :append_only)
                     (getf policy :public_status_redacted)
                     (getf policy :trend_baselines_versioned))
                "Codex sustainability records must be private and versioned")
  (require-false (getf policy :raw_evidence_public_allowed)
                 "Codex sustainability public status must not expose raw evidence")
  (require-positive-integer (getf policy :evidence_max_age_hours)
                            "Codex sustainability evidence max age required")
  (require-positive-integer (getf policy :trend_baseline_max_age_hours)
                            "Codex sustainability trend max age required")
  (require-positive-integer
   (getf policy :cost_per_accepted_change_review_threshold)
   "Codex sustainability cost threshold required")
  (validate-codex-sustainability-lists policy)
  (require-command policy "mhj codex-sustainability status")
  (require-command policy "mhj codex-sustainability record-quality")
  (require-command policy
                   "mhj codex-sustainability record-proposal <json-payload>"))

(defun validate-codex-sustainability-lists (policy)
  (require-members '("usage_sample" "cycle_sample" "trend_baseline"
                     "feature_proposal")
                   (policy-list policy :record_kinds)
                   "Codex sustainability record kind missing: ~A")
  (require-members '("codex_tokens" "codex_coin" "github_actions_minutes"
                     "elapsed_cycle_minutes" "rework_count" "cache_hit_count"
                     "cache_miss_count" "validation_failure_count"
                     "human_review_debt" "accepted_change_count")
                   (policy-list policy :metrics)
                   "Codex sustainability metric missing: ~A")
  (require-members '("record_count" "trend_posture"
                     "sustainability_posture" "evidence_freshness"
                     "review_gate_count" "cost_per_accepted_change")
                   (policy-list policy :public_summary_fields)
                   "Codex sustainability public summary missing: ~A"))

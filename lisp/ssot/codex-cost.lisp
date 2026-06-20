(in-package #:myhome-jarvis.ssot)

(defparameter *codex-cost-policy*
  (list :context "CodexCostGovernor"
        :version "v1"
        :generated_artifact "generated/codex_cost.generated.json"
        :private_usage_ledger "data/private/codex-cost/usage.jsonl"
        :private_attribution_ledger "data/private/codex-cost/attribution.jsonl"
        :append_only t
        :public_status_redacted t
        :raw_usage_public_allowed nil
        :semantic_hash_inputs #("scope" "unit_kind" "amount" "evidence_refs")
        :unit_kinds #("codex_tokens" "codex_coin" "github_actions_minutes"
                      "external_tool_cost")
        :loop_scopes #("assistant_loop" "linear_project" "repo"
                       "monetization_experiment")
        :record_statuses #("recorded" "review_required" "approved" "rejected")
        :warning_unit_threshold 100000
        :review_unit_threshold 500000
        :required_fields #("at" "scope" "unit_kind" "amount" "status"
                           "evidence_refs")
        :attribution_required_fields #("at" "scope" "subject_key"
                                        "unit_kind" "amount" "basis"
                                        "evidence_refs")
        :attribution_cost_ref_inputs #("unit_kind" "amount" "evidence_refs")
        :roi_accepted_change_inputs #("codex_sustainability_ledger"
                                      "git_merge_commits")
        :roi_merge_log_limit 200
        :attribution_subject_max_length 160
        :attribution_cost_ref_max_length 80
        :allowed_evidence_prefixes #("generated/" "docs/" "cmd/" "internal/"
                                     "apps/flutter/" "lisp/" "crates/"
                                     "fixtures/" "harness/" ".github/"
                                     "data/private/")
        :public_summary_fields #("policy_path" "ledger_path" "exists"
                                 "record_count" "invalid_record_count"
                                 "review_required_count"
                                 "missing_evidence_count" "total_units"
                                 "warning_unit_threshold"
                                 "review_unit_threshold" "budget_state"
                                 "by_unit_kind" "by_scope" "by_status"
                                 "last_observed_at" "checked_at")
        :forbidden_public_fields #("raw_prompt" "raw_transcript" "private_notes"
                                   "evidence_refs" "token" "secret"
                                   "credential" "cookie" "account_id"
                                   "card_number" "local_absolute_path"
                                   "linear_url" "private_evidence")
        :commands #("mhj codex-cost status"
                    "mhj codex-cost record <json-payload>"
                    "mhj codex-cost guard <json-payload>"
                    "mhj codex-cost attribute <json-payload>"
                    "mhj codex-cost roi")))

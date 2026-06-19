(in-package #:myhome-jarvis.ssot)

(defparameter *code-shape-policy*
  (list :context "AgentCluster"
        :version "v1"
        :generated_artifact "generated/code_shape.generated.json"
        :max_file_lines 75
        :public_status_redacted t
        :source_roots #("cmd" "internal" "apps/flutter/lib" "apps/flutter/test"
                        "lisp/ssot" "lisp/scripts" "crates")
        :extensions #(".go" ".dart" ".lisp" ".rs")
        :excluded_prefixes #("data/private/" "generated/" ".git/" "target/"
                             "build/" ".dart_tool/" "node_modules/")
        :legacy_debt_files #((:path "internal/incidents/status.go" :max_lines 389) (:path "internal/evidence/status.go" :max_lines 380)
                              (:path "internal/knowledge/index_test.go" :max_lines 377) (:path "internal/learning/ledger.go" :max_lines 370)
                              (:path "internal/evidencequality/status.go" :max_lines 369)
                              (:path "internal/linear/status.go" :max_lines 320) (:path "internal/commands/registry.go" :max_lines 265)
                              (:path "apps/flutter/test/widget_test.dart" :max_lines 245)
                              (:path "crates/mhj-core/src/finance.rs" :max_lines 231) (:path "internal/scheduler/scheduler.go" :max_lines 218)
                              (:path "internal/supervisor/status.go" :max_lines 215) (:path "crates/mhj-core/src/lib.rs" :max_lines 204)
                              (:path "internal/linear/replay_test.go" :max_lines 185)
                              (:path "crates/mhj-core/src/household.rs" :max_lines 178) (:path "crates/mhj-core/src/storage.rs" :max_lines 177))
        :public_summary_fields #("policy_path" "max_file_lines" "file_count"
                                 "over_budget_count" "legacy_debt_count"
                                 "budget_regression_count" "max_observed_path"
                                 "max_observed_lines" "top_debt"
                                 "regressions" "ok" "checked_at")
        :forbidden_public_fields #("local_absolute_path" "raw_source"
                                   "source_excerpt" "token" "secret"
                                   "credential" "account_id" "card_number"
                                   "linear_url" "private_evidence")
        :commands #("mhj code-shape status")))

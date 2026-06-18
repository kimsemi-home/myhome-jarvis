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
        :legacy_debt_files #((:path "internal/daemon/server_test.go" :max_lines 1488) (:path "apps/flutter/lib/main.dart" :max_lines 1209)
                              (:path "crates/mhj-storage/src/lib.rs" :max_lines 1143) (:path "lisp/ssot/codegen.lisp" :max_lines 1062)
                              (:path "apps/flutter/lib/daemon_client.dart" :max_lines 1018)
                              (:path "apps/flutter/test/daemon_client_test.dart" :max_lines 913) (:path "apps/flutter/lib/snapshot.dart" :max_lines 809)
                              (:path "internal/knowledge/index.go" :max_lines 763) (:path "internal/linear/issues_test.go" :max_lines 707)
                              (:path "internal/daemon/server.go" :max_lines 609) (:path "crates/mhj-finance/src/lib.rs" :max_lines 584)
                              (:path "internal/domain/summary.go" :max_lines 562) (:path "crates/mhj-harness/src/lib.rs" :max_lines 558)
                              (:path "crates/mhj-commerce/src/lib.rs" :max_lines 527) (:path "internal/linear/issues.go" :max_lines 526)
                              (:path "internal/security/security.go" :max_lines 499) (:path "internal/controlplane/status.go" :max_lines 467)
                              (:path "internal/authority/status.go" :max_lines 408) (:path "internal/translation/status.go" :max_lines 406)
                              (:path "crates/mhj-command/src/lib.rs" :max_lines 403) (:path "internal/linear/replay.go" :max_lines 392)
                              (:path "internal/incidents/status.go" :max_lines 389) (:path "internal/evidence/status.go" :max_lines 380)
                              (:path "internal/knowledge/index_test.go" :max_lines 377) (:path "internal/learning/ledger.go" :max_lines 370)
                              (:path "internal/evidencequality/status.go" :max_lines 369) (:path "lisp/ssot/ddd.lisp" :max_lines 378)
                              (:path "internal/linear/status.go" :max_lines 320) (:path "internal/agentcluster/status.go" :max_lines 301)
                              (:path "internal/security/security_test.go" :max_lines 299) (:path "internal/planner/status.go" :max_lines 282)
                              (:path "internal/commands/registry.go" :max_lines 265) (:path "internal/confidence/status.go" :max_lines 256)
                              (:path "apps/flutter/test/widget_test.dart" :max_lines 245) (:path "internal/authority/status_test.go" :max_lines 236)
                              (:path "crates/mhj-core/src/finance.rs" :max_lines 231) (:path "internal/scheduler/scheduler.go" :max_lines 218)
                              (:path "internal/supervisor/status.go" :max_lines 215) (:path "crates/mhj-core/src/lib.rs" :max_lines 204)
                              (:path "internal/commands/registry_test.go" :max_lines 202) (:path "internal/commands/harness.go" :max_lines 197)
                              (:path "internal/controlplane/status_test.go" :max_lines 194) (:path "lisp/ssot/authority.lisp" :max_lines 192)
                              (:path "internal/linear/replay_test.go" :max_lines 185)
                              (:path "crates/mhj-core/src/household.rs" :max_lines 178) (:path "crates/mhj-core/src/storage.rs" :max_lines 177)
                              (:path "internal/translation/status_test.go" :max_lines 167) (:path "crates/mhj-core/src/recommendations.rs" :max_lines 162)
                              (:path "internal/connectors/status.go" :max_lines 156) (:path "internal/planner/status_test.go" :max_lines 154)
                              (:path "internal/evidencequality/status_test.go" :max_lines 152) (:path "internal/audit/command_intent.go" :max_lines 152)
                              (:path "lisp/ssot/agent-cluster.lisp" :max_lines 150) (:path "internal/incidents/status_test.go" :max_lines 150)
                              (:path "apps/flutter/test/snapshot_test.dart" :max_lines 150) (:path "crates/mhj-core/src/commerce.rs" :max_lines 149)
                              (:path "internal/confidence/status_test.go" :max_lines 142) (:path "internal/repo/status.go" :max_lines 141)
                              (:path "internal/qualitylog/runs.go" :max_lines 138) (:path "internal/learning/ledger_test.go" :max_lines 130)
                              (:path "internal/commands/executor.go" :max_lines 129) (:path "internal/evidence/status_test.go" :max_lines 126)
                              (:path "internal/orchestrator/checkpoint_test.go" :max_lines 125) (:path "internal/linear/evidence.go" :max_lines 114)
                              (:path "internal/supervisor/status_test.go" :max_lines 104) (:path "internal/daemon/events.go" :max_lines 102)
                              (:path "lisp/ssot/evidence.lisp" :max_lines 97) (:path "crates/mhj-core/src/benchmark.rs" :max_lines 96))
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

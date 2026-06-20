(in-package #:myhome-jarvis.ssot)

(defparameter *storage-policy*
  (list :fixture_format "jsonl"
        :lake_layers #("raw" "bronze" "silver" "gold")
        :datasets #("finance_transactions" "commerce_purchases")
        :long_term_format "parquet"
        :compression "zstd"
        :private_root "data/lake"
        :private_log_sources
        #((:key "quality" :path "data/private/quality/runs.jsonl"
           :format "jsonl")
          (:key "learning" :path "data/private/learning/observations.jsonl"
           :format "jsonl")
          (:key "audit" :path "data/private/audit/command-intents.jsonl"
           :format "jsonl")
          (:key "linear_write" :path "data/private/linear-write-evidence.jsonl"
           :format "jsonl")
          (:key "codex_cost" :path "data/private/codex-cost/usage.jsonl"
           :format "jsonl")
          (:key "codex_sustainability"
           :path "data/private/codex-sustainability/evidence.jsonl"
           :format "jsonl"))
        :log_archive
        (list :mode "compress_then_archive"
              :compression "gzip"
              :archive_root "data/private/archive"
              :archive_extension ".jsonl.gz"
              :manifest_path "data/private/archive/manifest.jsonl"
              :raw_payload_public_allowed nil
              :config_is_evidence t
              :lifecycle #("collect_jsonl" "redact_summary"
                           "dedupe_low_signal" "compress_gzip"
                           "archive_manifest"))
        :evidence_noise_budget
        (list :enabled t
              :max_noise_ratio_percent 20
              :max_low_signal_records_per_window 10
              :window "per_quality_run"
              :dedupe_key_fields #("source" "kind" "evidence_ref")
              :config_evidence_field "evidence_noise_budget"
              :breach_blocks_archive t)))

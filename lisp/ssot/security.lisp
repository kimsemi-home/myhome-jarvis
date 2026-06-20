(in-package #:myhome-jarvis.ssot)

(defparameter *security-policy*
  (list :allowed_languages *allowed-languages*
        :forbidden_languages *forbidden-languages*
        :private_paths #("data/private" "data/lake" "secrets")
        :forbidden_file_markers #("token" "secret" "credential" "cookie")
        :current_content_scan t
        :current_content_scan_skips_private_paths t
        :private_identity_scan t
        :secret_literal_scan t
        :report_matched_secret_contents nil
        :dry_run_default t
        :default_bind_host "127.0.0.1"
        :local_token_file "data/private/local-token.txt"
        :lan_requires_bearer_token t
        :status_cache
        (list :enabled t
              :path "data/private/security/status-cache.json"
              :mode "history_aggregate_only"
              :key_inputs #("git_head" "generated/security.generated.json"
                            "internal/security/*.go")
              :validation_command "mhj security history"
              :miss_runs_full_history t
              :cache_hit_skips_history_scan t
              :current_scan_always_fresh t
              :raw_findings_public_allowed nil
              :public_safe t)))

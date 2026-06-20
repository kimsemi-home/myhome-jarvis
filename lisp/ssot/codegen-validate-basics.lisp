(in-package #:myhome-jarvis.ssot)

(defun validate-command-catalog ()
  (let ((names (mapcar (lambda (command) (getf command :name)) *commands*)))
    (require-members '("open_youtube" "open_youtube_search" "open_netflix"
                       "open_disney_plus" "open_tving" "open_wavve"
                       "open_coupang_play" "open_ott" "open_url"
                       "volume_set" "volume_up" "volume_down"
                       "display_sleep" "movie_mode" "sleep_mode")
                     names "Missing required command: ~A"))
  (require-false (find "Python" *allowed-languages* :test #'string=)
                 "Python must not be an allowed language"))

(defun validate-foundation-policies ()
  (require-member "raw" (policy-list *storage-policy* :lake_layers)
                  "Storage policy must include raw layer")
  (validate-storage-log-archive-policy *storage-policy*)
  (require-member "spouse" (policy-list *household-policy* :scopes)
                  "Household policy must include spouse scope")
  (require-member "subscription_review"
                  (policy-list *recommendation-policy* :kinds)
                  "Recommendation policy must include subscription review")
  (require-true (getf *scheduler-policy* :crash_recovery)
                "Scheduler policy must require crash recovery"))

(defun validate-security-policy-basics ()
  (require-true (getf *security-policy* :lan_requires_bearer_token)
                "LAN daemon access must require a bearer token")
  (require-true (getf *security-policy* :current_content_scan)
                "Current-tree content scanning must stay enabled")
  (require-true (getf *security-policy*
                      :current_content_scan_skips_private_paths)
                "Current-tree content scanning must skip private paths")
  (require-false (getf *security-policy* :report_matched_secret_contents)
                 "Security reports must not expose matched secret contents"))

(defun validate-connector-policy (policy)
  (require-true (getf policy :fixture_only)
                "Connector readiness must remain fixture-only in public SSOT")
  (require-false (getf policy :real_credentials_allowed)
                 "Connector readiness must not allow real credentials")
  (require-false (getf policy :external_api_calls_allowed)
                 "Connector readiness must not allow external API calls")
  (require-true (> (length (getf policy :connectors)) 0)
                "Connector readiness must declare planned connectors")
  (dolist (connector (policy-list policy :connectors))
    (validate-connector-entry connector)))

(defun validate-connector-entry (connector)
  (let ((key (getf connector :key))
        (allowed (policy-list connector :allowed_operations)))
    (require-string-value key "Connector key is required")
    (require-true (getf connector :fixture_mode)
                  "Connector ~A must stay in fixture mode" key)
    (require-false (find "external_api_call" allowed :test #'string=)
                   "Connector ~A must not allow external API calls" key)
    (require-false (find "credential_request" allowed :test #'string=)
                   "Connector ~A must not allow credential requests" key)))

(in-package #:myhome-jarvis.ssot)

(defun validate-external-evidence-policy (policy)
  (require-string-equal (getf policy :schema_version) "external_evidence/v1"
                        "External evidence schema mismatch")
  (require-true (getf policy :public_safe)
                "External evidence policy must stay public-safe")
  (require-false (getf policy :credentials_allowed)
                 "External evidence must not allow credentials")
  (require-false (getf policy :cookies_allowed)
                 "External evidence must not allow cookies")
  (require-false (getf policy :raw_payload_public_allowed)
                 "External evidence must not publish raw payloads")
  (dolist (path-key '(:private_root :manifest_path :raw_layer_path
                     :bronze_layer_path :silver_layer_path :gold_layer_path
                     :storage_archive_source_path))
    (require-private-path (getf policy path-key)
                          "External evidence path must stay private"))
  (require-positive-integer (getf policy :collection_max_bytes)
                            "External evidence collection max bytes required")
  (validate-external-evidence-sources policy)
  (validate-external-evidence-repo-split policy))

(defun validate-external-evidence-sources (policy)
  (dolist (source (policy-list policy :source_descriptors))
    (require-member (getf source :class) (policy-list policy :source_classes)
                    "External evidence source class is not allowed")
    (require-string-equal (getf source :method) "GET"
                          "External evidence source must use GET")
    (require-true (string-prefix-p "https://" (getf source :url))
                  "External evidence source must use HTTPS")
    (require-positive-integer (getf source :freshness_hours)
                              "External evidence source freshness required")))

(defun validate-external-evidence-repo-split (policy)
  (let ((assessment (getf policy :repo_split_assessment)))
    (require-string-equal (getf assessment :creation_gate)
                          "authority_review_required"
                          "External evidence repo split must be authority gated")
    (require-member "no_raw_payloads" (policy-list assessment :public_repo_rules)
                    "External evidence public repo rules must forbid raw payloads")
    (require-member "private_data_stays_private"
                    (policy-list assessment :public_repo_rules)
                    "External evidence public repo rules must keep data private")))

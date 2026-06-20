(in-package #:myhome-jarvis.ssot)

(defun validate-storage-log-archive-policy (policy)
  (dolist (source (policy-list policy :private_log_sources))
    (require-private-jsonl (getf source :path)
                           "Storage log source must stay private JSONL"))
  (let ((archive (getf policy :log_archive))
        (noise (getf policy :evidence_noise_budget)))
    (validate-storage-archive-config archive)
    (validate-storage-noise-budget noise)))

(defun validate-storage-archive-config (archive)
  (require-string-equal (getf archive :mode) "compress_then_archive"
                        "Storage archive mode mismatch")
  (require-string-equal (getf archive :compression) "gzip"
                        "Storage archive compression mismatch")
  (require-string-equal (getf archive :archive_extension) ".jsonl.gz"
                        "Storage archive extension mismatch")
  (require-private-path (getf archive :archive_root)
                        "Storage archive root must stay private")
  (require-private-jsonl (getf archive :manifest_path)
                         "Storage archive manifest must stay private JSONL")
  (require-false (getf archive :raw_payload_public_allowed)
                 "Storage archive must not publish raw payloads")
  (require-true (getf archive :config_is_evidence)
                "Storage archive config must be evidence"))

(defun validate-storage-noise-budget (noise)
  (require-true (getf noise :enabled)
                "Evidence noise budget must be enabled")
  (require-true (getf noise :breach_blocks_archive)
                "Evidence noise breaches must block archive promotion")
  (require-true (<= (getf noise :max_noise_ratio_percent) 25)
                "Evidence noise ratio budget is too loose")
  (require-true (> (getf noise :max_low_signal_records_per_window) 0)
                "Evidence noise low-signal window must be positive")
  (require-string-equal (getf noise :config_evidence_field) "evidence_noise_budget"
                        "Evidence noise config evidence field mismatch"))

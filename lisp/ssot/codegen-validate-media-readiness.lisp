(in-package #:myhome-jarvis.ssot)

(defun validate-media-readiness-policy (policy)
  (require-string-equal (getf policy :context)
                        "MediaReadinessBenchmark"
                        "Media readiness context mismatch")
  (require-true (and (getf policy :public_status_redacted)
                     (not (getf policy :execute_commands))
                     (not (getf policy :persist_payloads))
                     (not (getf policy :persist_urls)))
                "Media readiness must stay redacted and dry-run only")
  (require-true (> (getf policy :target_planning_latency_ms) 0)
                "Media readiness target latency must be positive")
  (validate-media-readiness-cases (getf policy :cases))
  (require-command policy "mhj media-readiness status"))

(defun validate-media-readiness-cases (cases)
  (require-true (and (vectorp cases) (>= (length cases) 4))
                "Media readiness requires at least four benchmark cases")
  (let ((ids (loop for index from 0 below (length cases)
                   collect (getf (aref cases index) :id))))
    (require-members '("youtube_launch" "youtube_search" "ott_netflix"
                       "playback_readiness")
                     ids
                     "Media readiness case missing: ~A")))

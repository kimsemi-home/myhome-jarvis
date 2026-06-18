(in-package #:myhome-jarvis.ssot)

(defun validate-confidence-policy (policy)
  (require-string-equal (getf policy :context) "AgentCluster"
                        "Confidence policy must belong to AgentCluster")
  (require-string-equal (getf policy :assessor_key) "confidence_assessor"
                        "Confidence assessor key must stay explicit")
  (require-true (getf policy :confidence_is_cap)
                "Confidence must be represented as a cap")
  (require-false (getf policy :self_report_allowed)
                 "Agent confidence self-reporting must stay forbidden")
  (require-true (getf policy :public_status_redacted)
                "Confidence status must stay redacted")
  (require-false (getf policy :raw_evidence_public_allowed)
                 "Confidence status must not expose raw evidence")
  (require-members '("blocked" "low" "medium" "high")
                   (policy-list policy :levels)
                   "Confidence level missing: ~A")
  (require-members '("evidence_graph" "learning_ledger" "quality_gate"
                     "public_safety")
                   (policy-list policy :inputs)
                   "Confidence input missing: ~A")
  (validate-confidence-rules policy)
  (require-members '("level_cap" "self_report_allowed" "evidence_link_count"
                     "public_safety_ok" "checked_at")
                   (policy-list policy :public_summary_fields)
                   "Confidence summary missing field: ~A")
  (require-command policy "mhj confidence status"))

(defun validate-confidence-rules (policy)
  (let ((rules (policy-list policy :cap_rules))
        (levels (policy-list policy :levels)))
    (require-true (> (length rules) 0) "Confidence cap rules are required")
    (dolist (rule rules)
      (let ((key (getf rule :key))
            (when-value (getf rule :when))
            (cap (getf rule :cap)))
        (require-string-value key "Confidence rule key is required")
        (require-string-value when-value
                              "Confidence rule must declare a condition")
        (require-member cap levels
                        "Confidence rule ~A has invalid cap ~A" key cap)))))

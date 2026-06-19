(in-package #:myhome-jarvis.ssot)

(defun validate-pdca-policy (policy)
  (require-string-equal (getf policy :context) "AgentCluster"
                        "PDCA context must be AgentCluster")
  (require-private-jsonl (getf policy :private_cycle_ledger)
                         "PDCA ledger must stay private JSONL")
  (require-true (getf policy :append_only)
                "PDCA ledger must be append-only")
  (require-true (getf policy :public_status_redacted)
                "PDCA status must be redacted")
  (require-false (getf policy :raw_cycle_public_allowed)
                 "PDCA raw cycles must stay private")
  (validate-pdca-steps policy)
  (require-members '("cycle_id" "at" "status" "owner"
                     "plan_ref" "do_ref" "check_ref" "act_ref")
                   (policy-list policy :required_fields)
                   "PDCA required field missing: ~A")
  (require-members '("open" "checking" "acting" "closed")
                   (policy-list policy :allowed_statuses)
                   "PDCA status missing: ~A")
  (require-members '("LearningLedger" "EvidenceGraph" "HumanReviewCapacity"
                     "AuthorityGate" "VerificationEvidenceOps"
                     "QualityLedger")
                   (policy-list policy :evidence_sources)
                   "PDCA evidence source missing: ~A")
  (require-members '("cycle_count" "open_count" "closed_count"
                     "invalid_cycle_count" "ready_step_count" "ready")
                   (policy-list policy :public_summary_fields)
                   "PDCA public summary field missing: ~A")
  (require-command policy "mhj pdca status"))

(defun validate-pdca-steps (policy)
  (let ((steps (policy-list policy :steps)))
    (unless (= (length steps) 4)
      (error "PDCA must have exactly four steps"))
    (loop for step in steps
          for expected in '("plan" "do" "check" "act")
          do (require-string-equal (getf step :id) expected
                                   "PDCA step order mismatch"))
    (dolist (step steps)
      (require-string-value (getf step :role) "PDCA step role is required")
      (require-string-value (getf step :artifact)
                            "PDCA step artifact is required")
      (require-string-value (getf step :command)
                            "PDCA step command is required"))))

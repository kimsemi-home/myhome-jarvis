(in-package #:myhome-jarvis.ssot)

(defun validate-incident-policy (policy)
  (require-string-equal (getf policy :context) "AgentCluster"
                        "Incident policy must belong to AgentCluster")
  (require-private-jsonl (getf policy :private_incident_ledger)
                         "Incident ledger must stay private JSONL")
  (require-true (and (getf policy :append_only)
                     (getf policy :public_status_redacted))
                "Incident ledger must be append-only and redacted")
  (require-false (getf policy :raw_incident_public_allowed)
                 "Incident status must not expose raw incident details")
  (require-positive-integer (getf policy :quarantine_stale_after_hours)
                            "Incident stale threshold must be positive")
  (require-members '("quality_regression" "public_safety" "evidence_gap"
                     "authority_violation" "quarantine" "feedback_loop_gap")
                   (policy-list policy :allowed_kinds)
                   "Incident kind missing: ~A")
  (require-members '("observed" "evidence_recorded" "classified"
                     "owner_assigned" "fix_planned" "verified"
                     "knowledge_updated")
                   (policy-list policy :lifecycle)
                   "Incident lifecycle missing stage: ~A")
  (require-members '("open" "mitigating" "verified" "closed" "quarantined")
                   (policy-list policy :allowed_statuses)
                   "Incident status missing: ~A")
  (require-members '("producer" "independent_reviewer"
                     "adversarial_reviewer" "deterministic_verifier"
                     "governance_steward")
                   (policy-list policy :owner_roles)
                   "Incident owner role missing: ~A")
  (require-members '("none" "quarantined" "release_requested" "released")
                   (policy-list policy :quarantine_states)
                   "Incident quarantine state missing: ~A")
  (validate-incident-fields policy)
  (require-command policy "mhj incidents status"))

(defun validate-incident-fields (policy)
  (require-members '("at" "kind" "stage" "status" "owner_role"
                     "evidence_refs")
                   (policy-list policy :required_fields)
                   "Incident required field missing: ~A")
  (require-members '("count" "open_count" "incident_debt_count"
                     "missing_owner_count" "missing_evidence_ref_count"
                     "stale_quarantine_count" "by_stage" "by_owner_role"
                     "checked_at")
                   (policy-list policy :public_summary_fields)
                   "Incident summary missing field: ~A"))

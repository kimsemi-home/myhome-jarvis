(in-package #:myhome-jarvis.ssot)

(defun validate-authority-policy (policy)
  (require-string-equal (getf policy :context) "AgentCluster"
                        "Authority policy must belong to AgentCluster")
  (require-true (getf policy :public_status_redacted)
                "Authority public status must stay redacted")
  (require-false (getf policy :self_authority_allowed)
                 "Authority policy must not allow self-authority")
  (require-false (getf policy :reasoning_tier_grants_approval)
                 "Reasoning tier alone must not grant approval")
  (require-true (getf policy :public_repo_high_risk_blocked)
                "Public repo mode must block high-risk authority decisions")
  (validate-authority-taxonomy policy)
  (validate-authority-decisions policy)
  (validate-authority-profiles policy)
  (validate-authority-summary policy)
  (validate-authority-review-recording policy)
  (require-command policy "mhj authority status")
  (require-command policy "mhj authority-review brief"))

(defun validate-authority-taxonomy (policy)
  (require-members '("confidence_assessor" "evidence_quality"
                     "incident_lifecycle" "control_plane" "translation"
                     "human_review" "public_safety")
                   (policy-list policy :required_inputs)
                   "Authority input missing: ~A")
  (require-members '("r0_compiler" "r1_low" "r2_medium" "r3_high"
                     "r4_governance")
                   (mapcar (lambda (item) (getf item :key))
                           (policy-list policy :reasoning_tiers))
                   "Authority reasoning tier missing: ~A")
  (require-members '("producer" "independent_reviewer"
                     "adversarial_reviewer" "deterministic_verifier"
                     "governance_steward")
                   (mapcar (lambda (item) (getf item :role))
                           (policy-list policy :role_permissions))
                   "Authority role permission missing: ~A")
  (require-members '("agent_reliability" "reasoning_tier"
                     "ontology_maturity" "evidence_quality"
                     "security_impact" "data_sensitivity" "change_risk"
                     "verification_scope" "lease_status" "quarantine_state"
                     "human_review_capacity")
                   (policy-list policy :domain_attributes)
                   "Authority domain attribute missing: ~A"))

(defun validate-authority-review-recording (policy)
  (require-private-jsonl (getf policy :private_review_request_ledger)
                         "Authority review ledger must stay private JSONL")
  (require-private-jsonl (getf policy :private_approval_decision_ledger)
                         "Authority approval ledger must stay private JSONL")
  (require-members '("request_id" "evidence_ref" "queue_item_ref"
                     "queue_state" "required_review_classes"
                     "approval_granted" "external_writes_allowed"
                     "self_approval_allowed")
                   (policy-list policy :review_record_required_fields)
                   "Authority review record field missing: ~A")
  (require-members '("decision_packet_ref" "decision_packet_context"
                     "decision_packet_checked_at" "scope" "target"
                     "reviewer_boundary" "reviewer_is_requester"
                     "expires_at" "approval_granted"
                     "repo_creation_allowed" "workflow_changes_allowed"
                     "external_writes_allowed" "self_approval_allowed")
                   (policy-list policy :approval_decision_required_fields)
                   "Authority approval decision field missing: ~A")
  (require-command policy "mhj authority-review refresh [linear-ref]")
  (require-command policy "mhj authority-review record <json-payload>")
  (require-command policy "mhj authority-review approval-status")
  (require-command policy "mhj authority-review approve <json-payload>"))

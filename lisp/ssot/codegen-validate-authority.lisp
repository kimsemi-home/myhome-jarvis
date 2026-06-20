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
  (require-command policy "mhj authority status"))

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

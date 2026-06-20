(in-package #:myhome-jarvis.ssot)

(defun validate-authority-decisions (policy)
  (let ((decisions (policy-list policy :decisions)))
    (require-true (> (length decisions) 0) "Authority decisions are required")
    (dolist (decision decisions)
      (let ((key (getf decision :key))
            (risk (getf decision :risk))
            (public-allowed (getf decision :public_repo_allowed)))
        (require-string-value key "Authority decision key is required")
        (require-member risk '("low" "medium" "high")
                        "Authority decision ~A has invalid risk ~A" key risk)
        (require-false (and (string= risk "high") public-allowed)
                       "Authority high-risk decision must not be public: ~A"
                       key)))
    (validate-authority-high-risk-decisions decisions))
  (require-members '("limited" "review_required" "blocked")
                   (policy-list policy :outcomes)
                   "Authority outcome missing: ~A")
  (require-members '("public_safety" "confidence_cap" "evidence_quality"
                     "incident" "control_plane" "translation" "human_review")
                   (policy-list policy :authority_debt_classes)
                   "Authority debt class missing: ~A"))

(defun validate-authority-high-risk-decisions (decisions)
  (dolist (key '("major_ontology_change" "security_boundary_change"
                 "production_change" "evidence_pruning" "quarantine_release"
                 "high_risk_automation"))
    (let ((entry (find key decisions :key (lambda (item) (getf item :key))
                       :test #'string=)))
      (require-true entry "Authority high-risk decision missing: ~A" key)
      (require-false (getf entry :public_repo_allowed)
                     "Authority high-risk decision must stay blocked: ~A"
                     key))))

(defun validate-authority-summary (policy)
  (require-members '("outcome" "active_rule" "allowed_decision_count"
                     "blocked_decision_count" "authority_debt_count"
                     "public_repo_mode" "reasoning_tier_grants_approval"
                     "self_authority_allowed" "public_safety_ok"
                     "confidence_cap" "human_review_debt_count"
                     "human_review_capacity_state" "allowed_decisions"
                     "blocked_decisions" "profile_count"
                     "review_required_profile_count"
                     "public_safety_gated_profile_count"
                     "self_approval_blocked_profile_count" "profile_keys"
                     "review_required_profiles" "checked_at")
                   (policy-list policy :public_summary_fields)
                   "Authority summary missing field: ~A"))

(in-package #:myhome-jarvis.ssot)

(defun validate-review-overload (policy)
  (let ((overload (policy-list policy :overload_policy))
        (frozen '("high_risk_change" "major_ontology_change"
                  "security_boundary_change" "production_change")))
    (require-members '("low_risk_only" "deterministic_verification"
                       "evidence_collection" "incident_response"
                       "revalidation" "high_risk_change"
                       "major_ontology_change" "security_boundary_change"
                       "production_change")
                     (mapcar (lambda (item) (getf item :key)) overload)
                     "Review overload policy missing: ~A")
    (dolist (key frozen)
      (let ((entry (find key overload :key (lambda (item) (getf item :key))
                         :test #'string=)))
        (require-false (getf entry :allowed_when_overloaded)
                       "Review overload policy must freeze key: ~A" key)))))

(defun validate-review-fields (policy)
  (require-members '("at" "item_key" "queue_class" "risk" "status"
                     "requester_role" "reviewer_role" "backup_available"
                     "evidence_refs")
                   (policy-list policy :required_fields)
                   "Review required field missing: ~A")
  (require-members '("count" "open_count" "high_risk_open_count"
                     "invalid_review_count" "missing_evidence_count"
                     "missing_reviewer_count" "backup_available_count"
                     "review_debt_count" "capacity_state" "active_rule"
                     "by_risk" "by_status" "by_reviewer_role"
                     "by_queue_class" "checked_at")
                   (policy-list policy :public_summary_fields)
                   "Review summary missing field: ~A"))

(in-package #:myhome-jarvis.ssot)

(defun validate-merge-evidence-policy (policy)
  (require-string-equal (getf policy :context)
                        "MergeEvidencePolicy"
                        "Merge evidence context mismatch")
  (require-string-equal (getf policy :default_behavior)
                        "merge_when_eligible"
                        "Merge evidence default behavior mismatch")
  (require-string-equal (getf policy :merge_preference)
                        "merge_after_checks_pass"
                        "Merge evidence merge preference mismatch")
  (require-true (getf policy :public_status_redacted)
                "Merge evidence status must be redacted")
  (require-false (getf policy :merge_without_review_allowed)
                 "Merge evidence must not allow unreviewed merges")
  (require-false (getf policy :persist_private_evidence)
                 "Merge evidence must not persist private evidence")
  (require-true (getf policy :post_merge_evidence_required)
                "Merge evidence must require post-merge evidence")
  (require-true (getf policy :linear_completion_required)
                "Merge evidence must require Linear completion evidence")
  (require-true (getf policy :main_quality_run_required)
                "Merge evidence must require main quality run evidence")
  (require-true (getf policy :private_data_scan_required)
                "Merge evidence must require private-data scan evidence")
  (validate-merge-evidence-gates (getf policy :gates))
  (require-members '("pr_url" "feature_commit" "merge_commit"
                     "push_quality_run" "pr_quality_run"
                     "pr_required_checks" "main_quality_run"
                     "linear_completion_comment" "public_safety_scan"
                     "private_data_scan" "merge_decision_comment")
                   (policy-list policy :required_evidence)
                   "Merge evidence required item missing: ~A")
  (require-command policy "mhj merge-evidence status"))

(defun validate-merge-evidence-gates (gates)
  (require-true (and (vectorp gates) (>= (length gates) 5))
                "Merge evidence requires at least five gates")
  (let ((keys (loop for index from 0 below (length gates)
                    collect (getf (aref gates index) :key))))
    (require-members '("clean_branch" "required_checks_success"
                       "public_safety_passed" "review_gate_clear"
                       "generated_drift_clear")
                     keys
                     "Merge evidence gate missing: ~A")))

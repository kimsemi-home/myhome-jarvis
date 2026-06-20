(in-package #:myhome-jarvis.ssot)

(defparameter *merge-evidence-gates*
  #((:key "clean_branch" :label "Clean branch"
     :evidence "git status --short --branch" :required t :blocks_merge t)
    (:key "required_checks_success" :label "Required checks success"
     :evidence "GitHub Actions required checks" :required t :blocks_merge t)
    (:key "public_safety_passed" :label "Public safety passed"
     :evidence "mhj security check and mhj security history" :required t :blocks_merge t)
    (:key "review_gate_clear" :label "Review gate clear"
     :evidence "unresolved review threads or explicit self-review" :required t :blocks_merge t)
    (:key "generated_drift_clear" :label "Generated drift clear"
     :evidence "codegen and generated artifact diff checks" :required t :blocks_merge t)))

(defparameter *merge-evidence-required-items*
  #("pr_url" "feature_commit" "merge_commit" "push_quality_run"
    "pr_quality_run" "main_quality_run" "linear_completion_comment"
    "public_safety_scan"))

(defparameter *merge-evidence-policy*
  (list :context "MergeEvidencePolicy"
        :version "v1"
        :generated_artifact "generated/merge_evidence.generated.json"
        :default_behavior "merge_when_eligible"
        :public_status_redacted t
        :merge_without_review_allowed nil
        :persist_private_evidence nil
        :gates *merge-evidence-gates*
        :required_evidence *merge-evidence-required-items*
        :public_summary_fields #("context" "version" "policy_path"
                                 "default_behavior" "eligible_gate_count"
                                 "required_evidence_count"
                                 "missing_required_evidence_count"
                                 "merge_ready"
                                 "merge_blocked_until_evidence"
                                 "checked_at")
        :forbidden_public_fields #("private_linear_url" "local_absolute_path"
                                   "token" "secret" "credential" "cookie"
                                   "raw_review_notes" "private_evidence"
                                   "raw_prompt" "raw_transcript")
        :commands #("mhj merge-evidence status")))

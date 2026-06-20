(in-package #:myhome-jarvis.ssot)

(defun validate-assistant-vision-policy (policy)
  (require-string-equal (getf policy :context)
                        "AssistantVision"
                        "Assistant vision context mismatch")
  (require-string-value (getf policy :mission)
                        "Assistant vision mission is required")
  (require-string-value (getf policy :operating_mode)
                        "Assistant vision operating mode is required")
  (require-members '("local_media_concierge"
                     "household_finance_copilot"
                     "shorts_factory_control_plane"
                     "monetization_console"
                     "codex_cost_governor"
                     "self_improvement_loop")
                   (assistant-pillar-keys policy)
                   "Assistant capability pillar missing: ~A")
  (require-members '("local_first" "public_safe_by_default"
                     "private_data_stays_private" "no_self_approval"
                     "cost_tracking_before_external_scale")
                   (policy-list policy :guardrails)
                   "Assistant guardrail missing: ~A")
  (validate-assistant-terms policy)
  t)

(defun assistant-pillar-keys (policy)
  (mapcar (lambda (pillar) (getf pillar :key))
          (policy-list policy :capability_pillars)))

(defun validate-assistant-terms (policy)
  (require-members '("intent" "capability" "evidence" "decision"
                     "cost_unit" "monetization_loop" "repo_factory"
                     "merge_evidence" "household_scope")
                   (mapcar (lambda (term) (getf term :key))
                           (policy-list policy :universal_terms))
                   "Assistant universal term missing: ~A"))

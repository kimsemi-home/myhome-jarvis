(in-package #:myhome-jarvis.ssot)

(defun validate-authority-profiles (policy)
  (let ((profiles (policy-list policy :assistant_authority_profiles)))
    (require-members '("local_media_concierge" "household_finance_copilot"
                       "shorts_factory_control_plane" "monetization_console"
                       "codex_cost_governor" "self_improvement_loop")
                     (mapcar (lambda (item) (getf item :key)) profiles)
                     "Assistant authority profile missing: ~A")
    (dolist (profile profiles)
      (require-false (getf profile :self_approval_allowed)
                     "Assistant profile must block self approval: ~A"
                     (getf profile :key)))
    (validate-risky-assistant-profiles profiles)))

(defun validate-risky-assistant-profiles (profiles)
  (let ((finance (find "household_finance_copilot" profiles
                       :key (lambda (item) (getf item :key)) :test #'string=))
        (repo (find "shorts_factory_control_plane" profiles
                    :key (lambda (item) (getf item :key)) :test #'string=))
        (money (find "monetization_console" profiles
                     :key (lambda (item) (getf item :key)) :test #'string=)))
    (require-true (getf finance :requires_human_review)
                  "Finance assistant profile must require review")
    (require-true (and (getf repo :requires_human_review)
                       (getf repo :public_safety_gate_required)
                       (getf repo :public_repo_change_gate_required)
                       (getf repo :workflow_change_gate_required))
                  "Repo factory profile must require public safety gates")
    (require-true (and (getf money :requires_human_review)
                       (getf money :public_safety_gate_required))
                  "Monetization profile must require review and safety gate")))

(in-package #:myhome-jarvis.ssot)

(defun assistant-authority-profile
    (key profile sensitivity review safety repo workflow external verifier evidence decisions)
  (list :key key
        :authority_profile profile
        :data_sensitivity sensitivity
        :requires_human_review review
        :public_safety_gate_required safety
        :public_repo_change_gate_required repo
        :workflow_change_gate_required workflow
        :external_writes_allowed external
        :verifier_separation_required verifier
        :self_approval_allowed nil
        :required_evidence evidence
        :allowed_decisions decisions))

(defparameter *assistant-authority-profiles*
  (vector
   (assistant-authority-profile
    "local_media_concierge" "local_interactive" "household_low"
    nil nil nil nil nil nil #("harness home")
    #("read_status" "evidence_collection" "deterministic_verification"))
   (assistant-authority-profile
    "household_finance_copilot" "finance_review_only" "household_finance"
    t nil nil nil nil t #("harness finance" "consent ledger")
    #("read_status" "evidence_collection" "deterministic_verification"))
   (assistant-authority-profile
    "shorts_factory_control_plane" "public_repo_review_required" "public_repo"
    t t t t nil t #("verification graph" "public safety scan")
    #("read_status" "evidence_collection" "deterministic_verification"))
   (assistant-authority-profile
    "monetization_console" "monetization_review_required" "revenue_experiment"
    t t nil t nil t #("experiment ledger" "cost governor")
    #("read_status" "evidence_collection" "deterministic_verification"))
   (assistant-authority-profile
    "codex_cost_governor" "local_readonly" "cost_metadata"
    nil nil nil nil nil nil #("private cost ledger")
    #("read_status" "evidence_collection" "deterministic_verification"))
   (assistant-authority-profile
    "self_improvement_loop" "authority_gated" "governance"
    t t nil t nil t #("pdca" "learning ledger" "incident lifecycle")
    #("read_status" "evidence_collection" "deterministic_verification"
      "revalidation"))))

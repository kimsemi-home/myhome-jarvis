(in-package #:myhome-jarvis.ssot)

(defparameter *authority-role-permissions*
  #((:role "producer"
     :may #("propose" "attach_evidence")
     :must_not #("self_approve" "final_confidence"))
    (:role "independent_reviewer"
     :may #("review_mapping" "review_risk")
     :must_not #("review_after_reading_adversarial_answer" "self_approve"))
    (:role "adversarial_reviewer"
     :may #("challenge_evidence" "find_authority_violation")
     :must_not #("mutate_artifact" "self_approve"))
    (:role "deterministic_verifier"
     :may #("run_checks" "verify_artifacts")
     :must_not #("change_policy" "approve_semantics"))
    (:role "governance_steward"
     :may #("gate_authority" "assign_revalidation")
     :must_not #("solo_major_ontology_change" "erase_evidence"))))

(defparameter *authority-domain-attributes*
  #("agent_reliability" "reasoning_tier" "ontology_maturity"
    "evidence_quality" "security_impact" "data_sensitivity"
    "change_risk" "verification_scope" "lease_status"
    "quarantine_state" "human_review_capacity"))

(in-package #:myhome-jarvis.ssot)

(defparameter *authority-reasoning-tiers*
  #((:key "r0_compiler"
     :label "R0 Compiler"
     :may #("deterministic_transform" "artifact_check")
     :must_not #("approve" "change_authority"))
    (:key "r1_low"
     :label "R1 Low Reasoning"
     :may #("small_candidate" "local_summary")
     :must_not #("approve" "raise_confidence"))
    (:key "r2_medium"
     :label "R2 Medium Reasoning"
     :may #("multi_file_candidate" "verification_plan_candidate")
     :must_not #("approve" "select_reviewer"))
    (:key "r3_high"
     :label "R3 High Reasoning"
     :may #("root_cause_candidate" "ontology_conflict_candidate")
     :must_not #("approve" "bypass_gate"))
    (:key "r4_governance"
     :label "R4 Governance Reasoning"
     :may #("policy_review" "authority_review")
     :must_not #("solo_approve" "replace_human_approval"))))

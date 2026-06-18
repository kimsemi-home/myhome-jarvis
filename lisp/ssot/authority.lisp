(in-package #:myhome-jarvis.ssot)

(defparameter *authority-policy*
  (list :context "AgentCluster"
        :version "v1"
        :generated_artifact "generated/authority.generated.json"
        :public_status_redacted t
        :self_authority_allowed nil
        :reasoning_tier_grants_approval nil
        :public_repo_high_risk_blocked t
        :required_inputs #("confidence_assessor"
                           "evidence_quality"
                           "incident_lifecycle"
                           "control_plane"
                           "translation"
                           "human_review"
                           "public_safety")
        :reasoning_tiers #((:key "r0_compiler"
                             :label "R0 Compiler"
                             :may #("deterministic_transform"
                                    "artifact_check")
                             :must_not #("approve"
                                         "change_authority"))
                            (:key "r1_low"
                             :label "R1 Low Reasoning"
                             :may #("small_candidate"
                                    "local_summary")
                             :must_not #("approve"
                                         "raise_confidence"))
                            (:key "r2_medium"
                             :label "R2 Medium Reasoning"
                             :may #("multi_file_candidate"
                                    "verification_plan_candidate")
                             :must_not #("approve"
                                         "select_reviewer"))
                            (:key "r3_high"
                             :label "R3 High Reasoning"
                             :may #("root_cause_candidate"
                                    "ontology_conflict_candidate")
                             :must_not #("approve"
                                         "bypass_gate"))
                            (:key "r4_governance"
                             :label "R4 Governance Reasoning"
                             :may #("policy_review"
                                    "authority_review")
                             :must_not #("solo_approve"
                                         "replace_human_approval")))
        :role_permissions #((:role "producer"
                              :may #("propose"
                                     "attach_evidence")
                              :must_not #("self_approve"
                                          "final_confidence"))
                             (:role "independent_reviewer"
                              :may #("review_mapping"
                                     "review_risk")
                              :must_not #("review_after_reading_adversarial_answer"
                                          "self_approve"))
                             (:role "adversarial_reviewer"
                              :may #("challenge_evidence"
                                     "find_authority_violation")
                              :must_not #("mutate_artifact"
                                          "self_approve"))
                             (:role "deterministic_verifier"
                              :may #("run_checks"
                                     "verify_artifacts")
                              :must_not #("change_policy"
                                          "approve_semantics"))
                             (:role "governance_steward"
                              :may #("gate_authority"
                                     "assign_revalidation")
                              :must_not #("solo_major_ontology_change"
                                          "erase_evidence")))
        :domain_attributes #("agent_reliability"
                             "reasoning_tier"
                             "ontology_maturity"
                             "evidence_quality"
                             "security_impact"
                             "data_sensitivity"
                             "change_risk"
                             "verification_scope"
                             "lease_status"
                             "quarantine_state"
                             "human_review_capacity")
        :decisions #((:key "read_status"
                      :risk "low"
                      :public_repo_allowed t
                      :requires_human_review nil
                      :allowed_when_blocked t)
                     (:key "evidence_collection"
                      :risk "low"
                      :public_repo_allowed t
                      :requires_human_review nil
                      :allowed_when_blocked t)
                     (:key "deterministic_verification"
                      :risk "low"
                      :public_repo_allowed t
                      :requires_human_review nil
                      :allowed_when_blocked t)
                     (:key "revalidation"
                      :risk "low"
                      :public_repo_allowed t
                      :requires_human_review nil
                      :allowed_when_blocked t)
                     (:key "low_risk_fixture_change"
                      :risk "medium"
                      :public_repo_allowed t
                      :requires_human_review nil
                      :allowed_when_blocked nil)
                     (:key "incident_response"
                      :risk "medium"
                      :public_repo_allowed t
                      :requires_human_review t
                      :allowed_when_blocked t)
                     (:key "major_ontology_change"
                      :risk "high"
                      :public_repo_allowed nil
                      :requires_human_review t
                      :allowed_when_blocked nil)
                     (:key "security_boundary_change"
                      :risk "high"
                      :public_repo_allowed nil
                      :requires_human_review t
                      :allowed_when_blocked nil)
                     (:key "production_change"
                      :risk "high"
                      :public_repo_allowed nil
                      :requires_human_review t
                      :allowed_when_blocked nil)
                     (:key "evidence_pruning"
                      :risk "high"
                      :public_repo_allowed nil
                      :requires_human_review t
                      :allowed_when_blocked nil)
                     (:key "quarantine_release"
                      :risk "high"
                      :public_repo_allowed nil
                      :requires_human_review t
                      :allowed_when_blocked nil)
                     (:key "high_risk_automation"
                      :risk "high"
                      :public_repo_allowed nil
                      :requires_human_review t
                      :allowed_when_blocked nil))
        :outcomes #("limited"
                    "review_required"
                    "blocked")
        :authority_debt_classes #("public_safety"
                                  "confidence_cap"
                                  "evidence_quality"
                                  "incident"
                                  "control_plane"
                                  "translation"
                                  "human_review")
        :public_summary_fields #("policy_path"
                                 "outcome"
                                 "active_rule"
                                 "input_count"
                                 "decision_count"
                                 "allowed_decision_count"
                                 "blocked_decision_count"
                                 "authority_debt_count"
                                 "public_repo_mode"
                                 "reasoning_tier_grants_approval"
                                 "self_authority_allowed"
                                 "public_safety_ok"
                                 "confidence_cap"
                                 "evidence_quality_debt_count"
                                 "incident_debt_count"
                                 "control_plane_debt_count"
                                 "translation_debt_count"
                                 "human_review_debt_count"
                                 "human_review_capacity_state"
                                 "allowed_decisions"
                                 "blocked_decisions"
                                 "by_risk"
                                 "checked_at")
        :forbidden_public_fields #("raw_rationale"
                                   "raw_evidence"
                                   "evidence_ref"
                                   "evidence_refs"
                                   "raw_prompt"
                                   "raw_transcript"
                                   "token"
                                   "secret"
                                   "credential"
                                   "cookie"
                                   "account_id"
                                   "card_number"
                                   "local_absolute_path"
                                   "linear_url"
                                   "private_evidence")
        :commands #("mhj authority status")))

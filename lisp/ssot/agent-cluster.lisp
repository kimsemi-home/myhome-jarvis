(in-package #:myhome-jarvis.ssot)

(defparameter *agent-cluster-policy*
  (list :context "AgentCluster"
        :version "v1"
        :public_safe t
        :raw_transcript_storage_allowed nil
        :private_data_in_evidence_allowed nil
        :external_agent_execution_allowed nil
        :self_approval_allowed nil
        :confidence_self_report_allowed nil
        :authority_gate_required t
        :evidence_flow #("reality" "observation" "evidence" "interpretation" "claim" "rulebook" "design" "code" "verification_evidence" "knowledge_update")
        :ontology_rules #("term_requires_context" "term_requires_version" "translation_requires_manifest" "loss_requires_ledger" "ssot_defect_requires_impact_analysis")
        :agent_roles #((:key "producer" :label "Producer agent" :reasoning_tier "R2" :authority "propose"
                        :must_produce #("change_candidate" "verification_plan" "evidence_links")
                        :must_not #("self_approve" "declare_final_confidence" "bypass_authority_gate"))
                       (:key "independent_reviewer" :label "Independent reviewer" :reasoning_tier "R3" :authority "review"
                        :must_produce #("risk_findings" "ontology_mapping_check")
                        :must_not #("reuse_producer_reasoning" "approve_without_evidence"))
                       (:key "adversarial_reviewer" :label "Adversarial reviewer" :reasoning_tier "R3" :authority "challenge"
                        :must_produce #("failure_modes" "missing_evidence" "rollback_risks")
                        :must_not #("block_without_reason" "skip_security_boundary"))
                       (:key "deterministic_verifier" :label "Deterministic verifier" :reasoning_tier "R0" :authority "verify"
                        :must_produce #("command_result" "artifact_diff" "quality_status")
                        :must_not #("infer_unobserved_success" "mutate_external_state"))
                       (:key "governance_steward" :label "Governance steward" :reasoning_tier "R4" :authority "gate"
                        :must_produce #("approval_record" "revalidation_date" "owner_assignment")
                        :must_not #("replace_human_approval" "erase_audit_evidence")))
        :sidecars #((:key "verification" :label "Verification sidecar" :checks #("contract" "tests" "generated_artifacts"))
                    (:key "confidence" :label "Confidence assessor" :checks #("evidence_links" "coverage" "agent_reliability"))
                    (:key "evidence_quality" :label "Evidence quality assessor" :checks #("age" "schema_version" "mapping_confidence"))
                    (:key "security_audit" :label "Security audit sidecar" :checks #("authority_boundary" "public_safety" "secret_scan"))
                    (:key "translation_verifier" :label "Translation verifier" :checks #("source_context" "target_context" "loss_ledger"))
                    (:key "control_plane_verifier" :label "Control plane verifier" :checks #("orchestration_manifest" "policy_hash" "lease_state")))
        :incident_lifecycle #("observed" "evidence_recorded" "classified" "owner_assigned" "fix_planned" "verified" "knowledge_updated")
        :debt_types #("evidence_debt" "review_debt" "revalidation_debt" "quarantine_debt" "translation_debt")
        :quarantine_triggers #("fabricated_evidence" "prompt_injection_suspected" "authority_violation" "security_boundary_break" "tool_output_integrity_risk" "reviewer_isolation_break")
        :failure_conditions #("contextless_concept" "versionless_ontology" "missing_evidence_link" "self_assessed_confidence" "missing_orchestration_manifest" "lease_expired_artifact" "lossy_translation_without_ledger" "feedback_loop_gap")
        :signals #((:key "evidence_first" :label "Evidence first" :status "active" :evidence "observation and evidence precede code")
                   (:key "authority_gated" :label "Authority gated" :status "gated" :evidence "producer, reviewer, verifier, and steward roles are separated")
                   (:key "feedback_loop" :label "Feedback loop" :status "tracked" :evidence "incidents must end in verification and knowledge update"))
        :forbidden_public_fields #("token" "secret" "credential" "cookie" "raw_prompt" "raw_transcript" "account_id" "card_number" "local_absolute_path")
        :commands #("mhj agent-cluster status")))

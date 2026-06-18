(in-package #:myhome-jarvis.ssot)

(defparameter *translation-policy*
  (list :context "AgentCluster"
        :version "v1"
        :generated_artifact "generated/translation.generated.json"
        :private_loss_ledger "data/private/translation/losses.jsonl"
        :private_manifest_root "data/private/translation/manifests"
        :manifest_required t
        :public_status_redacted t
        :raw_loss_public_allowed nil
        :allowed_contexts #("HomeControl" "HouseholdFinance" "CommerceIntelligence"
                            "ConnectorReadiness" "AgentCluster" "StorageLake"
                            "SecurityPolicy" "AgentOps" "KnowledgeIndex")
        :required_manifest_fields #("source_context" "target_context" "source_version"
                                    "target_version" "preserved_rules" "known_losses"
                                    "owner" "evidence_refs")
        :loss_levels #("l0_none" "l1_note" "l2_degraded" "l3_review_required"
                       "l4_forbidden")
        :allowed_loss_categories #("none" "mapping_gap" "version_drift" "field_drop"
                                   "precision_loss" "review_needed" "authority"
                                   "security_boundary" "user_consent"
                                   "deletion_semantics" "audit_record"
                                   "legal_obligation" "financial_commitment")
        :forbidden_loss_categories #("authority" "security_boundary" "user_consent"
                                     "deletion_semantics" "audit_record"
                                     "legal_obligation" "financial_commitment")
        :allowed_evidence_prefixes #("data/private/" "generated/" "docs/" "cmd/"
                                     "internal/" "apps/flutter/" "lisp/" "crates/"
                                     "fixtures/" "harness/" ".github/")
        :public_summary_fields #("policy_path" "ledger_path" "manifest_root"
                                 "ledger_exists" "manifest_root_exists" "manifest_count"
                                 "invalid_manifest_count" "missing_manifest_count"
                                 "loss_count" "open_loss_count" "closed_loss_count"
                                 "invalid_loss_count" "open_debt_count"
                                 "forbidden_loss_count" "by_level" "by_source_context"
                                 "by_target_context" "last_observed_at" "checked_at")
        :forbidden_public_fields #("summary" "semantic_notes" "raw_mapping" "known_losses"
                                   "evidence_refs" "token" "secret" "credential"
                                   "cookie" "raw_prompt" "raw_transcript" "account_id"
                                   "card_number" "local_absolute_path")
        :commands #("mhj translation status")))

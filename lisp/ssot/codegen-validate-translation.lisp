(in-package #:myhome-jarvis.ssot)

(defun validate-translation-policy (policy)
  (require-string-equal (getf policy :context) "AgentCluster"
                        "Translation policy must belong to AgentCluster")
  (require-private-path (getf policy :private_loss_ledger)
                        "Translation loss ledger must stay under data/private")
  (require-private-path (getf policy :private_manifest_root)
                        "Translation manifest root must stay under data/private")
  (require-true (getf policy :manifest_required)
                "Translation manifests must be required for context moves")
  (require-true (getf policy :public_status_redacted)
                "Translation public status must stay redacted")
  (require-false (getf policy :raw_loss_public_allowed)
                 "Translation status must not expose raw semantic losses")
  (require-members '("AgentCluster" "KnowledgeIndex" "AgentOps"
                     "SecurityPolicy")
                   (policy-list policy :allowed_contexts)
                   "Translation policy missing context: ~A")
  (require-members '("source_context" "target_context" "source_version"
                     "target_version" "preserved_rules" "known_losses"
                     "owner" "evidence_refs")
                   (policy-list policy :required_manifest_fields)
                   "Translation manifest missing required field: ~A")
  (require-members '("l0_none" "l1_note" "l2_degraded" "l3_review_required"
                     "l4_forbidden")
                   (policy-list policy :loss_levels)
                   "Translation loss level missing: ~A")
  (require-members '("mapping_gap" "authority" "security_boundary"
                     "financial_commitment")
                   (policy-list policy :allowed_loss_categories)
                   "Translation loss category missing: ~A")
  (require-members '("authority" "security_boundary" "user_consent"
                     "deletion_semantics" "audit_record" "legal_obligation"
                     "financial_commitment")
                   (policy-list policy :forbidden_loss_categories)
                   "Translation forbidden loss category missing: ~A")
  (require-members '("open_debt_count" "forbidden_loss_count"
                     "invalid_manifest_count" "missing_manifest_count"
                     "checked_at")
                   (policy-list policy :public_summary_fields)
                   "Translation summary missing field: ~A")
  (require-command policy "mhj translation status"))

(in-package #:myhome-jarvis.ssot)

(defun validate-evidence-quality-policy (policy)
  (require-string-equal (getf policy :context) "AgentCluster"
                        "Evidence quality policy must belong to AgentCluster")
  (require-private-jsonl (getf policy :private_snapshot_ledger)
                         "Evidence quality ledger must stay private JSONL")
  (require-true (and (getf policy :append_only)
                     (getf policy :public_status_redacted))
                "Evidence quality snapshots must be append-only and redacted")
  (require-false (getf policy :raw_snapshot_public_allowed)
                 "Evidence quality status must not expose raw snapshots")
  (require-positive-integer (getf policy :stale_after_hours)
                            "Evidence quality stale threshold must be positive")
  (require-members '("high" "medium" "low" "blocked")
                   (policy-list policy :quality_levels)
                   "Evidence quality level missing: ~A")
  (require-members '("high" "medium" "low" "unknown")
                   (policy-list policy :mapping_confidence_levels)
                   "Evidence quality mapping confidence missing: ~A")
  (require-members '("root_cause" "confidence_assessment" "incident_review"
                     "release_gate" "conformance" "revalidation")
                   (policy-list policy :allowed_purposes)
                   "Evidence quality purpose missing: ~A")
  (require-members '("age" "schema_version_change" "ontology_version_change"
                     "counter_evidence" "security_incident" "quarantine"
                     "translation_loss")
                   (policy-list policy :reassessment_reasons)
                   "Evidence quality reassessment reason missing: ~A")
  (validate-evidence-quality-fields policy)
  (require-command policy "mhj evidence-quality status"))

(defun validate-evidence-quality-fields (policy)
  (require-members '("at" "evidence_ref" "purpose" "quality_level"
                     "schema_version" "ontology_version" "mapping_confidence"
                     "assessed_by" "reassessment_reasons")
                   (policy-list policy :required_fields)
                   "Evidence quality required field missing: ~A")
  (require-members '("snapshot_count" "invalid_snapshot_count"
                     "reassessment_debt_count" "missing_evidence_count"
                     "stale_snapshot_count" "low_quality_count"
                     "blocked_quality_count" "mapping_drift_count"
                     "by_quality_level" "by_mapping_confidence" "checked_at")
                   (policy-list policy :public_summary_fields)
                   "Evidence quality summary missing field: ~A"))

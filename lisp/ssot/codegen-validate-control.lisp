(in-package #:myhome-jarvis.ssot)

(defun validate-control-plane-policy (policy)
  (require-string-equal (getf policy :context) "AgentOps"
                        "Control-plane policy must belong to AgentOps")
  (require-private-jsonl (getf policy :private_manifest_ledger)
                         "Control-plane ledger must stay private JSONL")
  (require-true (and (getf policy :manifest_required)
                     (getf policy :append_only)
                     (getf policy :public_status_redacted))
                "Control-plane manifests must be append-only and redacted")
  (require-false (getf policy :raw_rationale_public_allowed)
                 "Control-plane status must not expose raw rationale")
  (require-true (getf policy :verifier_separation_required)
                "Control-plane verifier separation must be required")
  (require-string-equal (getf policy :verifier_generated_artifact)
                        "generated/control_plane_verification.generated.json"
                        "Control-plane verifier artifact must be generated")
  (require-string-equal (getf policy :verification_command)
                        "mhj control-plane verify"
                        "Control-plane verifier command must stay executable")
  (let ((min-lease (getf policy :min_lease_seconds))
        (max-lease (getf policy :max_lease_seconds)))
    (require-true (and (integerp min-lease)
                       (integerp max-lease)
                       (< min-lease max-lease))
                  "Control-plane lease bounds must be valid"))
  (require-members '("loop_once" "loop_worker_cycle" "checkpoint_write")
                   (policy-list policy :allowed_decision_kinds)
                   "Control-plane decision kind missing: ~A")
  (require-members '("local_readonly" "external_write_gated")
                   (policy-list policy :allowed_authority_profiles)
                   "Control-plane authority profile missing: ~A")
  (require-members '("issued" "active" "finished" "aborted" "quarantined")
                   (policy-list policy :allowed_lease_statuses)
                   "Control-plane lease status missing: ~A")
  (validate-control-plane-fields policy)
  (require-control-plane-verifier policy)
  (require-command policy "mhj control-plane status")
  (require-command policy "mhj control-plane verify"))

(defun require-control-plane-verifier (policy)
  (require-members '("policy-json-valid" "status-public-redacted"
                     "lease-bounds-valid" "verifier-separation-required"
                     "manifest-debt-evaluated")
                   (policy-list policy :verifier_checks)
                   "Control-plane verifier check missing: ~A"))

(defun validate-control-plane-fields (policy)
  (require-members '("decision_kind" "policy_version" "ontology_version"
                     "authority_profile" "selected_route" "reviewer_role"
                     "verifier_role" "lease_seconds" "lease_status"
                     "evidence_refs" "output_ref")
                   (policy-list policy :required_fields)
                   "Control-plane manifest missing required field: ~A")
  (require-members '("count" "invalid_manifest_count" "manifest_debt_count"
                     "verifier_violation_count" "by_decision_kind"
                     "by_authority_profile" "by_lease_status" "checked_at")
                   (policy-list policy :public_summary_fields)
                   "Control-plane summary missing field: ~A"))

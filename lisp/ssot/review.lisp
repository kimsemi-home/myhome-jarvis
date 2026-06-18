(in-package #:myhome-jarvis.ssot)

(defparameter *review-policy*
  (list :context "AgentCluster"
        :version "v1"
        :generated_artifact "generated/review.generated.json"
        :private_review_queue "data/private/review/queue.jsonl"
        :append_only t
        :public_status_redacted t
        :raw_review_public_allowed nil
        :max_open_reviews 5
        :max_high_risk_open_reviews 0
        :min_backup_reviewers 1
        :allowed_risks #("low" "medium" "high")
        :queue_classes #("security_incident" "production_incident" "ssot_defect"
                         "user_impact_regression" "authority_boundary_change"
                         "release_blocking" "major_ontology_change"
                         "spec_change" "documentation")
        :priority_order #("security_incident" "production_incident" "ssot_defect"
                          "user_impact_regression" "authority_boundary_change"
                          "release_blocking" "major_ontology_change"
                          "spec_change" "documentation")
        :allowed_statuses #("requested" "assigned" "in_review" "approved"
                            "rejected" "deferred" "escalated")
        :requester_roles #("producer" "independent_reviewer"
                           "adversarial_reviewer" "deterministic_verifier"
                           "governance_steward")
        :reviewer_roles #("independent_reviewer" "adversarial_reviewer"
                          "deterministic_verifier" "governance_steward"
                          "backup_steward")
        :overload_policy #((:key "low_risk_only" :allowed_when_overloaded t)
                            (:key "deterministic_verification" :allowed_when_overloaded t)
                            (:key "evidence_collection" :allowed_when_overloaded t)
                            (:key "incident_response" :allowed_when_overloaded t)
                            (:key "revalidation" :allowed_when_overloaded t)
                            (:key "high_risk_change" :allowed_when_overloaded nil)
                            (:key "major_ontology_change" :allowed_when_overloaded nil)
                            (:key "security_boundary_change" :allowed_when_overloaded nil)
                            (:key "production_change" :allowed_when_overloaded nil))
        :required_fields #("at" "item_key" "queue_class" "risk" "status"
                           "requester_role" "reviewer_role"
                           "backup_available" "evidence_refs")
        :allowed_evidence_prefixes #("data/private/" "generated/" "docs/"
                                     "cmd/" "internal/" "apps/flutter/"
                                     "lisp/" "crates/" "fixtures/"
                                     "harness/" ".github/")
        :public_summary_fields #("policy_path" "queue_path" "exists" "count"
                                 "open_count" "high_risk_open_count"
                                 "invalid_review_count" "missing_evidence_count"
                                 "missing_reviewer_count" "backup_available_count"
                                 "review_debt_count" "capacity_state" "active_rule"
                                 "max_open_reviews" "max_high_risk_open_reviews"
                                 "by_risk" "by_status" "by_reviewer_role"
                                 "by_queue_class" "last_observed_at" "checked_at")
        :forbidden_public_fields #("raw_rationale" "raw_review" "raw_review_notes"
                                   "reviewer_identity" "reviewer_name"
                                   "reviewer_email" "evidence_ref" "evidence_refs"
                                   "raw_prompt" "raw_transcript" "token" "secret"
                                   "credential" "cookie" "account_id" "card_number"
                                   "local_absolute_path" "linear_url"
                                   "private_evidence")
        :commands #("mhj review status")))

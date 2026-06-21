(in-package #:myhome-jarvis.ssot)

(defparameter *authority-policy*
  (append (authority-policy-core)
          (list :reasoning_tiers *authority-reasoning-tiers*
                :role_permissions *authority-role-permissions*
                :domain_attributes *authority-domain-attributes*
                :decisions *authority-decisions*
                :assistant_authority_profiles *assistant-authority-profiles*
                :outcomes *authority-outcomes*
                :authority_debt_classes *authority-debt-classes*
                :private_review_request_ledger
                "data/private/authority-review/requests.jsonl"
                :private_approval_decision_ledger
                "data/private/authority-review/approvals.jsonl"
                :review_record_required_fields
                #("request_id" "evidence_ref" "queue_item_ref" "queue_state"
                  "required_review_classes" "approval_granted"
                  "external_writes_allowed" "self_approval_allowed")
                :approval_decision_required_fields
                #("decision_packet_ref" "decision_packet_context"
                  "decision_packet_checked_at" "scope" "target"
                  "reviewer_boundary" "reviewer_is_requester" "expires_at"
                  "approval_granted" "repo_creation_allowed"
                  "workflow_changes_allowed" "external_writes_allowed"
                  "self_approval_allowed")
                :public_summary_fields *authority-public-summary-fields*
                :forbidden_public_fields *authority-forbidden-public-fields*
                :commands #("mhj authority status"
                            "mhj authority-review status"
                            "mhj authority-review brief"
                            "mhj authority-review request"
                            "mhj authority-review evidence"
                            "mhj authority-review queue"
                            "mhj authority-review refresh [linear-ref]"
                            "mhj authority-review record <json-payload>"
                            "mhj authority-review approval-status"
                            "mhj authority-review approve <json-payload>"))))

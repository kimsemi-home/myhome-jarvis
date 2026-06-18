(in-package #:myhome-jarvis.ssot)

(defparameter *learning-policy*
  (list :context "AgentCluster"
        :version "v1"
        :private_ledger "data/private/learning/observations.jsonl"
        :generated_artifact "generated/learning.generated.json"
        :append_only t
        :private_journal_required t
        :public_status_redacted t
        :raw_observation_public_allowed nil
        :required_fields #("kind"
                           "source"
                           "summary"
                           "evidence_refs"
                           "owner"
                           "next_action")
        :allowed_kinds #("loop_gap"
                         "evidence_debt"
                         "review_debt"
                         "revalidation_debt"
                         "translation_loss"
                         "quality_regression"
                         "ssot_defect_candidate")
        :lifecycle #("observed"
                     "evidence_recorded"
                     "classified"
                     "owner_assigned"
                     "fix_planned"
                     "verified"
                     "knowledge_updated")
        :allowed_statuses #("open" "closed")
        :allowed_evidence_prefixes #("data/private/"
                                     "generated/"
                                     "docs/"
                                     "cmd/"
                                     "internal/"
                                     "apps/flutter/"
                                     "lisp/"
                                     "crates/"
                                     "fixtures/"
                                     "harness/"
                                     ".github/")
        :public_summary_fields #("path"
                                 "policy_path"
                                 "exists"
                                 "count"
                                 "open_count"
                                 "closed_count"
                                 "by_kind"
                                 "by_stage"
                                 "last_kind"
                                 "last_stage"
                                 "last_status"
                                 "last_observed_at"
                                 "checked_at")
        :forbidden_private_markers #("raw_prompt"
                                     "raw_transcript"
                                     "bearer_token"
                                     "api_secret"
                                     "account_id"
                                     "card_number"
                                     "local_absolute_path")
        :commands #("mhj learning status"
                    "mhj learning record <json-payload>")))
